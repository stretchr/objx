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

func (o *O) Get(selector interface{}) *O {
	return o.access(selector, nil)
}

// String gets a string representation of the object
// contained.
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
