package objx

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestKind(t *testing.T) {

	var o *Obj

	o = New(bool(true))
	assert.Equal(t, o.Kind(), reflect.Bool)

	o = New(bool(false))
	assert.Equal(t, o.Kind(), reflect.Bool)

	o = New(int(1))
	assert.Equal(t, o.Kind(), reflect.Int)

	o = New(int8(1))
	assert.Equal(t, o.Kind(), reflect.Int8)

	o = New(int16(1))
	assert.Equal(t, o.Kind(), reflect.Int16)

	o = New(int32(1))
	assert.Equal(t, o.Kind(), reflect.Int32)

	o = New(int64(1))
	assert.Equal(t, o.Kind(), reflect.Int64)

	o = New(uint(1))
	assert.Equal(t, o.Kind(), reflect.Uint)

	o = New(uint8(1))
	assert.Equal(t, o.Kind(), reflect.Uint8)

	o = New(uint16(1))
	assert.Equal(t, o.Kind(), reflect.Uint16)

	o = New(uint32(1))
	assert.Equal(t, o.Kind(), reflect.Uint32)

	o = New(uint64(1))
	assert.Equal(t, o.Kind(), reflect.Uint64)

	o = New(float32(1))
	assert.Equal(t, o.Kind(), reflect.Float32)

	o = New(float64(1))
	assert.Equal(t, o.Kind(), reflect.Float64)

	o = New(complex64(1))
	assert.Equal(t, o.Kind(), reflect.Complex64)

	o = New(complex128(1))
	assert.Equal(t, o.Kind(), reflect.Complex128)

	o = New(string("1"))
	assert.Equal(t, o.Kind(), reflect.String)

	o = New(func() {})
	assert.Equal(t, o.Kind(), reflect.Func)

}
