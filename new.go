package objx

// New creates a new O containing the specified object.
func New(object interface{}) *O {
	return &O{Obj: object}
}
