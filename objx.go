package objx

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
	"strings"
)

// Objx provides extended functionality for working with
// untyped data, particularly map[string]interface{},
// []interface{} and interface{}.
type Objx struct {
	// data is the raw data being managed by this Objx
	data interface{}
}

// Nil represents a nil Objx.
var Nil *Objx = New(nil)

// New creates a new *Objx containing the data
// in the data argument.
func New(data interface{}) *Objx {
	return &Objx{data: data}
}

// MSI creates a map[string]interface{} and puts it inside a new Objx.
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
//     // creates an Objx equivalent to
//     m := objx.New(map[string]interface{}{"name": "Mat", "age": 29, "subobj": map[string]interface{}{"active": true}})
func MSI(keyAndValuePairs ...interface{}) *Objx {

	newMap := make(map[string]interface{})
	keyAndValuePairsLen := len(keyAndValuePairs)

	if keyAndValuePairsLen%2 != 0 {
		panic("objx: MSI must have an even number of arguments following the 'key, value' pattern.")
	}

	for i := 0; i < keyAndValuePairsLen; i = i + 2 {

		key := keyAndValuePairs[i]
		value := keyAndValuePairs[i+1]

		// make sure the key is a string
		keyString, keyStringOK := key.(string)
		if !keyStringOK {
			panic("objx: MSI must follow 'string, interface{}' pattern.  " + keyString + " is not a valid key.")
		}

		newMap[keyString] = value

	}

	return New(newMap)
}

// Slice creates an []interface{} and puts it inside an Objx.
// Arguments are a comma separated list of objects.
func Slice(objects ...interface{}) *Objx {
	return New(objects)
}

// ****** Conversion Constructors

// MustFromJSON creates a new Objx containing the data specified in the
// jsonString.
//
// Panics if the JSON is invalid.
func MustFromJSON(jsonString string) *Objx {
	o, err := FromJSON(jsonString)

	if err != nil {
		panic("objx: MustFromJSON failed with error: " + err.Error())
	}

	return o
}

// FromJSON creates a new Objx containing the data specified in the
// jsonString.
//
// Returns an error if the JSON is invalid.
func FromJSON(jsonString string) (*Objx, error) {

	var data interface{}
	err := json.Unmarshal([]byte(jsonString), &data)

	if err != nil {
		return Nil, err
	}

	return New(data), nil

}

// FromBase64 creates a new Obj containing the data specified
// in the Base64 string.
//
// The string is an encoded JSON string returned by Base64
func FromBase64(base64String string) (*Objx, error) {

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
func MustFromBase64(base64String string) *Objx {

	result, err := FromBase64(base64String)

	if err != nil {
		panic("objx: MustFromBase64 failed with error: " + err.Error())
	}

	return result
}

// FromSignedBase64 creates a new Obj containing the data specified
// in the Base64 string.
//
// The string is an encoded JSON string returned by SignedBase64
func FromSignedBase64(base64String, key string) (*Objx, error) {
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
func MustFromSignedBase64(base64String, key string) *Objx {

	result, err := FromSignedBase64(base64String, key)

	if err != nil {
		panic("objx: MustFromSignedBase64 failed with error: " + err.Error())
	}

	return result
}

// FromURLQuery generates a new Obj by parsing the specified
// query.
//
// For queries with multiple values, the first value is selected.
func FromURLQuery(query string) (*Objx, error) {

	vals, err := url.ParseQuery(query)

	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	for k, vals := range vals {
		m[k] = vals[0]
	}

	return New(m), nil
}

// MustFromURLQuery generates a new Obj by parsing the specified
// query.
//
// For queries with multiple values, the first value is selected.
//
// Panics if it encounters an error
func MustFromURLQuery(query string) *Objx {

	o, err := FromURLQuery(query)

	if err != nil {
		panic("objx: MustFromURLQuery failed with error: " + err.Error())
	}

	return o

}
