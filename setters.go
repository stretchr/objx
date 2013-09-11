package objx

// Set sets a value inside the object using the specified
// selector.
func (o *O) Set(selector interface{}, value interface{}) *O {
	return o.access(selector, value)
}
