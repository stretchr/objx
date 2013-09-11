package objx

// Obj gets the underlying object contained within this
// O.
func (o *O) Obj() interface{} {
	return o.obj
}

func (o *O) Get(selector interface{}) *O {
	return o.access(selector, nil)
}
