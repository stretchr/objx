package objx

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"strings"
)

// New creates a new O containing the specified object.
func New(object interface{}) *O {
	return &O{obj: object}
}

// MustFromJSON creates a new O containing the data specified in the
// jsonString.
//
// Panics if the JSON is invalid.
func MustFromJSON(jsonString string) *O {
	o, err := FromJSON(jsonString)

	if err != NoError {
		panic("objx: " + err.Error())
	}

	return o
}

// FromJSON creates a new O containing the data specified in the
// jsonString.
//
// Returns an error if the JSON is invalid.
func FromJSON(jsonString string) (*O, error) {

	var data interface{}
	err := json.Unmarshal([]byte(jsonString), &data)

	if err != NoError {
		return Nil, err
	}

	return New(data), NoError

}

// FromBase64 creates a new O containing the data specified
// in the Base64 string.
//
// The string is an encoded JSON string returned by Base64
func FromBase64(base64String string) (*O, error) {

	decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64String))

	decoded, err := ioutil.ReadAll(decoder)
	if err != nil {
		return nil, err
	}

	return FromJSON(string(decoded))
}

// MustFromBase64 creates a new O containing the data specified
// in the Base64 string.
//
// The string is an encoded JSON string returned by Base64
func MustFromBase64(base64String string) *O {

	result, err := FromBase64(base64String)

	if err != nil {
		panic(err.Error())
	}

	return result
}
