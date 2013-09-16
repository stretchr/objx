package objx

// Copy creates a shallow copy of the Obj.
func (m *Map) Copy() *Map {
	copied := make(map[string]interface{})
	for k, v := range m.value.MSI() {
		copied[k] = v
	}
	return New(copied)
}

// Merge blends the specified map with a copy of this map and returns the result.
//
// Keys that appear in both will be selected from the specified map.
// This method requires that the wrapped object be a map[string]interface{}
func (m *Map) Merge(merge *Map) *Map {
	return m.Copy().MergeHere(merge)
}

// Merge blends the specified map with this map and returns the current map.
//
// Keys that appear in both will be selected from the specified map.  The original map
// will be modified. This method requires that
// the wrapped object be a map[string]interface{}
func (m *Map) MergeHere(merge *Map) *Map {

	oMSI := m.value.MSI()
	for k, v := range merge.value.MSI() {
		oMSI[k] = v
	}

	return New(oMSI)

}

// Transform builds a new Obj giving the transformer a chance
// to change the keys and values as it goes. This method requires that
// the wrapped object be a map[string]interface{}
func (m *Map) Transform(transformer func(key string, value interface{}) (string, interface{})) *Map {
	newMap := make(map[string]interface{})
	for k, v := range m.value.MSI() {
		modifiedKey, modifiedVal := transformer(k, v)
		newMap[modifiedKey] = modifiedVal
	}
	return New(newMap)
}

// TransformKeys builds a new map using the specified key mapping.
//
// Unspecified keys will be unaltered.
// This method requires that the wrapped object be a map[string]interface{}
func (m *Map) TransformKeys(mapping map[string]string) *Map {
	return m.Transform(func(key string, value interface{}) (string, interface{}) {

		if newKey, ok := mapping[key]; ok {
			return newKey, value
		}

		return key, value
	})
}
