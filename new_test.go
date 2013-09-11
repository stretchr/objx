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

func TestMustJsonWithError(t *testing.T) {

	_, err := Json(`"name":"Mat"}`)
	assert.Error(t, err)

}

func TestJson(t *testing.T) {

	o := MustJson(`{"name":"Mat"}`)

	if assert.NotNil(t, o) {
		if assert.NotNil(t, o.Obj()) {
			assert.Equal(t, "Mat", o.Obj().(map[string]interface{})["name"])
		}
	}

}

func TestJsonWithError(t *testing.T) {

	var o *O

	assert.Panics(t, func() {
		o = MustJson(`"name":"Mat"}`)
	})

	assert.Nil(t, o)

}
