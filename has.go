package objx

// Has gets whether there is something at the specified selector
// or not.
func (o *Obj) Has(selector string) bool {
	if o == nil {
		return false
	}
	return !o.Get(selector).IsNil()
}
