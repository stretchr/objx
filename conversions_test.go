package objx_test

import (
	"testing"

	"github.com/stretchr/objx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConversionJSON(t *testing.T) {
	jsonString := `{"name":"Mat"}`
	o := objx.MustFromJSON(jsonString)

	result, err := o.JSON()

	require.NoError(t, err)
	assert.Equal(t, jsonString, result)
	assert.Equal(t, jsonString, o.MustJSON())
}

func TestConversionJSONWithError(t *testing.T) {
	o := objx.MSI()
	o["test"] = func() {}

	assert.Panics(t, func() {
		o.MustJSON()
	})

	_, err := o.JSON()

	assert.Error(t, err)
}

func TestConversionBase64(t *testing.T) {
	o := objx.Map{"name": "Mat"}

	result, err := o.Base64()

	require.NoError(t, err)
	assert.Equal(t, "eyJuYW1lIjoiTWF0In0=", result)
	assert.Equal(t, "eyJuYW1lIjoiTWF0In0=", o.MustBase64())
}

func TestConversionBase64WithError(t *testing.T) {
	o := objx.MSI()
	o["test"] = func() {}

	assert.Panics(t, func() {
		o.MustBase64()
	})

	_, err := o.Base64()

	assert.Error(t, err)
}

func TestConversionSignedBase64(t *testing.T) {
	o := objx.Map{"name": "Mat"}

	result, err := o.SignedBase64("key")

	require.NoError(t, err)
	assert.Equal(t, "eyJuYW1lIjoiTWF0In0=_67ee82916f90b2c0d68c903266e8998c9ef0c3d6", result)
	assert.Equal(t, "eyJuYW1lIjoiTWF0In0=_67ee82916f90b2c0d68c903266e8998c9ef0c3d6", o.MustSignedBase64("key"))
}

func TestConversionSignedBase64WithError(t *testing.T) {
	o := objx.MSI()
	o["test"] = func() {}

	assert.Panics(t, func() {
		o.MustSignedBase64("key")
	})

	_, err := o.SignedBase64("key")

	assert.Error(t, err)
}
