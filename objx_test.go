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

func TestObjxMustFromJSONWithError(t *testing.T) {

	_, err := FromJSON(`"name":"Mat"}`)
	assert.Error(t, err)

}

func TestObjxFromJSON(t *testing.T) {

	o := MustFromJSON(`{"name":"Mat"}`)

	if assert.NotNil(t, o) {
		if assert.NotNil(t, o.data) {
			assert.Equal(t, "Mat", o.data.(map[string]interface{})["name"])
		}
	}

}

func TestObjxFromJSONWithError(t *testing.T) {

	var o *Objx

	assert.Panics(t, func() {
		o = MustFromJSON(`"name":"Mat"}`)
	})

	assert.Nil(t, o)

}

func TestObjxFromBase64String(t *testing.T) {

	base64String := "eyJuYW1lIjoiTWF0In0="

	o, err := FromBase64(base64String)

	if assert.NoError(t, err) {
		assert.Equal(t, o.Get("name").Str(), "Mat")
	}

	assert.Equal(t, MustFromBase64(base64String).Get("name").Str(), "Mat")

}

func TestObjxFromBase64StringWithError(t *testing.T) {

	base64String := "eyJuYW1lIjoiTWFasd0In0="

	_, err := FromBase64(base64String)

	assert.Error(t, err)

	assert.Panics(t, func() {
		MustFromBase64(base64String)
	})

}

func TestObjxFromSignedBase64String(t *testing.T) {

	base64String := "eyJuYW1lIjoiTWF0In0=_67ee82916f90b2c0d68c903266e8998c9ef0c3d6"

	o, err := FromSignedBase64(base64String, "key")

	if assert.NoError(t, err) {
		assert.Equal(t, o.Get("name").Str(), "Mat")
	}

	assert.Equal(t, MustFromSignedBase64(base64String, "key").Get("name").Str(), "Mat")

}

func TestObjxFromSignedBase64StringWithError(t *testing.T) {

	base64String := "eyJuYW1lasdIjoiTWF0In0=_67ee82916f90b2c0d68c903266e8998c9ef0c3d6"

	_, err := FromSignedBase64(base64String, "key")

	assert.Error(t, err)

	assert.Panics(t, func() {
		MustFromSignedBase64(base64String, "key")
	})

}

func TestObjxFromURLQuery(t *testing.T) {

	m, err := FromURLQuery("name=tyler&state=UT")
	if assert.NoError(t, err) && assert.NotNil(t, m) {
		assert.Equal(t, "tyler", m.Get("name").Str())
		assert.Equal(t, "UT", m.Get("state").Str())
	}

}
