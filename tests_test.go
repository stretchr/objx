package objx_test

import (
	"testing"

	"github.com/stretchr/objx"
)

func TestHas(t *testing.T) {
	m := objx.Map(TestMap)

	assert.True(t, m.Has("name"))
	assert.True(t, m.Has("address.state"))
	assert.True(t, m.Has("numbers[4]"))

	assert.False(t, m.Has("address.state.nope"))
	assert.False(t, m.Has("address.nope"))
	assert.False(t, m.Has("nope"))
	assert.False(t, m.Has("numbers[5]"))

	m = nil

	assert.False(t, m.Has("nothing"))
}

func TestValue_IsNil(t *testing.T) {
	var nilVal *objx.Value
	assert.True(t, nilVal.IsNil())

	m := objx.Map{
		"nil":  nil,
		"str":  "hello",
		"int":  123,
		"bool": true,
	}

	assert.True(t, m.Get("nil").IsNil())
	assert.True(t, m.Get("nonexistent").IsNil())
	assert.False(t, m.Get("str").IsNil())
	assert.False(t, m.Get("int").IsNil())
	assert.False(t, m.Get("bool").IsNil())
}


