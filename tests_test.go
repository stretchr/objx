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

func TestHasWithNil(t *testing.T) {
	m := objx.Map{
		"nilField": nil,
		"nested": objx.Map{
			"nilChild": nil,
			"valid":    "hello",
		},
		"nilSlice": []interface{}{nil, "value"},
	}

	assert.True(t, m.Has("nilField"))
	assert.True(t, m.Has("nested.nilChild"))
	assert.True(t, m.Has("nested.valid"))
	assert.True(t, m.Has("nilSlice[0]"))
	assert.True(t, m.Has("nilSlice[1]"))

	assert.False(t, m.Has("nonExistent"))
	assert.False(t, m.Has("nested.nonExistent"))
	assert.False(t, m.Has("nilSlice[2]"))
	assert.False(t, m.Has("nilField.child"))
	assert.False(t, m.Has("nilSlice[0].child"))

	// Verify we can differentiate nil value from not found
	assert.True(t, m.Has("nilField"))
	assert.True(t, m.Get("nilField").IsNil())

	assert.False(t, m.Has("nonExistent"))
	assert.True(t, m.Get("nonExistent").IsNil())
}

func TestHasDifferentiateNilFromNotFound(t *testing.T) {
	obj := objx.Map(map[string]interface{}{
		"foo": map[string]interface{}{
			"bar": 5,
			"baz": nil,
		},
	})

	assert.True(t, obj.Has("foo.bar"))
	assert.True(t, obj.Has("foo.baz"))
	assert.False(t, obj.Has("foo.qux"))
}
