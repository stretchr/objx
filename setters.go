package objx

// Set sets a value inside the object using the specified
// selector.
func (o *Obj) Set(selector interface{}, value interface{}) *Obj {
	return o.access(selector, value, true, false)
}
