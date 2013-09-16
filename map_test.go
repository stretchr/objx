package objx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapCreation(t *testing.T) {

	o := new(Map)
	if assert.NotNil(t, o) {
		assert.Nil(t, o.value)
	}

	o = New(nil)
	assert.Nil(t, o)

	o = New("Tyler")
	assert.Nil(t, o)

	o = MSI()
	if assert.NotNil(t, o) {
		assert.NotNil(t, o.value.data)
	}

	o = MSI("name", "Tyler")
	if assert.NotNil(t, o) {
		if assert.NotNil(t, o.value.data) {
			assert.Equal(t, o.value.data.(map[string]interface{})["name"], "Tyler")
		}
	}

}

func TestMapMustFromJSONWithError(t *testing.T) {

	_, err := FromJSON(`"name":"Mat"}`)
	assert.Error(t, err)

}

func TestMapFromJSON(t *testing.T) {

	o := MustFromJSON(`{"name":"Mat"}`)

	if assert.NotNil(t, o) {
		if assert.NotNil(t, o.value.data) {
			assert.Equal(t, "Mat", o.value.data.(map[string]interface{})["name"])
		}
	}

}

func TestMapFromJSONWithError(t *testing.T) {

	var m *Map

	assert.Panics(t, func() {
		m = MustFromJSON(`"name":"Mat"}`)
	})

	assert.Nil(t, m)

}

func TestMapFromBase64String(t *testing.T) {

	base64String := "eyJuYW1lIjoiTWF0In0="

	o, err := FromBase64(base64String)

	if assert.NoError(t, err) {
		assert.Equal(t, o.Get("name").Str(), "Mat")
	}

	assert.Equal(t, MustFromBase64(base64String).Get("name").Str(), "Mat")

}

func TestMapFromBase64StringWithError(t *testing.T) {

	base64String := "eyJuYW1lIjoiTWFasd0In0="

	_, err := FromBase64(base64String)

	assert.Error(t, err)

	assert.Panics(t, func() {
		MustFromBase64(base64String)
	})

}

func TestMapFromSignedBase64String(t *testing.T) {

	base64String := "eyJuYW1lIjoiTWF0In0=_67ee82916f90b2c0d68c903266e8998c9ef0c3d6"

	o, err := FromSignedBase64(base64String, "key")

	if assert.NoError(t, err) {
		assert.Equal(t, o.Get("name").Str(), "Mat")
	}

	assert.Equal(t, MustFromSignedBase64(base64String, "key").Get("name").Str(), "Mat")

}

func TestMapFromSignedBase64StringWithError(t *testing.T) {

	base64String := "eyJuYW1lasdIjoiTWF0In0=_67ee82916f90b2c0d68c903266e8998c9ef0c3d6"

	_, err := FromSignedBase64(base64String, "key")

	assert.Error(t, err)

	assert.Panics(t, func() {
		MustFromSignedBase64(base64String, "key")
	})

}

func TestMapFromURLQuery(t *testing.T) {

	m, err := FromURLQuery("name=tyler&state=UT")
	if assert.NoError(t, err) && assert.NotNil(t, m) {
		assert.Equal(t, "tyler", m.Get("name").Str())
		assert.Equal(t, "UT", m.Get("state").Str())
	}

}
