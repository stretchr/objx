package objx

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

// MustBool gets the value as a bool.
//
// Panics if the object is not a bool.
func (o *O) MustBool() bool {
	return o.obj.(bool)
}

// MustBoolSlice gets the value as a []bool.
//
// Panics if the object is not a []bool.
func (o *O) MustBoolSlice() []bool {
	return o.obj.([]bool)
}

// BoolSlice gets the value as a []bool, returns the optinoalDefault
// value or nil if the value is not a []bool.
func (o *O) BoolSlice(optionalDefault ...[]bool) []bool {
	if s, ok := o.obj.([]bool); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
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

// MustStr gets the value as a string.
//
// Panics if the object is not a string.
func (o *O) MustStr() string {
	return o.obj.(string)
}

// MustStrSlice gets the value as a []string.
//
// Panics if the object is not a []string.
func (o *O) MustStrSlice() []string {
	return o.obj.([]string)
}

// StringSlice gets the value as a []string, returns the optinoalDefault
// value or nil if the value is not a []string.
func (o *O) StrSlice(optionalDefault ...[]string) []string {
	if s, ok := o.obj.([]string); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
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

// MustInt gets the value as a int.
//
// Panics if the object is not a int.
func (o *O) MustInt() int {
	return o.obj.(int)
}

// MustIntSlice gets the value as a []int.
//
// Panics if the object is not a []int.
func (o *O) MustIntSlice() []int {
	return o.obj.([]int)
}

// IntSlice gets the value as a []int, returns the optinoalDefault
// value or nil if the value is not a []int.
func (o *O) IntSlice(optionalDefault ...[]int) []int {
	if s, ok := o.obj.([]int); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
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

// MustInt8 gets the value as a int8.
//
// Panics if the object is not a int8.
func (o *O) MustInt8() int8 {
	return o.obj.(int8)
}

// MustInt8Slice gets the value as a []int8.
//
// Panics if the object is not a []int8.
func (o *O) MustInt8Slice() []int8 {
	return o.obj.([]int8)
}

// Int8Slice gets the value as a []int8, returns the optinoalDefault
// value or nil if the value is not a []int8.
func (o *O) Int8Slice(optionalDefault ...[]int8) []int8 {
	if s, ok := o.obj.([]int8); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
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

// MustInt16 gets the value as a int16.
//
// Panics if the object is not a int16.
func (o *O) MustInt16() int16 {
	return o.obj.(int16)
}

// MustInt16Slice gets the value as a []int16.
//
// Panics if the object is not a []int16.
func (o *O) MustInt16Slice() []int16 {
	return o.obj.([]int16)
}

// Int16Slice gets the value as a []int16, returns the optinoalDefault
// value or nil if the value is not a []int16.
func (o *O) Int16Slice(optionalDefault ...[]int16) []int16 {
	if s, ok := o.obj.([]int16); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
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

// MustInt32 gets the value as a int32.
//
// Panics if the object is not a int32.
func (o *O) MustInt32() int32 {
	return o.obj.(int32)
}

// MustInt32Slice gets the value as a []int32.
//
// Panics if the object is not a []int32.
func (o *O) MustInt32Slice() []int32 {
	return o.obj.([]int32)
}

// Int32Slice gets the value as a []int32, returns the optinoalDefault
// value or nil if the value is not a []int32.
func (o *O) Int32Slice(optionalDefault ...[]int32) []int32 {
	if s, ok := o.obj.([]int32); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
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

// MustInt64 gets the value as a int64.
//
// Panics if the object is not a int64.
func (o *O) MustInt64() int64 {
	return o.obj.(int64)
}

// MustInt64Slice gets the value as a []int64.
//
// Panics if the object is not a []int64.
func (o *O) MustInt64Slice() []int64 {
	return o.obj.([]int64)
}

// Int64Slice gets the value as a []int64, returns the optinoalDefault
// value or nil if the value is not a []int64.
func (o *O) Int64Slice(optionalDefault ...[]int64) []int64 {
	if s, ok := o.obj.([]int64); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
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

// MustUint gets the value as a uint.
//
// Panics if the object is not a uint.
func (o *O) MustUint() uint {
	return o.obj.(uint)
}

// MustUintSlice gets the value as a []uint.
//
// Panics if the object is not a []uint.
func (o *O) MustUintSlice() []uint {
	return o.obj.([]uint)
}

// UintSlice gets the value as a []uint, returns the optinoalDefault
// value or nil if the value is not a []uint.
func (o *O) UintSlice(optionalDefault ...[]uint) []uint {
	if s, ok := o.obj.([]uint); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
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

// MustUint8 gets the value as a uint8.
//
// Panics if the object is not a uint8.
func (o *O) MustUint8() uint8 {
	return o.obj.(uint8)
}

// MustUint8Slice gets the value as a []uint8.
//
// Panics if the object is not a []uint8.
func (o *O) MustUint8Slice() []uint8 {
	return o.obj.([]uint8)
}

// Uint8Slice gets the value as a []uint8, returns the optinoalDefault
// value or nil if the value is not a []uint8.
func (o *O) Uint8Slice(optionalDefault ...[]uint8) []uint8 {
	if s, ok := o.obj.([]uint8); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
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

// MustUint16 gets the value as a uint16.
//
// Panics if the object is not a uint16.
func (o *O) MustUint16() uint16 {
	return o.obj.(uint16)
}

// MustUint16Slice gets the value as a []uint16.
//
// Panics if the object is not a []uint16.
func (o *O) MustUint16Slice() []uint16 {
	return o.obj.([]uint16)
}

// Uint16Slice gets the value as a []uint16, returns the optinoalDefault
// value or nil if the value is not a []uint16.
func (o *O) Uint16Slice(optionalDefault ...[]uint16) []uint16 {
	if s, ok := o.obj.([]uint16); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
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

// MustUint32 gets the value as a uint32.
//
// Panics if the object is not a uint32.
func (o *O) MustUint32() uint32 {
	return o.obj.(uint32)
}

// MustUint32Slice gets the value as a []uint32.
//
// Panics if the object is not a []uint32.
func (o *O) MustUint32Slice() []uint32 {
	return o.obj.([]uint32)
}

// Uint32Slice gets the value as a []uint32, returns the optinoalDefault
// value or nil if the value is not a []uint32.
func (o *O) Uint32Slice(optionalDefault ...[]uint32) []uint32 {
	if s, ok := o.obj.([]uint32); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
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

// MustUint64 gets the value as a uint64.
//
// Panics if the object is not a uint64.
func (o *O) MustUint64() uint64 {
	return o.obj.(uint64)
}

// MustUint64Slice gets the value as a []uint64.
//
// Panics if the object is not a []uint64.
func (o *O) MustUint64Slice() []uint64 {
	return o.obj.([]uint64)
}

// Uint64Slice gets the value as a []uint64, returns the optinoalDefault
// value or nil if the value is not a []uint64.
func (o *O) Uint64Slice(optionalDefault ...[]uint64) []uint64 {
	if s, ok := o.obj.([]uint64); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
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

// MustUintptr gets the value as a uintptr.
//
// Panics if the object is not a uintptr.
func (o *O) MustUintptr() uintptr {
	return o.obj.(uintptr)
}

// MustUintptrSlice gets the value as a []uintptr.
//
// Panics if the object is not a []uintptr.
func (o *O) MustUintptrSlice() []uintptr {
	return o.obj.([]uintptr)
}

// UintptrSlice gets the value as a []uintptr, returns the optinoalDefault
// value or nil if the value is not a []uintptr.
func (o *O) UintptrSlice(optionalDefault ...[]uintptr) []uintptr {
	if s, ok := o.obj.([]uintptr); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
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

// MustFloat32 gets the value as a float32.
//
// Panics if the object is not a float32.
func (o *O) MustFloat32() float32 {
	return o.obj.(float32)
}

// MustFloat32Slice gets the value as a []float32.
//
// Panics if the object is not a []float32.
func (o *O) MustFloat32Slice() []float32 {
	return o.obj.([]float32)
}

// Float32Slice gets the value as a []float32, returns the optinoalDefault
// value or nil if the value is not a []float32.
func (o *O) Float32Slice(optionalDefault ...[]float32) []float32 {
	if s, ok := o.obj.([]float32); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
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

// MustFloat64 gets the value as a float64.
//
// Panics if the object is not a float64.
func (o *O) MustFloat64() float64 {
	return o.obj.(float64)
}

// MustFloat64Slice gets the value as a []float64.
//
// Panics if the object is not a []float64.
func (o *O) MustFloat64Slice() []float64 {
	return o.obj.([]float64)
}

// Float64Slice gets the value as a []float64, returns the optinoalDefault
// value or nil if the value is not a []float64.
func (o *O) Float64Slice(optionalDefault ...[]float64) []float64 {
	if s, ok := o.obj.([]float64); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
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

// MustComplex64 gets the value as a complex64.
//
// Panics if the object is not a complex64.
func (o *O) MustComplex64() complex64 {
	return o.obj.(complex64)
}

// MustComplex64Slice gets the value as a []complex64.
//
// Panics if the object is not a []complex64.
func (o *O) MustComplex64Slice() []complex64 {
	return o.obj.([]complex64)
}

// Complex64Slice gets the value as a []complex64, returns the optinoalDefault
// value or nil if the value is not a []complex64.
func (o *O) Complex64Slice(optionalDefault ...[]complex64) []complex64 {
	if s, ok := o.obj.([]complex64); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
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

// MustComplex128 gets the value as a complex128.
//
// Panics if the object is not a complex128.
func (o *O) MustComplex128() complex128 {
	return o.obj.(complex128)
}

// MustComplex128Slice gets the value as a []complex128.
//
// Panics if the object is not a []complex128.
func (o *O) MustComplex128Slice() []complex128 {
	return o.obj.([]complex128)
}

// Complex128Slice gets the value as a []complex128, returns the optinoalDefault
// value or nil if the value is not a []complex128.
func (o *O) Complex128Slice(optionalDefault ...[]complex128) []complex128 {
	if s, ok := o.obj.([]complex128); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}
