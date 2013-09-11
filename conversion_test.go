package objx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConversionJSON(t *testing.T) {

	jsonString := `{"name":"Mat"}`
	o := MustFromJSON(jsonString)

	result, err := o.JSON()

	if assert.NoError(t, err) {
		assert.Equal(t, jsonString, result)
	}

	assert.Equal(t, jsonString, o.MustJSON())

}

func TestConversionJSONWithError(t *testing.T) {

	o := new(O)
	o.obj = func() {}

	assert.Panics(t, func() {
		o.MustJSON()
	})

	_, err := o.JSON()

	assert.Error(t, err)

}

func TestConversionBase64(t *testing.T) {

	o := New(map[string]interface{}{"name": "Mat"})

	result, err := o.Base64()

	if assert.NoError(t, err) {
		assert.Equal(t, "eyJuYW1lIjoiTWF0In0=", result)
	}

	assert.Equal(t, "eyJuYW1lIjoiTWF0In0=", o.MustBase64())

}

func TestConversionBase64WithError(t *testing.T) {

	o := new(O)
	o.obj = func() {}

	assert.Panics(t, func() {
		o.MustBase64()
	})

	_, err := o.Base64()

	assert.Error(t, err)

}
