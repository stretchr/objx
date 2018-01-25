package objx

import (
	"regexp"
	"strconv"
	"strings"
)

// arrayAccesRegexString is the regex used to extract the array number
// from the access path
const arrayAccesRegexString = `^(.+)\[([0-9]+)\]$`

// arrayAccesRegex is the compiled arrayAccesRegexString
var arrayAccesRegex = regexp.MustCompile(arrayAccesRegexString)

// Get gets the value using the specified selector and
// returns it inside a new Obj object.
//
// If it cannot find the value, Get will return a nil
// value inside an instance of Obj.
//
// Get can only operate directly on map[string]interface{} and []interface.
//
// Example
//
// To access the title of the third chapter of the second book, do:
//
//    o.Get("books[1].chapters[2].title")
func (m Map) Get(selector string) *Value {
	rawObj := access(m, selector, nil, false)
	return &Value{data: rawObj}
}

// Set sets the value using the specified selector and
// returns the object on which Set was called.
//
// Set can only operate directly on map[string]interface{} and []interface
//
// Example
//
// To set the title of the third chapter of the second book, do:
//
//    o.Set("books[1].chapters[2].title","Time to Go")
func (m Map) Set(selector string, value interface{}) Map {
	access(m, selector, value, true)
	return m
}

// access accesses the object using the selector and performs the
// appropriate action.
func access(current interface{}, selector string, value interface{}, isSet bool) interface{} {
	selSegs := strings.SplitN(selector, PathSeparator, 2)
	thisSel := selSegs[0]
	index := -1
	var err error

	if strings.Contains(thisSel, "[") {
		arrayMatches := arrayAccesRegex.FindStringSubmatch(thisSel)
		if len(arrayMatches) > 0 {
			// Get the key into the map
			thisSel = arrayMatches[1]

			// Get the index into the array at the key
			index, err = strconv.Atoi(arrayMatches[2])

			if err != nil {
				// This should never happen. If it does, something has gone
				// seriously wrong. Panic.
				panic("objx: Array index is not an integer.  Must use array[int].")
			}
		}
	}
	if curMap, ok := current.(Map); ok {
		current = map[string]interface{}(curMap)
	}
	// get the object in question
	switch current.(type) {
	case map[string]interface{}:
		curMSI := current.(map[string]interface{})
		if len(selSegs) <= 1 && isSet {
			curMSI[thisSel] = value
			return nil
		}
		current = curMSI[thisSel]
	default:
		current = nil
	}
	// do we need to access the item of an array?
	if index > -1 {
		if array, ok := current.([]interface{}); ok {
			if index < len(array) {
				current = array[index]
			} else {
				current = nil
			}
		}
	}
	if len(selSegs) > 1 {
		current = access(current, selSegs[1], value, isSet)
	}
	return current
}
