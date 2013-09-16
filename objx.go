package objx

import (
	"fmt"
)

// Objx provides extended functionality for working with
// untyped data, particularly map[string]interface{},
// []interface{} and interface{}.
type Objx struct {
	// data is the raw data being managed by this Objx
	data interface{}
}

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
			panic(fmt.Sprintf("objx: MSI must follow 'string, interface{}' pattern.  %s is not a valid key.", keyString))
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
