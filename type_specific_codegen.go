package objx

import (
	"reflect"
)

/*
	Inter (interface{} and []interface{})
	--------------------------------------------------
*/

// Inter gets the value as a interface{}, returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *O) Inter(optionalDefault ...interface{}) interface{} {
	if s, ok := o.obj.(interface{}); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustInter gets the value as a interface{}.
//
// Panics if the object is not a interface{}.
func (o *O) MustInter() interface{} {
	return o.obj.(interface{})
}

// InterSlice gets the value as a []interface{}, returns the optionalDefault
// value or nil if the value is not a []interface{}.
func (o *O) InterSlice(optionalDefault ...[]interface{}) []interface{} {
	if s, ok := o.obj.([]interface{}); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustInterSlice gets the value as a []interface{}.
//
// Panics if the object is not a []interface{}.
func (o *O) MustInterSlice() []interface{} {
	return o.obj.([]interface{})
}

// IsInterSlice gets whether the object contained is a []interface{} or not.
func (o *O) IsInterSlice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]interface{})
	return ok
}

// EachInter calls the specified callback for each object
// in the []interface{}.
//
// Panics if the object is the wrong type.
func (o *O) EachInter(callback func(int, interface{}) bool) *O {

	list := o.MustInterSlice()

	for index, val := range list {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

/*
	Bool (bool and []bool)
	--------------------------------------------------
*/

// Bool gets the value as a bool, returns the optionalDefault
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

// BoolSlice gets the value as a []bool, returns the optionalDefault
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

// MustBoolSlice gets the value as a []bool.
//
// Panics if the object is not a []bool.
func (o *O) MustBoolSlice() []bool {
	return o.obj.([]bool)
}

// IsBool gets whether the object contained is a bool or not.
func (o *O) IsBool() bool {
	return o.IsKind(reflect.Bool)
}

// IsBoolSlice gets whether the object contained is a []bool or not.
func (o *O) IsBoolSlice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]bool)
	return ok
}

// EachBool calls the specified callback for each object
// in the []bool.
//
// Panics if the object is the wrong type.
func (o *O) EachBool(callback func(int, bool) bool) *O {

	list := o.MustBoolSlice()

	for index, val := range list {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

/*
	Str (string and []string)
	--------------------------------------------------
*/

// Str gets the value as a string, returns the optionalDefault
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

// StrSlice gets the value as a []string, returns the optionalDefault
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

// MustStrSlice gets the value as a []string.
//
// Panics if the object is not a []string.
func (o *O) MustStrSlice() []string {
	return o.obj.([]string)
}

// IsStr gets whether the object contained is a string or not.
func (o *O) IsStr() bool {
	return o.IsKind(reflect.String)
}

// IsStrSlice gets whether the object contained is a []string or not.
func (o *O) IsStrSlice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]string)
	return ok
}

// EachStr calls the specified callback for each object
// in the []string.
//
// Panics if the object is the wrong type.
func (o *O) EachStr(callback func(int, string) bool) *O {

	list := o.MustStrSlice()

	for index, val := range list {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

/*
	Int (int and []int)
	--------------------------------------------------
*/

// Int gets the value as a int, returns the optionalDefault
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

// IntSlice gets the value as a []int, returns the optionalDefault
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

// MustIntSlice gets the value as a []int.
//
// Panics if the object is not a []int.
func (o *O) MustIntSlice() []int {
	return o.obj.([]int)
}

// IsInt gets whether the object contained is a int or not.
func (o *O) IsInt() bool {
	return o.IsKind(reflect.Int)
}

// IsIntSlice gets whether the object contained is a []int or not.
func (o *O) IsIntSlice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]int)
	return ok
}

// EachInt calls the specified callback for each object
// in the []int.
//
// Panics if the object is the wrong type.
func (o *O) EachInt(callback func(int, int) bool) *O {

	list := o.MustIntSlice()

	for index, val := range list {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

/*
	Int8 (int8 and []int8)
	--------------------------------------------------
*/

// Int8 gets the value as a int8, returns the optionalDefault
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

// Int8Slice gets the value as a []int8, returns the optionalDefault
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

// MustInt8Slice gets the value as a []int8.
//
// Panics if the object is not a []int8.
func (o *O) MustInt8Slice() []int8 {
	return o.obj.([]int8)
}

// IsInt8 gets whether the object contained is a int8 or not.
func (o *O) IsInt8() bool {
	return o.IsKind(reflect.Int8)
}

// IsInt8Slice gets whether the object contained is a []int8 or not.
func (o *O) IsInt8Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]int8)
	return ok
}

// EachInt8 calls the specified callback for each object
// in the []int8.
//
// Panics if the object is the wrong type.
func (o *O) EachInt8(callback func(int, int8) bool) *O {

	list := o.MustInt8Slice()

	for index, val := range list {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

/*
	Int16 (int16 and []int16)
	--------------------------------------------------
*/

// Int16 gets the value as a int16, returns the optionalDefault
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

// Int16Slice gets the value as a []int16, returns the optionalDefault
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

// MustInt16Slice gets the value as a []int16.
//
// Panics if the object is not a []int16.
func (o *O) MustInt16Slice() []int16 {
	return o.obj.([]int16)
}

// IsInt16 gets whether the object contained is a int16 or not.
func (o *O) IsInt16() bool {
	return o.IsKind(reflect.Int16)
}

// IsInt16Slice gets whether the object contained is a []int16 or not.
func (o *O) IsInt16Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]int16)
	return ok
}

// EachInt16 calls the specified callback for each object
// in the []int16.
//
// Panics if the object is the wrong type.
func (o *O) EachInt16(callback func(int, int16) bool) *O {

	list := o.MustInt16Slice()

	for index, val := range list {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

/*
	Int32 (int32 and []int32)
	--------------------------------------------------
*/

// Int32 gets the value as a int32, returns the optionalDefault
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

// Int32Slice gets the value as a []int32, returns the optionalDefault
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

// MustInt32Slice gets the value as a []int32.
//
// Panics if the object is not a []int32.
func (o *O) MustInt32Slice() []int32 {
	return o.obj.([]int32)
}

// IsInt32 gets whether the object contained is a int32 or not.
func (o *O) IsInt32() bool {
	return o.IsKind(reflect.Int32)
}

// IsInt32Slice gets whether the object contained is a []int32 or not.
func (o *O) IsInt32Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]int32)
	return ok
}

// EachInt32 calls the specified callback for each object
// in the []int32.
//
// Panics if the object is the wrong type.
func (o *O) EachInt32(callback func(int, int32) bool) *O {

	list := o.MustInt32Slice()

	for index, val := range list {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

/*
	Int64 (int64 and []int64)
	--------------------------------------------------
*/

// Int64 gets the value as a int64, returns the optionalDefault
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

// Int64Slice gets the value as a []int64, returns the optionalDefault
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

// MustInt64Slice gets the value as a []int64.
//
// Panics if the object is not a []int64.
func (o *O) MustInt64Slice() []int64 {
	return o.obj.([]int64)
}

// IsInt64 gets whether the object contained is a int64 or not.
func (o *O) IsInt64() bool {
	return o.IsKind(reflect.Int64)
}

// IsInt64Slice gets whether the object contained is a []int64 or not.
func (o *O) IsInt64Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]int64)
	return ok
}

// EachInt64 calls the specified callback for each object
// in the []int64.
//
// Panics if the object is the wrong type.
func (o *O) EachInt64(callback func(int, int64) bool) *O {

	list := o.MustInt64Slice()

	for index, val := range list {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

/*
	Uint (uint and []uint)
	--------------------------------------------------
*/

// Uint gets the value as a uint, returns the optionalDefault
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

// UintSlice gets the value as a []uint, returns the optionalDefault
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

// MustUintSlice gets the value as a []uint.
//
// Panics if the object is not a []uint.
func (o *O) MustUintSlice() []uint {
	return o.obj.([]uint)
}

// IsUint gets whether the object contained is a uint or not.
func (o *O) IsUint() bool {
	return o.IsKind(reflect.Uint)
}

// IsUintSlice gets whether the object contained is a []uint or not.
func (o *O) IsUintSlice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]uint)
	return ok
}

// EachUint calls the specified callback for each object
// in the []uint.
//
// Panics if the object is the wrong type.
func (o *O) EachUint(callback func(int, uint) bool) *O {

	list := o.MustUintSlice()

	for index, val := range list {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

/*
	Uint8 (uint8 and []uint8)
	--------------------------------------------------
*/

// Uint8 gets the value as a uint8, returns the optionalDefault
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

// Uint8Slice gets the value as a []uint8, returns the optionalDefault
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

// MustUint8Slice gets the value as a []uint8.
//
// Panics if the object is not a []uint8.
func (o *O) MustUint8Slice() []uint8 {
	return o.obj.([]uint8)
}

// IsUint8 gets whether the object contained is a uint8 or not.
func (o *O) IsUint8() bool {
	return o.IsKind(reflect.Uint8)
}

// IsUint8Slice gets whether the object contained is a []uint8 or not.
func (o *O) IsUint8Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]uint8)
	return ok
}

// EachUint8 calls the specified callback for each object
// in the []uint8.
//
// Panics if the object is the wrong type.
func (o *O) EachUint8(callback func(int, uint8) bool) *O {

	list := o.MustUint8Slice()

	for index, val := range list {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

/*
	Uint16 (uint16 and []uint16)
	--------------------------------------------------
*/

// Uint16 gets the value as a uint16, returns the optionalDefault
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

// Uint16Slice gets the value as a []uint16, returns the optionalDefault
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

// MustUint16Slice gets the value as a []uint16.
//
// Panics if the object is not a []uint16.
func (o *O) MustUint16Slice() []uint16 {
	return o.obj.([]uint16)
}

// IsUint16 gets whether the object contained is a uint16 or not.
func (o *O) IsUint16() bool {
	return o.IsKind(reflect.Uint16)
}

// IsUint16Slice gets whether the object contained is a []uint16 or not.
func (o *O) IsUint16Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]uint16)
	return ok
}

// EachUint16 calls the specified callback for each object
// in the []uint16.
//
// Panics if the object is the wrong type.
func (o *O) EachUint16(callback func(int, uint16) bool) *O {

	list := o.MustUint16Slice()

	for index, val := range list {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

/*
	Uint32 (uint32 and []uint32)
	--------------------------------------------------
*/

// Uint32 gets the value as a uint32, returns the optionalDefault
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

// Uint32Slice gets the value as a []uint32, returns the optionalDefault
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

// MustUint32Slice gets the value as a []uint32.
//
// Panics if the object is not a []uint32.
func (o *O) MustUint32Slice() []uint32 {
	return o.obj.([]uint32)
}

// IsUint32 gets whether the object contained is a uint32 or not.
func (o *O) IsUint32() bool {
	return o.IsKind(reflect.Uint32)
}

// IsUint32Slice gets whether the object contained is a []uint32 or not.
func (o *O) IsUint32Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]uint32)
	return ok
}

// EachUint32 calls the specified callback for each object
// in the []uint32.
//
// Panics if the object is the wrong type.
func (o *O) EachUint32(callback func(int, uint32) bool) *O {

	list := o.MustUint32Slice()

	for index, val := range list {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

/*
	Uint64 (uint64 and []uint64)
	--------------------------------------------------
*/

// Uint64 gets the value as a uint64, returns the optionalDefault
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

// Uint64Slice gets the value as a []uint64, returns the optionalDefault
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

// MustUint64Slice gets the value as a []uint64.
//
// Panics if the object is not a []uint64.
func (o *O) MustUint64Slice() []uint64 {
	return o.obj.([]uint64)
}

// IsUint64 gets whether the object contained is a uint64 or not.
func (o *O) IsUint64() bool {
	return o.IsKind(reflect.Uint64)
}

// IsUint64Slice gets whether the object contained is a []uint64 or not.
func (o *O) IsUint64Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]uint64)
	return ok
}

// EachUint64 calls the specified callback for each object
// in the []uint64.
//
// Panics if the object is the wrong type.
func (o *O) EachUint64(callback func(int, uint64) bool) *O {

	list := o.MustUint64Slice()

	for index, val := range list {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

/*
	Uintptr (uintptr and []uintptr)
	--------------------------------------------------
*/

// Uintptr gets the value as a uintptr, returns the optionalDefault
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

// UintptrSlice gets the value as a []uintptr, returns the optionalDefault
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

// MustUintptrSlice gets the value as a []uintptr.
//
// Panics if the object is not a []uintptr.
func (o *O) MustUintptrSlice() []uintptr {
	return o.obj.([]uintptr)
}

// IsUintptr gets whether the object contained is a uintptr or not.
func (o *O) IsUintptr() bool {
	return o.IsKind(reflect.Uintptr)
}

// IsUintptrSlice gets whether the object contained is a []uintptr or not.
func (o *O) IsUintptrSlice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]uintptr)
	return ok
}

// EachUintptr calls the specified callback for each object
// in the []uintptr.
//
// Panics if the object is the wrong type.
func (o *O) EachUintptr(callback func(int, uintptr) bool) *O {

	list := o.MustUintptrSlice()

	for index, val := range list {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

/*
	Float32 (float32 and []float32)
	--------------------------------------------------
*/

// Float32 gets the value as a float32, returns the optionalDefault
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

// Float32Slice gets the value as a []float32, returns the optionalDefault
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

// MustFloat32Slice gets the value as a []float32.
//
// Panics if the object is not a []float32.
func (o *O) MustFloat32Slice() []float32 {
	return o.obj.([]float32)
}

// IsFloat32 gets whether the object contained is a float32 or not.
func (o *O) IsFloat32() bool {
	return o.IsKind(reflect.Float32)
}

// IsFloat32Slice gets whether the object contained is a []float32 or not.
func (o *O) IsFloat32Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]float32)
	return ok
}

// EachFloat32 calls the specified callback for each object
// in the []float32.
//
// Panics if the object is the wrong type.
func (o *O) EachFloat32(callback func(int, float32) bool) *O {

	list := o.MustFloat32Slice()

	for index, val := range list {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

/*
	Float64 (float64 and []float64)
	--------------------------------------------------
*/

// Float64 gets the value as a float64, returns the optionalDefault
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

// Float64Slice gets the value as a []float64, returns the optionalDefault
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

// MustFloat64Slice gets the value as a []float64.
//
// Panics if the object is not a []float64.
func (o *O) MustFloat64Slice() []float64 {
	return o.obj.([]float64)
}

// IsFloat64 gets whether the object contained is a float64 or not.
func (o *O) IsFloat64() bool {
	return o.IsKind(reflect.Float64)
}

// IsFloat64Slice gets whether the object contained is a []float64 or not.
func (o *O) IsFloat64Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]float64)
	return ok
}

// EachFloat64 calls the specified callback for each object
// in the []float64.
//
// Panics if the object is the wrong type.
func (o *O) EachFloat64(callback func(int, float64) bool) *O {

	list := o.MustFloat64Slice()

	for index, val := range list {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

/*
	Complex64 (complex64 and []complex64)
	--------------------------------------------------
*/

// Complex64 gets the value as a complex64, returns the optionalDefault
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

// Complex64Slice gets the value as a []complex64, returns the optionalDefault
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

// MustComplex64Slice gets the value as a []complex64.
//
// Panics if the object is not a []complex64.
func (o *O) MustComplex64Slice() []complex64 {
	return o.obj.([]complex64)
}

// IsComplex64 gets whether the object contained is a complex64 or not.
func (o *O) IsComplex64() bool {
	return o.IsKind(reflect.Complex64)
}

// IsComplex64Slice gets whether the object contained is a []complex64 or not.
func (o *O) IsComplex64Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]complex64)
	return ok
}

// EachComplex64 calls the specified callback for each object
// in the []complex64.
//
// Panics if the object is the wrong type.
func (o *O) EachComplex64(callback func(int, complex64) bool) *O {

	list := o.MustComplex64Slice()

	for index, val := range list {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

/*
	Complex128 (complex128 and []complex128)
	--------------------------------------------------
*/

// Complex128 gets the value as a complex128, returns the optionalDefault
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

// Complex128Slice gets the value as a []complex128, returns the optionalDefault
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

// MustComplex128Slice gets the value as a []complex128.
//
// Panics if the object is not a []complex128.
func (o *O) MustComplex128Slice() []complex128 {
	return o.obj.([]complex128)
}

// IsComplex128 gets whether the object contained is a complex128 or not.
func (o *O) IsComplex128() bool {
	return o.IsKind(reflect.Complex128)
}

// IsComplex128Slice gets whether the object contained is a []complex128 or not.
func (o *O) IsComplex128Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]complex128)
	return ok
}

// EachComplex128 calls the specified callback for each object
// in the []complex128.
//
// Panics if the object is the wrong type.
func (o *O) EachComplex128(callback func(int, complex128) bool) *O {

	list := o.MustComplex128Slice()

	for index, val := range list {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}
