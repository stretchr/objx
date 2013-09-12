package objx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHas(t *testing.T) {

	o := New(TestMap)

	assert.True(t, o.Has("name"))
	assert.True(t, o.Has("address.state"))
	assert.True(t, o.Has("numbers[4]"))

	assert.False(t, o.Has("address.state.nope"))
	assert.False(t, o.Has("address.nope"))
	assert.False(t, o.Has("nope"))
	assert.False(t, o.Has("numbers[5]"))

	o = nil
	assert.False(t, o.Has("nothing"))

}
