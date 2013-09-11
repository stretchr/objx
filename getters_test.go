package objx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObj(t *testing.T) {

	obj := "hello objx"
	o := New(obj)
	assert.Equal(t, o.Obj(), obj)

}

func TestString(t *testing.T) {

	obj := "hello objx"
	o := New(obj)
	assert.Equal(t, o.Obj(), obj)
	assert.Equal(t, objxStringPrefix+"hello objx"+objxStringSuffix, o.String())

}

func TestStr(t *testing.T) {

	m := map[string]interface{}{"name": "Mat", "age": 29}
	assert.Equal(t, "Mat", New(m).Get("name").Str())
	assert.Equal(t, "", New(m).Get("age").Str())
	assert.Equal(t, "def", New(m).Get("age").Str("def"))

	assert.Panics(t, func() {
		New(m).Get("age").MustStr()
	})

}

func TestBool(t *testing.T) {

	m := map[string]interface{}{"value": bool(true), "nothing": nil}
	assert.Equal(t, bool(true), New(m).Get("value").Bool())
	assert.Equal(t, bool(false), New(m).Get("nothing").Bool(false))
	assert.Equal(t, bool(true), New(m).Get("nothing").Bool(true))

	assert.Panics(t, func() {
		New(m).Get("age").MustBool()
	})

}
func TestInt(t *testing.T) {

	m := map[string]interface{}{"value": int(1), "nothing": nil}
	assert.Equal(t, int(1), New(m).Get("value").Int())
	assert.Equal(t, int(0), New(m).Get("nothing").Int())
	assert.Equal(t, int(123), New(m).Get("nothing").Int(123))

	assert.Panics(t, func() {
		New(m).Get("age").MustInt()
	})

}
func TestInt8(t *testing.T) {

	m := map[string]interface{}{"value": int8(1), "nothing": nil}
	assert.Equal(t, int8(1), New(m).Get("value").Int8())
	assert.Equal(t, int8(0), New(m).Get("nothing").Int8())
	assert.Equal(t, int8(123), New(m).Get("nothing").Int8(123))

	assert.Panics(t, func() {
		New(m).Get("age").MustInt8()
	})

}
func TestInt16(t *testing.T) {

	m := map[string]interface{}{"value": int16(1), "nothing": nil}
	assert.Equal(t, int16(1), New(m).Get("value").Int16())
	assert.Equal(t, int16(0), New(m).Get("nothing").Int16())
	assert.Equal(t, int16(123), New(m).Get("nothing").Int16(123))

	assert.Panics(t, func() {
		New(m).Get("age").MustInt16()
	})

}
func TestInt32(t *testing.T) {

	m := map[string]interface{}{"value": int32(1), "nothing": nil}
	assert.Equal(t, int32(1), New(m).Get("value").Int32())
	assert.Equal(t, int32(0), New(m).Get("nothing").Int32())
	assert.Equal(t, int32(123), New(m).Get("nothing").Int32(123))

	assert.Panics(t, func() {
		New(m).Get("age").MustInt32()
	})

}
func TestInt64(t *testing.T) {

	m := map[string]interface{}{"value": int64(1), "nothing": nil}
	assert.Equal(t, int64(1), New(m).Get("value").Int64())
	assert.Equal(t, int64(0), New(m).Get("nothing").Int64())
	assert.Equal(t, int64(123), New(m).Get("nothing").Int64(123))

	assert.Panics(t, func() {
		New(m).Get("age").MustInt64()
	})

}
func TestUint(t *testing.T) {

	m := map[string]interface{}{"value": uint(1), "nothing": nil}
	assert.Equal(t, uint(1), New(m).Get("value").Uint())
	assert.Equal(t, uint(0), New(m).Get("nothing").Uint())
	assert.Equal(t, uint(123), New(m).Get("nothing").Uint(123))

	assert.Panics(t, func() {
		New(m).Get("age").MustUint()
	})

}
func TestUint8(t *testing.T) {

	m := map[string]interface{}{"value": uint8(1), "nothing": nil}
	assert.Equal(t, uint8(1), New(m).Get("value").Uint8())
	assert.Equal(t, uint8(0), New(m).Get("nothing").Uint8())
	assert.Equal(t, uint8(123), New(m).Get("nothing").Uint8(123))

	assert.Panics(t, func() {
		New(m).Get("age").MustUint8()
	})

}
func TestUint16(t *testing.T) {

	m := map[string]interface{}{"value": uint16(1), "nothing": nil}
	assert.Equal(t, uint16(1), New(m).Get("value").Uint16())
	assert.Equal(t, uint16(0), New(m).Get("nothing").Uint16())
	assert.Equal(t, uint16(123), New(m).Get("nothing").Uint16(123))

	assert.Panics(t, func() {
		New(m).Get("age").MustUint16()
	})

}
func TestUint32(t *testing.T) {

	m := map[string]interface{}{"value": uint32(1), "nothing": nil}
	assert.Equal(t, uint32(1), New(m).Get("value").Uint32())
	assert.Equal(t, uint32(0), New(m).Get("nothing").Uint32())
	assert.Equal(t, uint32(123), New(m).Get("nothing").Uint32(123))

	assert.Panics(t, func() {
		New(m).Get("age").MustUint32()
	})

}
func TestUint64(t *testing.T) {

	m := map[string]interface{}{"value": uint64(1), "nothing": nil}
	assert.Equal(t, uint64(1), New(m).Get("value").Uint64())
	assert.Equal(t, uint64(0), New(m).Get("nothing").Uint64())
	assert.Equal(t, uint64(123), New(m).Get("nothing").Uint64(123))

	assert.Panics(t, func() {
		New(m).Get("age").MustUint64()
	})

}
func TestUintptr(t *testing.T) {

	m := map[string]interface{}{"value": uintptr(1), "nothing": nil}
	assert.Equal(t, uintptr(1), New(m).Get("value").Uintptr())
	assert.Equal(t, uintptr(0), New(m).Get("nothing").Uintptr())
	assert.Equal(t, uintptr(123), New(m).Get("nothing").Uintptr(123))

	assert.Panics(t, func() {
		New(m).Get("age").MustUintptr()
	})

}
func TestFloat32(t *testing.T) {

	m := map[string]interface{}{"value": float32(1), "nothing": nil}
	assert.Equal(t, float32(1), New(m).Get("value").Float32())
	assert.Equal(t, float32(0), New(m).Get("nothing").Float32())
	assert.Equal(t, float32(123), New(m).Get("nothing").Float32(123))

	assert.Panics(t, func() {
		New(m).Get("age").MustFloat32()
	})

}
func TestFloat64(t *testing.T) {

	m := map[string]interface{}{"value": float64(1), "nothing": nil}
	assert.Equal(t, float64(1), New(m).Get("value").Float64())
	assert.Equal(t, float64(0), New(m).Get("nothing").Float64())
	assert.Equal(t, float64(123), New(m).Get("nothing").Float64(123))

	assert.Panics(t, func() {
		New(m).Get("age").MustFloat64()
	})

}
func TestComplex64(t *testing.T) {

	m := map[string]interface{}{"value": complex64(1), "nothing": nil}
	assert.Equal(t, complex64(1), New(m).Get("value").Complex64())
	assert.Equal(t, complex64(0), New(m).Get("nothing").Complex64())
	assert.Equal(t, complex64(123), New(m).Get("nothing").Complex64(123))

	assert.Panics(t, func() {
		New(m).Get("age").MustComplex64()
	})

}
func TestComplex128(t *testing.T) {

	m := map[string]interface{}{"value": complex128(1), "nothing": nil}
	assert.Equal(t, complex128(1), New(m).Get("value").Complex128())
	assert.Equal(t, complex128(0), New(m).Get("nothing").Complex128())
	assert.Equal(t, complex128(123), New(m).Get("nothing").Complex128(123))

	assert.Panics(t, func() {
		New(m).Get("age").MustComplex128()
	})

}

func TestBoolSlice(t *testing.T) {

	m := map[string]interface{}{"value": []bool{bool(true)}, "nothing": nil}
	assert.Equal(t, true, New(m).Get("value").BoolSlice()[0])
	assert.Equal(t, []bool(nil), New(m).Get("nothing").BoolSlice())
	assert.Equal(t, true, New(m).Get("nothing").BoolSlice([]bool{true})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustBool()
	})

}
func TestIntSlice(t *testing.T) {

	m := map[string]interface{}{"value": []int{int(1)}, "nothing": nil}
	assert.Equal(t, int(1), New(m).Get("value").IntSlice()[0])
	assert.Equal(t, []int(nil), New(m).Get("nothing").IntSlice())
	assert.Equal(t, int(123), New(m).Get("nothing").IntSlice([]int{123})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustInt()
	})

}
func TestInt8Slice(t *testing.T) {

	m := map[string]interface{}{"value": []int8{int8(1)}, "nothing": nil}
	assert.Equal(t, int8(1), New(m).Get("value").Int8Slice()[0])
	assert.Equal(t, []int8(nil), New(m).Get("nothing").Int8Slice())
	assert.Equal(t, int8(123), New(m).Get("nothing").Int8Slice([]int8{123})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustInt8()
	})

}
func TestInt16Slice(t *testing.T) {

	m := map[string]interface{}{"value": []int16{int16(1)}, "nothing": nil}
	assert.Equal(t, int16(1), New(m).Get("value").Int16Slice()[0])
	assert.Equal(t, []int16(nil), New(m).Get("nothing").Int16Slice())
	assert.Equal(t, int16(123), New(m).Get("nothing").Int16Slice([]int16{123})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustInt16()
	})

}
func TestInt32Slice(t *testing.T) {

	m := map[string]interface{}{"value": []int32{int32(1)}, "nothing": nil}
	assert.Equal(t, int32(1), New(m).Get("value").Int32Slice()[0])
	assert.Equal(t, []int32(nil), New(m).Get("nothing").Int32Slice())
	assert.Equal(t, int32(123), New(m).Get("nothing").Int32Slice([]int32{123})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustInt32()
	})

}
func TestInt64Slice(t *testing.T) {

	m := map[string]interface{}{"value": []int64{int64(1)}, "nothing": nil}
	assert.Equal(t, int64(1), New(m).Get("value").Int64Slice()[0])
	assert.Equal(t, []int64(nil), New(m).Get("nothing").Int64Slice())
	assert.Equal(t, int64(123), New(m).Get("nothing").Int64Slice([]int64{123})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustInt64()
	})

}
func TestUintSlice(t *testing.T) {

	m := map[string]interface{}{"value": []uint{uint(1)}, "nothing": nil}
	assert.Equal(t, uint(1), New(m).Get("value").UintSlice()[0])
	assert.Equal(t, []uint(nil), New(m).Get("nothing").UintSlice())
	assert.Equal(t, uint(123), New(m).Get("nothing").UintSlice([]uint{123})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustUint()
	})

}
func TestUint8Slice(t *testing.T) {

	m := map[string]interface{}{"value": []uint8{uint8(1)}, "nothing": nil}
	assert.Equal(t, uint8(1), New(m).Get("value").Uint8Slice()[0])
	assert.Equal(t, []uint8(nil), New(m).Get("nothing").Uint8Slice())
	assert.Equal(t, uint8(123), New(m).Get("nothing").Uint8Slice([]uint8{123})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustUint8()
	})

}
func TestUint16Slice(t *testing.T) {

	m := map[string]interface{}{"value": []uint16{uint16(1)}, "nothing": nil}
	assert.Equal(t, uint16(1), New(m).Get("value").Uint16Slice()[0])
	assert.Equal(t, []uint16(nil), New(m).Get("nothing").Uint16Slice())
	assert.Equal(t, uint16(123), New(m).Get("nothing").Uint16Slice([]uint16{123})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustUint16()
	})

}
func TestUint32Slice(t *testing.T) {

	m := map[string]interface{}{"value": []uint32{uint32(1)}, "nothing": nil}
	assert.Equal(t, uint32(1), New(m).Get("value").Uint32Slice()[0])
	assert.Equal(t, []uint32(nil), New(m).Get("nothing").Uint32Slice())
	assert.Equal(t, uint32(123), New(m).Get("nothing").Uint32Slice([]uint32{123})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustUint32()
	})

}
func TestUint64Slice(t *testing.T) {

	m := map[string]interface{}{"value": []uint64{uint64(1)}, "nothing": nil}
	assert.Equal(t, uint64(1), New(m).Get("value").Uint64Slice()[0])
	assert.Equal(t, []uint64(nil), New(m).Get("nothing").Uint64Slice())
	assert.Equal(t, uint64(123), New(m).Get("nothing").Uint64Slice([]uint64{123})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustUint64()
	})

}
func TestUintptrSlice(t *testing.T) {

	m := map[string]interface{}{"value": []uintptr{uintptr(1)}, "nothing": nil}
	assert.Equal(t, uintptr(1), New(m).Get("value").UintptrSlice()[0])
	assert.Equal(t, []uintptr(nil), New(m).Get("nothing").UintptrSlice())
	assert.Equal(t, uintptr(123), New(m).Get("nothing").UintptrSlice([]uintptr{123})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustUintptr()
	})

}
func TestFloat32Slice(t *testing.T) {

	m := map[string]interface{}{"value": []float32{float32(1)}, "nothing": nil}
	assert.Equal(t, float32(1), New(m).Get("value").Float32Slice()[0])
	assert.Equal(t, []float32(nil), New(m).Get("nothing").Float32Slice())
	assert.Equal(t, float32(123), New(m).Get("nothing").Float32Slice([]float32{123})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustFloat32()
	})

}
func TestFloat64Slice(t *testing.T) {

	m := map[string]interface{}{"value": []float64{float64(1)}, "nothing": nil}
	assert.Equal(t, float64(1), New(m).Get("value").Float64Slice()[0])
	assert.Equal(t, []float64(nil), New(m).Get("nothing").Float64Slice())
	assert.Equal(t, float64(123), New(m).Get("nothing").Float64Slice([]float64{123})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustFloat64()
	})

}
func TestComplex64Slice(t *testing.T) {

	m := map[string]interface{}{"value": []complex64{complex64(1)}, "nothing": nil}
	assert.Equal(t, complex64(1), New(m).Get("value").Complex64Slice()[0])
	assert.Equal(t, []complex64(nil), New(m).Get("nothing").Complex64Slice())
	assert.Equal(t, complex64(123), New(m).Get("nothing").Complex64Slice([]complex64{123})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustComplex64()
	})

}
func TestComplex128Slice(t *testing.T) {

	m := map[string]interface{}{"value": []complex128{complex128(1)}, "nothing": nil}
	assert.Equal(t, complex128(1), New(m).Get("value").Complex128Slice()[0])
	assert.Equal(t, []complex128(nil), New(m).Get("nothing").Complex128Slice())
	assert.Equal(t, complex128(123), New(m).Get("nothing").Complex128Slice([]complex128{123})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustComplex128()
	})

}
