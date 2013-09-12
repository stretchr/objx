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

func TestRootObj(t *testing.T) {

	obj := "hello objx"
	o := New(obj)
	o2 := New(o)
	o3 := New(o2)
	o4 := New(o3)

	assert.Equal(t, o4.RootObj(), obj)

}

func TestString(t *testing.T) {

	obj := "hello objx"
	o := New(obj)
	assert.Equal(t, o.Obj(), obj)
	assert.Equal(t, objxStringPrefix+"hello objx"+objxStringSuffix, o.String())

}

func TestGetWithinNestedObj(t *testing.T) {

	o := MSI("one", MSI("two", MSI("three", 3)))
	assert.Equal(t, 3, o.Get("one.two.three").Int())

}

func TestGetWithinNestedObj_ViaArrayAccessors(t *testing.T) {

	o := MSI("one", MSI("two", MSI("three", []*Obj{MSI("name", "Mat"), MSI("name", "Tyler")})))
	assert.Equal(t, "Tyler", o.Get("one.two.three[1].name").Str())

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
