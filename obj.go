package objx

import (
	"reflect"
)

// Obj is the main container for objects.
type Obj struct {
	// obj is the raw object being managed by this
	// Obj object.
	obj interface{}

	// kind stores the reflect.Kind of the object stored
	// in this Obj.
	kind reflect.Kind
}

// Nil represents a nil Obj.
var Nil *Obj = nil
