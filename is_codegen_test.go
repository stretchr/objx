package objx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBool(t *testing.T) {

	var o *O

	o = New(true)
	assert.True(t, o.IsBool())

	o = New([]bool{true})
	assert.True(t, o.IsBoolSlice())

}

func TestIsString(t *testing.T) {

	var o *O

	o = New(string(1))
	assert.True(t, o.IsString())

	o = New([]string{string(1)})
	assert.True(t, o.IsStringSlice())

}

func TestIsInt(t *testing.T) {

	var o *O

	o = New(int(1))
	assert.True(t, o.IsInt())

	o = New([]int{int(1)})
	assert.True(t, o.IsIntSlice())

}

func TestIsInt8(t *testing.T) {

	var o *O

	o = New(int8(1))
	assert.True(t, o.IsInt8())

	o = New([]int8{int8(1)})
	assert.True(t, o.IsInt8Slice())

}

func TestIsInt16(t *testing.T) {

	var o *O

	o = New(int16(1))
	assert.True(t, o.IsInt16())

	o = New([]int16{int16(1)})
	assert.True(t, o.IsInt16Slice())

}

func TestIsInt32(t *testing.T) {

	var o *O

	o = New(int32(1))
	assert.True(t, o.IsInt32())

	o = New([]int32{int32(1)})
	assert.True(t, o.IsInt32Slice())

}

func TestIsInt64(t *testing.T) {

	var o *O

	o = New(int64(1))
	assert.True(t, o.IsInt64())

	o = New([]int64{int64(1)})
	assert.True(t, o.IsInt64Slice())

}

func TestIsUint(t *testing.T) {

	var o *O

	o = New(uint(1))
	assert.True(t, o.IsUint())

	o = New([]uint{uint(1)})
	assert.True(t, o.IsUintSlice())

}

func TestIsUint8(t *testing.T) {

	var o *O

	o = New(uint8(1))
	assert.True(t, o.IsUint8())

	o = New([]uint8{uint8(1)})
	assert.True(t, o.IsUint8Slice())

}

func TestIsUint16(t *testing.T) {

	var o *O

	o = New(uint16(1))
	assert.True(t, o.IsUint16())

	o = New([]uint16{uint16(1)})
	assert.True(t, o.IsUint16Slice())

}

func TestIsUint32(t *testing.T) {

	var o *O

	o = New(uint32(1))
	assert.True(t, o.IsUint32())

	o = New([]uint32{uint32(1)})
	assert.True(t, o.IsUint32Slice())

}

func TestIsUint64(t *testing.T) {

	var o *O

	o = New(uint64(1))
	assert.True(t, o.IsUint64())

	o = New([]uint64{uint64(1)})
	assert.True(t, o.IsUint64Slice())

}

func TestIsUintptr(t *testing.T) {

	var o *O

	o = New(uintptr(1))
	assert.True(t, o.IsUintptr())

	o = New([]uintptr{uintptr(1)})
	assert.True(t, o.IsUintptrSlice())

}

func TestIsFloat32(t *testing.T) {

	var o *O

	o = New(float32(1))
	assert.True(t, o.IsFloat32())

	o = New([]float32{float32(1)})
	assert.True(t, o.IsFloat32Slice())

}

func TestIsFloat64(t *testing.T) {

	var o *O

	o = New(float64(1))
	assert.True(t, o.IsFloat64())

	o = New([]float64{float64(1)})
	assert.True(t, o.IsFloat64Slice())

}

func TestIsComplex64(t *testing.T) {

	var o *O

	o = New(complex64(1))
	assert.True(t, o.IsComplex64())

	o = New([]complex64{complex64(1)})
	assert.True(t, o.IsComplex64Slice())

}

func TestIsComplex128(t *testing.T) {

	var o *O

	o = New(complex128(1))
	assert.True(t, o.IsComplex128())

	o = New([]complex128{complex128(1)})
	assert.True(t, o.IsComplex128Slice())

}
