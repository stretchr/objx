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

func TestNewMSI(t *testing.T) {

	o := MSI("name", "Tyler")

	if assert.NotNil(t, o) {
		assert.Equal(t, o.Get("name").Str(), "Tyler")
	}

}

func TesNewtMustFromJSONWithError(t *testing.T) {

	_, err := FromJSON(`"name":"Mat"}`)
	assert.Error(t, err)

}

func TestNewFromJSON(t *testing.T) {

	o := MustFromJSON(`{"name":"Mat"}`)

	if assert.NotNil(t, o) {
		if assert.NotNil(t, o.Obj()) {
			assert.Equal(t, "Mat", o.Obj().(map[string]interface{})["name"])
		}
	}

}

func TestNewFromJSONWithError(t *testing.T) {

	var o *Obj

	assert.Panics(t, func() {
		o = MustFromJSON(`"name":"Mat"}`)
	})

	assert.Nil(t, o)

}

func TestNewFromBase64String(t *testing.T) {

	base64String := "eyJuYW1lIjoiTWF0In0="

	o, err := FromBase64(base64String)

	if assert.NoError(t, err) {
		assert.Equal(t, o.Get("name").Str(), "Mat")
	}

	assert.Equal(t, MustFromBase64(base64String).Get("name").Str(), "Mat")

}

func TestNewFromBase64StringWithError(t *testing.T) {

	base64String := "eyJuYW1lIjoiTWFasd0In0="

	_, err := FromBase64(base64String)

	assert.Error(t, err)

	assert.Panics(t, func() {
		MustFromBase64(base64String)
	})

}

func TestNewFromSignedBase64String(t *testing.T) {

	base64String := "eyJuYW1lIjoiTWF0In0=_67ee82916f90b2c0d68c903266e8998c9ef0c3d6"

	o, err := FromSignedBase64(base64String, "key")

	if assert.NoError(t, err) {
		assert.Equal(t, o.Get("name").Str(), "Mat")
	}

	assert.Equal(t, MustFromSignedBase64(base64String, "key").Get("name").Str(), "Mat")

}

func TestNewFromSignedBase64StringWithError(t *testing.T) {

	base64String := "eyJuYW1lasdIjoiTWF0In0=_67ee82916f90b2c0d68c903266e8998c9ef0c3d6"

	_, err := FromSignedBase64(base64String, "key")

	assert.Error(t, err)

	assert.Panics(t, func() {
		MustFromSignedBase64(base64String, "key")
	})

}

func TestNewFromURLQuery(t *testing.T) {

	m, err := FromURLQuery("name=tyler&state=UT")
	if assert.NoError(t, err) && assert.NotNil(t, m) {
		assert.Equal(t, "tyler", m.Get("name").Str())
		assert.Equal(t, "UT", m.Get("state").Str())
	}

}
