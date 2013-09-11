package objx

import (
	"encoding/json"
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
