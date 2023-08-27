package objx

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Regex to parse an array index access
var parseIndexRegex = regexp.MustCompile(`\[([\d]+)\]`)

type notFoundError struct{}

func (m *notFoundError) Error() string {
	return "NotFound"
}

// Get gets the value using the specified selector and
// returns it inside a new Obj object.
//
// If it cannot find the value, Get will return a nil
// value inside an instance of Obj.
//
// Get can only operate directly on map[string]interface{} and []interface.
//
// # Example
//
// To access the title of the third chapter of the second book, do:
//
//	o.Get("books[1].chapters[2].title")
func (m Map) Get(selector string) *Value {
	rawObj, _ := access(m, selector, nil, false)
	return &Value{data: rawObj}
}

// Set sets the value using the specified selector and
// returns the object on which Set was called.
//
// Set can only operate directly on map[string]interface{} and []interface
//
// # Example
//
// To set the title of the third chapter of the second book, do:
//
//	o.Set("books[1].chapters[2].title","Time to Go")
func (m Map) Set(selector string, value interface{}) Map {
	var newObj reflect.Value
	if value == nil {
		newObj = reflect.ValueOf(&value).Elem()
	} else {
		newObj = reflect.ValueOf(value)
	}
	access(m, selector, &newObj, true)
	return m
}

func parsePath(path string) ([]string, error) {
	res := make([]string, 0, 8)
	path = strings.TrimPrefix(path, ".")

	for {
		pos := strings.IndexAny(path, ".[")

		if pos == 0 && path[pos] == '[' {
			pos = strings.IndexAny(path[1:], "[]")
			if pos < 0 || path[pos+1] == '[' {
				return nil, fmt.Errorf("invalid path")
			}
			pos += 2
		}
		var elem string
		if pos < 0 {
			elem = path
		} else {
			elem = path[:pos]
		}

		if elem[0] == '[' {
			if !parseIndexRegex.MatchString(elem) {
				// its not an index so drop the backets for normal access
				if len(elem) <= 2 {
					return nil, fmt.Errorf("invalid path")
				}
				elem = elem[1 : len(elem)-1]
			}
		}

		res = append(res, elem)

		if pos < 0 || pos >= len(path) {
			break
		}

		if path[pos] == '.' {
			pos++
		}
		path = path[pos:]
	}
	return res, nil
}

func getArrayIndex(key string) int {
	if key[0] == '[' {
		idx, err := strconv.ParseUint(key[1:len(key)-1], 10, 64)
		if err != nil {
			// should not happen, otherwise the path is invalid
			panic(err)
		}
		return int(idx)
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func access(object interface{}, selector string, value *reflect.Value, createPath bool) (interface{}, error) {
	path, err := parsePath(selector)
	if err != nil {
		return nil, err
	}

	length := len(path)
	lastIndex := length - 1

	currentObj := reflect.ValueOf(object)
	for index := 0; index < length && currentObj.IsValid(); index++ {
		key := path[index]
		arrayIndex := getArrayIndex(key)
		var nextObj reflect.Value
		// fmt.Printf("current: %s %v\n", key, currentObj.Kind())
		if arrayIndex >= 0 {
			if currentObj.Kind() != reflect.Slice || currentObj.Len() <= arrayIndex {
				if createPath {
					return nil, fmt.Errorf("set with invalid type. Expected currentObj to be a Slice(Len > %d) but got %v(Len %d)", arrayIndex, currentObj.Kind(), currentObj.Len())
				}
				return nil, &notFoundError{}
			}
			nextObj = currentObj.Index(arrayIndex)
		} else {
			if currentObj.Kind() != reflect.Map {
				if createPath {
					return nil, fmt.Errorf("set with invalid type. Expected currentObj to be a Map but got %v", currentObj.Kind())
				}
				return nil, &notFoundError{}
			}
			nextObj = currentObj.MapIndex(reflect.ValueOf(key))
			if !createPath && !nextObj.IsValid() {
				return nil, &notFoundError{}
			}
		}
		if nextObj.Kind() == reflect.Interface {
			nextObj = nextObj.Elem()
		}
		// fmt.Printf("key: %s %v\n", key, nextObj.Kind())

		newObj := nextObj
		if index == lastIndex && value != nil {
			// we are in the last path, assign the value
			newObj = *value
		} else {
			if createPath {
				nextArrayIndex := getArrayIndex(path[index+1])
				if nextArrayIndex >= 0 {
					// next element will be an array
					if !nextObj.IsValid() || nextObj.Kind() != reflect.Slice {
						// fmt.Printf("new slice for %s\n", key)
						newObj = reflect.ValueOf(make([]interface{}, nextArrayIndex+1, max(nextArrayIndex+1, 8)))
					} else if nextObj.Len() <= nextArrayIndex {
						// fmt.Printf("nextObj %s len %d cap %d\n", key, nextObj.Len(), nextObj.Cap())
						newObj = reflect.AppendSlice(nextObj, reflect.ValueOf(make([]interface{}, nextArrayIndex-nextObj.Len()+1)))
					}
				} else {
					// next element will be an normal object
					if !nextObj.IsValid() || nextObj.Kind() != reflect.Map {
						// fmt.Printf("new map for %s\n", key)
						newObj = reflect.ValueOf(map[string]interface{}{})
					}
				}
			}
		}

		if newObj != nextObj || (value != nil && index == lastIndex) {
			// fmt.Printf("assign key %s to %v\n", key, newObj)
			if arrayIndex >= 0 {
				if newObj.IsValid() {
					currentObj.Index(arrayIndex).Set(newObj)
				} else {
					// delete op
					// TODO: implement array shrinking if its the last element
					//if currentObj.Len() == arrayIndex+1 {
					// val := currentObj.Slice(0, arrayIndex)
					// currentObj.Set(val)
					var val interface{}
					currentObj.Index(arrayIndex).Set(reflect.ValueOf(&val).Elem())
					currentObj = nextObj
					break
				}
				nextObj = newObj
			} else {
				currentObj.SetMapIndex(reflect.ValueOf(key), newObj)
				if !newObj.IsValid() {
					// delete op
					currentObj = nextObj
					break
				}
				nextObj = newObj
			}

		}

		currentObj = nextObj
	}
	if !currentObj.IsValid() {
		// JSON NULL
		return nil, nil
	}
	return currentObj.Interface(), nil
}
