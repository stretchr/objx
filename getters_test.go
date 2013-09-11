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
	assert.Equal(t, "29", New(m).Get("age").Str())

	assert.Panics(t, func() {
		New(m).Get("age").MustStr()
	})

}
