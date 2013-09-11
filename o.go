package objx

import (
	"reflect"
)

// O is the main container for objects.
type O struct {
	// obj is the raw object being managed by this
	// O object.
	obj interface{}

	// kind stores the reflect.Kind of the object stored
	// in this O.
	kind reflect.Kind
}

// Nil represents a nil O.
var Nil *O = nil
