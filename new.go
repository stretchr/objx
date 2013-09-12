package objx

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

// New creates a new Obj containing the specified object.
func New(object interface{}) *Obj {
	return &Obj{obj: object}
}

// MSI creates a map[string]interface{} and puts it inside a new Obj.
//
// The arguments follow a key, value pattern.
//
// Panics
//
// Panics if any key arugment is non-string or if there are an odd number of arguments.
//
// Example
//
// To easily create Maps:
//
//     m := objx.MSI("name", "Mat", "age", 29, "subobj", objx.MSI("active", true))
//
//     // creates an Obj equivalent to
//     m := objx.New(map[string]interface{}{"name": "Mat", "age": 29, "subobj": map[string]interface{}{"active": true}})
func MSI(keyAndValuePairs ...interface{}) *Obj {

	newMap := make(map[string]interface{})
	keyAndValuePairsLen := len(keyAndValuePairs)

	if keyAndValuePairsLen%2 != 0 {
		panic("MSI must have an even number of arguments following the 'key, value' pattern.")
	}

	for i := 0; i < keyAndValuePairsLen; i = i + 2 {

		key := keyAndValuePairs[i]
		value := keyAndValuePairs[i+1]

		// make sure the key is a string
		keyString, keyStringOK := key.(string)
		if !keyStringOK {
			panic(fmt.Sprintf("MSI must follow 'string, interface{}' pattern.  %s is not a valid key.", keyString))
		}

		newMap[keyString] = value

	}

	return New(newMap)
}

// MustFromJSON creates a new Obj containing the data specified in the
// jsonString.
//
// Panics if the JSON is invalid.
func MustFromJSON(jsonString string) *Obj {
	o, err := FromJSON(jsonString)

	if err != NoError {
		panic("objx: " + err.Error())
	}

	return o
}

// FromJSON creates a new Obj containing the data specified in the
// jsonString.
//
// Returns an error if the JSON is invalid.
func FromJSON(jsonString string) (*Obj, error) {

	var data interface{}
	err := json.Unmarshal([]byte(jsonString), &data)

	if err != NoError {
		return Nil, err
	}

	return New(data), NoError

}

// FromBase64 creates a new Obj containing the data specified
// in the Base64 string.
//
// The string is an encoded JSON string returned by Base64
func FromBase64(base64String string) (*Obj, error) {

	decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64String))

	decoded, err := ioutil.ReadAll(decoder)
	if err != nil {
		return nil, err
	}

	return FromJSON(string(decoded))
}

// MustFromBase64 creates a new Obj containing the data specified
// in the Base64 string and panics if there is an error.
//
// The string is an encoded JSON string returned by Base64
func MustFromBase64(base64String string) *Obj {

	result, err := FromBase64(base64String)

	if err != nil {
		panic(err.Error())
	}

	return result
}

// FromSignedBase64 creates a new Obj containing the data specified
// in the Base64 string.
//
// The string is an encoded JSON string returned by SignedBase64
func FromSignedBase64(base64String, key string) (*Obj, error) {
	parts := strings.Split(base64String, SignatureSeparator)
	if len(parts) != 2 {
		return nil, errors.New("objx: Signed base64 string is malformed.")
	}

	sig := HashWithKey(parts[0], key)
	if parts[1] != sig {
		return nil, errors.New("objx: Signature for base64 data does not match.")
	}

	return FromBase64(parts[0])
}

// MustFromSignedBase64 creates a new Obj containing the data specified
// in the Base64 string and panics if there is an error.
//
// The string is an encoded JSON string returned by Base64
func MustFromSignedBase64(base64String, key string) *Obj {

	result, err := FromSignedBase64(base64String, key)

	if err != nil {
		panic(err.Error())
	}

	return result
}
