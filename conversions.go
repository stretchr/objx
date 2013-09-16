package objx

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

// JSON converts the contained object to a JSON string
// representation
func (o *Map) JSON() (string, error) {

	result, err := json.Marshal(o.value.data)

	if err != nil {
		err = errors.New("objx: JSON encode failed with: " + err.Error())
	}

	return string(result), err

}

// MustJSON converts the contained object to a JSON string
// representation and panics if there is an error
func (o *Map) MustJSON() string {
	result, err := o.JSON()
	if err != nil {
		panic(err.Error())
	}
	return result
}

// Base64 converts the contained object to a Base64 string
// representation of the JSON string representation
func (o *Map) Base64() (string, error) {

	var buf bytes.Buffer

	jsonData, err := o.JSON()
	if err != nil {
		return "", err
	}

	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	encoder.Write([]byte(jsonData))
	encoder.Close()

	return buf.String(), nil

}

// MustBase64 converts the contained object to a Base64 string
// representation of the JSON string representation and panics
// if there is an error
func (o *Map) MustBase64() string {
	result, err := o.Base64()
	if err != nil {
		panic(err.Error())
	}
	return result
}

// SignedBase64 converts the contained object to a Base64 string
// representation of the JSON string representation and signs it
// using the provided key.
func (o *Map) SignedBase64(key string) (string, error) {

	base64, err := o.Base64()
	if err != nil {
		return "", err
	}

	sig := HashWithKey(base64, key)

	return base64 + SignatureSeparator + sig, nil

}

// MustSignedBase64 converts the contained object to a Base64 string
// representation of the JSON string representation and signs it
// using the provided key and panics if there is an error
func (o *Map) MustSignedBase64(key string) string {
	result, err := o.SignedBase64(key)
	if err != nil {
		panic(err.Error())
	}
	return result
}

/*
	URL Query
	------------------------------------------------
*/

// URLValues creates a url.Values object from an Obj. This
// function requires that the wrapped object be a map[string]interface{}
func (o *Map) URLValues() url.Values {

	vals := make(url.Values)

	for k, v := range o.value.MSI() {
		//TODO: can this be done without sprintf?
		vals.Set(k, fmt.Sprintf("%v", v))
	}

	return vals
}

// URLQuery gets an encoded URL query representing the given
// Obj. This function requires that the wrapped object be a
// map[string]interface{}
func (o *Map) URLQuery() (string, error) {
	return o.URLValues().Encode(), nil
}
