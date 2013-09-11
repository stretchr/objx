package objx

import (
	"reflect"
)

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

// IsString gets whether the object contained is a string or not.
func (o *O) IsString() bool {
	return o.IsKind(reflect.String)
}

// IsStringSlice gets whether the object contained is a []string or not.
func (o *O) IsStringSlice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]string)
	return ok
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
