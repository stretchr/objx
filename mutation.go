package objx

// Copy creates a shallow copy of the Obj.
func (o *Obj) Copy() *Obj {
	copied := make(map[string]interface{})
	for k, v := range o.MSI() {
		copied[k] = v
	}
	return New(copied)
}

// Merge blends the specified map with a copy of this map and returns the result.
//
// Keys that appear in both will be selected from the specified map.
// This method requires that the wrapped object be a map[string]interface{}
func (o *Obj) Merge(merge *Obj) *Obj {
	return o.Copy().MergeHere(merge)
}

// Merge blends the specified map with this map and returns the current map.
//
// Keys that appear in both will be selected from the specified map.  The original map
// will be modified. This method requires that
// the wrapped object be a map[string]interface{}
func (o *Obj) MergeHere(merge *Obj) *Obj {

	oMSI := o.MSI()
	for k, v := range merge.MSI() {
		oMSI[k] = v
	}

	return New(oMSI)

}

// Transform builds a new Obj giving the transformer a chance
// to change the keys and values as it goes. This method requires that
// the wrapped object be a map[string]interface{}
func (o *Obj) Transform(transformer func(key string, value interface{}) (string, interface{})) *Obj {
	m := make(map[string]interface{})
	for k, v := range o.MSI() {
		modifiedKey, modifiedVal := transformer(k, v)
		m[modifiedKey] = modifiedVal
	}
	return New(m)
}

// TransformKeys builds a new map using the specified key mapping.
//
// Unspecified keys will be unaltered.
// This method requires that the wrapped object be a map[string]interface{}
func (o *Obj) TransformKeys(mapping map[string]string) *Obj {
	return o.Transform(func(key string, value interface{}) (string, interface{}) {

		if newKey, ok := mapping[key]; ok {
			return newKey, value
		}

		return key, value
	})
}
