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

func TestMustGet(t *testing.T) {

	o := New(TestMap)

	assert.Equal(t, "UT", o.MustGet("address.state").Obj())

	assert.Panics(t, func() {
		o.MustGet("address.state.nope")
	})
	assert.Panics(t, func() {
		o.MustGet("nothing-here")
	})
	assert.Panics(t, func() {
		o.MustGet("numbers[5]")
	})

}
