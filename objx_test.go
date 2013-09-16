package objx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObjxCreation(t *testing.T) {

	objx := new(Objx)
	if assert.NotNil(t, objx) {
		assert.Nil(t, objx.data)
	}

	objx = New(nil)
	if assert.NotNil(t, objx) {
		assert.Nil(t, objx.data)
	}

	objx = New("Tyler")
	if assert.NotNil(t, objx) {
		if assert.NotNil(t, objx.data) {
			assert.Equal(t, "Tyler", objx.data)
		}
	}

	objx = MSI("name", "Tyler")
	if assert.NotNil(t, objx) {
		if assert.NotNil(t, objx.data) {
			assert.Equal(t, objx.data.(map[string]interface{})["name"], "Tyler")
		}
	}

	objx = Slice("Tyler", "Mat", "Ryan")
	if assert.NotNil(t, objx) {
		if assert.NotNil(t, objx.data) {
			array := objx.data.([]interface{})
			assert.Equal(t, array[0], "Tyler")
			assert.Equal(t, array[1], "Mat")
			assert.Equal(t, array[2], "Ryan")
		}
	}

}
