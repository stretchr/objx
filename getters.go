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

// Str gets the value as a string, returns the optinoalDefault
// value or a system default object if the value is the wrong type.
func (o *O) Str(optionalDefault ...string) string {
	if s, ok := o.obj.(string); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return ""
}

// Str gets the value as a string.
//
// Panics if the object is not a string.
func (o *O) MustStr() string {
	return o.obj.(string)
}

// Bool gets the value as a bool, returns the optinoalDefault
// value or a system default object if the value is the wrong type.
func (o *O) Bool(optionalDefault ...bool) bool {
	if s, ok := o.obj.(bool); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return false
}

// Bool gets the value as a bool.
//
// Panics if the object is not a bool.
func (o *O) MustBool() bool {
	return o.obj.(bool)
}

// Int gets the value as a int, returns the optinoalDefault
// value or a system default object if the value is the wrong type.
func (o *O) Int(optionalDefault ...int) int {
	if s, ok := o.obj.(int); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// Int gets the value as a int.
//
// Panics if the object is not a int.
func (o *O) MustInt() int {
	return o.obj.(int)
}

// Int8 gets the value as a int8, returns the optinoalDefault
// value or a system default object if the value is the wrong type.
func (o *O) Int8(optionalDefault ...int8) int8 {
	if s, ok := o.obj.(int8); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// Int8 gets the value as a int8.
//
// Panics if the object is not a int8.
func (o *O) MustInt8() int8 {
	return o.obj.(int8)
}

// Int16 gets the value as a int16, returns the optinoalDefault
// value or a system default object if the value is the wrong type.
func (o *O) Int16(optionalDefault ...int16) int16 {
	if s, ok := o.obj.(int16); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// Int16 gets the value as a int16.
//
// Panics if the object is not a int16.
func (o *O) MustInt16() int16 {
	return o.obj.(int16)
}

// Int32 gets the value as a int32, returns the optinoalDefault
// value or a system default object if the value is the wrong type.
func (o *O) Int32(optionalDefault ...int32) int32 {
	if s, ok := o.obj.(int32); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// Int32 gets the value as a int32.
//
// Panics if the object is not a int32.
func (o *O) MustInt32() int32 {
	return o.obj.(int32)
}

// Int64 gets the value as a int64, returns the optinoalDefault
// value or a system default object if the value is the wrong type.
func (o *O) Int64(optionalDefault ...int64) int64 {
	if s, ok := o.obj.(int64); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// Int64 gets the value as a int64.
//
// Panics if the object is not a int64.
func (o *O) MustInt64() int64 {
	return o.obj.(int64)
}

// Uint gets the value as a uint, returns the optinoalDefault
// value or a system default object if the value is the wrong type.
func (o *O) Uint(optionalDefault ...uint) uint {
	if s, ok := o.obj.(uint); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// Uint gets the value as a uint.
//
// Panics if the object is not a uint.
func (o *O) MustUint() uint {
	return o.obj.(uint)
}

// Uint8 gets the value as a uint8, returns the optinoalDefault
// value or a system default object if the value is the wrong type.
func (o *O) Uint8(optionalDefault ...uint8) uint8 {
	if s, ok := o.obj.(uint8); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// Uint8 gets the value as a uint8.
//
// Panics if the object is not a uint8.
func (o *O) MustUint8() uint8 {
	return o.obj.(uint8)
}

// Uint16 gets the value as a uint16, returns the optinoalDefault
// value or a system default object if the value is the wrong type.
func (o *O) Uint16(optionalDefault ...uint16) uint16 {
	if s, ok := o.obj.(uint16); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// Uint16 gets the value as a uint16.
//
// Panics if the object is not a uint16.
func (o *O) MustUint16() uint16 {
	return o.obj.(uint16)
}

// Uint32 gets the value as a uint32, returns the optinoalDefault
// value or a system default object if the value is the wrong type.
func (o *O) Uint32(optionalDefault ...uint32) uint32 {
	if s, ok := o.obj.(uint32); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// Uint32 gets the value as a uint32.
//
// Panics if the object is not a uint32.
func (o *O) MustUint32() uint32 {
	return o.obj.(uint32)
}

// Uint64 gets the value as a uint64, returns the optinoalDefault
// value or a system default object if the value is the wrong type.
func (o *O) Uint64(optionalDefault ...uint64) uint64 {
	if s, ok := o.obj.(uint64); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// Uint64 gets the value as a uint64.
//
// Panics if the object is not a uint64.
func (o *O) MustUint64() uint64 {
	return o.obj.(uint64)
}

// Uintptr gets the value as a uintptr, returns the optinoalDefault
// value or a system default object if the value is the wrong type.
func (o *O) Uintptr(optionalDefault ...uintptr) uintptr {
	if s, ok := o.obj.(uintptr); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// Uintptr gets the value as a uintptr.
//
// Panics if the object is not a uintptr.
func (o *O) MustUintptr() uintptr {
	return o.obj.(uintptr)
}

// Float32 gets the value as a float32, returns the optinoalDefault
// value or a system default object if the value is the wrong type.
func (o *O) Float32(optionalDefault ...float32) float32 {
	if s, ok := o.obj.(float32); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// Float32 gets the value as a float32.
//
// Panics if the object is not a float32.
func (o *O) MustFloat32() float32 {
	return o.obj.(float32)
}

// Float64 gets the value as a float64, returns the optinoalDefault
// value or a system default object if the value is the wrong type.
func (o *O) Float64(optionalDefault ...float64) float64 {
	if s, ok := o.obj.(float64); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// Float64 gets the value as a float64.
//
// Panics if the object is not a float64.
func (o *O) MustFloat64() float64 {
	return o.obj.(float64)
}

// Complex64 gets the value as a complex64, returns the optinoalDefault
// value or a system default object if the value is the wrong type.
func (o *O) Complex64(optionalDefault ...complex64) complex64 {
	if s, ok := o.obj.(complex64); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// Complex64 gets the value as a complex64.
//
// Panics if the object is not a complex64.
func (o *O) MustComplex64() complex64 {
	return o.obj.(complex64)
}

// Complex128 gets the value as a complex128, returns the optinoalDefault
// value or a system default object if the value is the wrong type.
func (o *O) Complex128(optionalDefault ...complex128) complex128 {
	if s, ok := o.obj.(complex128); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// Complex128 gets the value as a complex128.
//
// Panics if the object is not a complex128.
func (o *O) MustComplex128() complex128 {
	return o.obj.(complex128)
}
