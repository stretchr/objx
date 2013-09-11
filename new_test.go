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

func TestFromBase64String(t *testing.T) {

	base64String := "eyJuYW1lIjoiTWF0In0="

	o, err := FromBase64(base64String)

	if assert.NoError(t, err) {
		assert.Equal(t, o.Get("name").Str(), "Mat")
	}

	assert.Equal(t, MustFromBase64(base64String).Get("name").Str(), "Mat")

}

func TestFromBase64StringWithError(t *testing.T) {

	base64String := "eyJuYW1lIjoiTWFasd0In0="

	_, err := FromBase64(base64String)

	assert.Error(t, err)

	assert.Panics(t, func() {
		MustFromBase64(base64String)
	})

}

func TestFromSignedBase64String(t *testing.T) {

	base64String := "eyJuYW1lIjoiTWF0In0=_67ee82916f90b2c0d68c903266e8998c9ef0c3d6"

	o, err := FromSignedBase64(base64String, "key")

	if assert.NoError(t, err) {
		assert.Equal(t, o.Get("name").Str(), "Mat")
	}

	assert.Equal(t, MustFromSignedBase64(base64String, "key").Get("name").Str(), "Mat")

}

func TestFromSignedBase64StringWithError(t *testing.T) {

	base64String := "eyJuYW1lasdIjoiTWF0In0=_67ee82916f90b2c0d68c903266e8998c9ef0c3d6"

	_, err := FromSignedBase64(base64String, "key")

	assert.Error(t, err)

	assert.Panics(t, func() {
		MustFromSignedBase64(base64String, "key")
	})

}
