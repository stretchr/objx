package objx

// O is the main container for objects.
type O struct {
	// Obj is the raw object being managed by this
	// O object.
	Obj interface{}
}

// Nil represents a nil O.
var Nil *O = nil
