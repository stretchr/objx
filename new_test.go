package objx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {

	o := New(TestMap)

	if assert.NotNil(t, o) {
		assert.Equal(t, TestMap, o.Obj())
	}

}

func TestMustFromJSONWithError(t *testing.T) {

	_, err := FromJSON(`"name":"Mat"}`)
	assert.Error(t, err)

}

func TestFromJSON(t *testing.T) {

	o := MustFromJSON(`{"name":"Mat"}`)

	if assert.NotNil(t, o) {
		if assert.NotNil(t, o.Obj()) {
			assert.Equal(t, "Mat", o.Obj().(map[string]interface{})["name"])
		}
	}

}

func TestFromJSONWithError(t *testing.T) {

	var o *O

	assert.Panics(t, func() {
		o = MustFromJSON(`"name":"Mat"}`)
	})

	assert.Nil(t, o)

}
