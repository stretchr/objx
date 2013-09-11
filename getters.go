package objx

import (
	"fmt"
)

const (
	objxStringPrefix = "•"
	objxStringSuffix = "•"
)

// Obj gets the underlying object contained within this
// O.
func (o *O) Obj() interface{} {
	return o.obj
}

// Get gets the value using the specified selector and
// returns it inside a new O object.
func (o *O) Get(selector interface{}) *O {
	return o.access(selector, nil, false)
}

// String gets a string representation of the object
// contained.  This is not to be confused with Str, which gets
// the real string value currently contained.
func (o *O) String() string {

	// if the object has a string method, just call it
	if s, ok := o.obj.(interface {
		String() string
	}); ok {
		return objxStringPrefix + s.String() + objxStringSuffix
	}

	// otherwise, let fmt do the work
	return fmt.Sprintf(objxStringPrefix+"%v"+objxStringSuffix, o.obj)

}
