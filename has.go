package objx

// Has gets whether there is something at the specified selector
// or not.
//
// If o is nil, Has will always return false.
func (o *Obj) Has(selector interface{}) bool {
	if o == nil {
		return false
	}
	return !o.Get(selector).IsNil()
}
