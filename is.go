package objx

import (
	"reflect"
)

// IsNil gets whether the data is nil or not.
func (o *O) IsNil() bool {
	return o.Obj() == nil
}

// IsKind gets whether the object contained is of the
// specified reflect.Kind or not.
func (o *O) IsKind(kind reflect.Kind) bool {
	return o.Kind() == kind
}

// Gets whether the object contained is of this type
// or not.
func (o *O) IsBool() bool {
	return o.IsKind(reflect.Bool)
}

// Gets whether the object contained is of this type
// or not.
func (o *O) IsInt() bool {
	return o.IsKind(reflect.Int)
}

// Gets whether the object contained is of this type
// or not.
func (o *O) IsInt8() bool {
	return o.IsKind(reflect.Int8)
}

// Gets whether the object contained is of this type
// or not.
func (o *O) IsInt16() bool {
	return o.IsKind(reflect.Int16)
}

// Gets whether the object contained is of this type
// or not.
func (o *O) IsInt32() bool {
	return o.IsKind(reflect.Int32)
}

// Gets whether the object contained is of this type
// or not.
func (o *O) IsInt64() bool {
	return o.IsKind(reflect.Int64)
}

// Gets whether the object contained is of this type
// or not.
func (o *O) IsUint() bool {
	return o.IsKind(reflect.Uint)
}

// Gets whether the object contained is of this type
// or not.
func (o *O) IsUint8() bool {
	return o.IsKind(reflect.Uint8)
}

// Gets whether the object contained is of this type
// or not.
func (o *O) IsUint16() bool {
	return o.IsKind(reflect.Uint16)
}

// Gets whether the object contained is of this type
// or not.
func (o *O) IsUint32() bool {
	return o.IsKind(reflect.Uint32)
}

// Gets whether the object contained is of this type
// or not.
func (o *O) IsUint64() bool {
	return o.IsKind(reflect.Uint64)
}

// Gets whether the object contained is of this type
// or not.
func (o *O) IsFloat32() bool {
	return o.IsKind(reflect.Float32)
}

// Gets whether the object contained is of this type
// or not.
func (o *O) IsFloat64() bool {
	return o.IsKind(reflect.Float64)
}

// Gets whether the object contained is of this type
// or not.
func (o *O) IsComplex64() bool {
	return o.IsKind(reflect.Complex64)
}

// Gets whether the object contained is of this type
// or not.
func (o *O) IsComplex128() bool {
	return o.IsKind(reflect.Complex128)
}

// Gets whether the object contained is of this type
// or not.
func (o *O) IsString() bool {
	return o.IsKind(reflect.String)
}

// Gets whether the object contained is of this type
// or not.
func (o *O) IsUintPtr() bool {
	return o.IsKind(reflect.Uintptr)
}

// Gets whether the object contained is of this type
// or not.
func (o *O) IsFunc() bool {
	return o.IsKind(reflect.Func)
}
