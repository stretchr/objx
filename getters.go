package objx

import (
	"fmt"
)

const (
	objxStringPrefix = "•"
	objxStringSuffix = "•"
)

// Obj gets the underlying object contained within this
// Obj.
func (o *Obj) Obj() interface{} {
	return o.obj
}

// Get gets the value using the specified selector and
// returns it inside a new Obj object.
//
// If it cannot find the value, Get will return a nil
// value inside an instance of Obj.
func (o *Obj) Get(selector interface{}) *Obj {
	return o.access(selector, nil, false, false)
}

// MustGet gets the value using the specified selector and
// returns it inside a new Obj object.
//
// If it cannot find the value, MustGet will panic.
func (o *Obj) MustGet(selector interface{}) *Obj {
	return o.access(selector, nil, false, true)
}

// String gets a string representation of the object
// contained.  This is not to be confused with Str, which gets
// the real string value currently contained.
func (o *Obj) String() string {

	// if the object has a string method, just call it
	if s, ok := o.obj.(interface {
		String() string
	}); ok {
		return objxStringPrefix + s.String() + objxStringSuffix
	}

	// otherwise, let fmt do the work
	return fmt.Sprintf(objxStringPrefix+"%v"+objxStringSuffix, o.obj)

}
