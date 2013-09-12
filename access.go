package objx

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// arrayAccesRegexString is the regex used to extract the array number
// from the access path
const arrayAccesRegexString = `^(.+)\[([0-9]+)\]$`

// arrayAccesRegex is the compiled arrayAccesRegexString
var arrayAccesRegex = regexp.MustCompile(arrayAccesRegexString)

// access accesses the object using the selector and performs the
// appropriate action.
func access(current interface{}, selector interface{}, panics bool) interface{} {

	switch selector.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:

		indexInt := intFromInterface(selector)
		current = arrayAccess(current, indexInt, panics)

	case string:

		selStr := selector.(string)
		selSegs := strings.SplitN(selStr, PathSeparator, 2)
		thisSel := selSegs[0]
		arrIndex := -1
		var err error

		arrayMatches := arrayAccesRegex.FindStringSubmatch(thisSel)

		if len(arrayMatches) > 0 {

			// Get the key into the map
			thisSel = arrayMatches[1]

			// Get the index into the array at the key
			arrIndex, err = strconv.Atoi(arrayMatches[2])

			if err != nil {
				// This should never happen. If it does, something has gone
				// seriously wrong. Panic.
				panic("objx: Array index is not an integer.  Must use array[int].")
			}

		}

		// get the object in question
		switch current.(type) {
		case map[string]interface{}:
			current = current.(map[string]interface{})[thisSel]
		default:
			current = nil
		}

		if current == nil && panics {
			panic(fmt.Sprintf("objx: '%v' invalid on object.", selector))
		}

		// do we need to access the item of an array?
		if arrIndex > -1 {

			current = arrayAccess(current, arrIndex, panics)

		}

		if len(selSegs) > 1 {

			current = access(current, selSegs[1], panics)

		}

	}

	return current

}

func arrayAccess(object interface{}, index int, panics bool) interface{} {

	switch object.(type) {
	case []interface{}:
		a := object.([]interface{})
		if index >= len(a) {
			if panics {
				panic(fmt.Sprintf("objx: Index %d is out of range because the []interface{} only contains %d items.", index, len(a)))
			}
			return nil
		} else {
			return a[index]
		}

	case []map[string]interface{}:
		a := object.([]map[string]interface{})
		if index >= len(a) {
			if panics {
				panic(fmt.Sprintf("objx: Index %d is out of range because the []map[string]interface{} only contains %d items.", index, len(a)))
			}
			return nil
		} else {
			return a[index]
		}

	case []*Obj:
		a := object.([]*Obj)
		if index >= len(a) {
			if panics {
				panic(fmt.Sprintf("objx: Index %d is out of range because the []*Obj only contains %d items.", index, len(a)))
			}
			return nil
		} else {
			return a[index]
		}

	case []bool:
		a := object.([]bool)
		if index >= len(a) {
			if panics {
				panic(fmt.Sprintf("objx: Index %d is out of range because the []bool only contains %d items.", index, len(a)))
			}
			return nil
		} else {
			return a[index]
		}

	case []string:
		a := object.([]string)
		if index >= len(a) {
			if panics {
				panic(fmt.Sprintf("objx: Index %d is out of range because the []string only contains %d items.", index, len(a)))
			}
			return nil
		} else {
			return a[index]
		}

	case []int:
		a := object.([]int)
		if index >= len(a) {
			if panics {
				panic(fmt.Sprintf("objx: Index %d is out of range because the []int only contains %d items.", index, len(a)))
			}
			return nil
		} else {
			return a[index]
		}

	case []int8:
		a := object.([]int8)
		if index >= len(a) {
			if panics {
				panic(fmt.Sprintf("objx: Index %d is out of range because the []int8 only contains %d items.", index, len(a)))
			}
			return nil
		} else {
			return a[index]
		}

	case []int16:
		a := object.([]int16)
		if index >= len(a) {
			if panics {
				panic(fmt.Sprintf("objx: Index %d is out of range because the []int16 only contains %d items.", index, len(a)))
			}
			return nil
		} else {
			return a[index]
		}

	case []int32:
		a := object.([]int32)
		if index >= len(a) {
			if panics {
				panic(fmt.Sprintf("objx: Index %d is out of range because the []int32 only contains %d items.", index, len(a)))
			}
			return nil
		} else {
			return a[index]
		}

	case []int64:
		a := object.([]int64)
		if index >= len(a) {
			if panics {
				panic(fmt.Sprintf("objx: Index %d is out of range because the []int64 only contains %d items.", index, len(a)))
			}
			return nil
		} else {
			return a[index]
		}

	case []uint:
		a := object.([]uint)
		if index >= len(a) {
			if panics {
				panic(fmt.Sprintf("objx: Index %d is out of range because the []uint only contains %d items.", index, len(a)))
			}
			return nil
		} else {
			return a[index]
		}

	case []uint8:
		a := object.([]uint8)
		if index >= len(a) {
			if panics {
				panic(fmt.Sprintf("objx: Index %d is out of range because the []uint8 only contains %d items.", index, len(a)))
			}
			return nil
		} else {
			return a[index]
		}

	case []uint16:
		a := object.([]uint16)
		if index >= len(a) {
			if panics {
				panic(fmt.Sprintf("objx: Index %d is out of range because the []uint16 only contains %d items.", index, len(a)))
			}
			return nil
		} else {
			return a[index]
		}

	case []uint32:
		a := object.([]uint32)
		if index >= len(a) {
			if panics {
				panic(fmt.Sprintf("objx: Index %d is out of range because the []uint32 only contains %d items.", index, len(a)))
			}
			return nil
		} else {
			return a[index]
		}

	case []uint64:
		a := object.([]uint64)
		if index >= len(a) {
			if panics {
				panic(fmt.Sprintf("objx: Index %d is out of range because the []uint64 only contains %d items.", index, len(a)))
			}
			return nil
		} else {
			return a[index]
		}

	case []uintptr:
		a := object.([]uintptr)
		if index >= len(a) {
			if panics {
				panic(fmt.Sprintf("objx: Index %d is out of range because the []uintptr only contains %d items.", index, len(a)))
			}
			return nil
		} else {
			return a[index]
		}

	case []float32:
		a := object.([]float32)
		if index >= len(a) {
			if panics {
				panic(fmt.Sprintf("objx: Index %d is out of range because the []float32 only contains %d items.", index, len(a)))
			}
			return nil
		} else {
			return a[index]
		}

	case []float64:
		a := object.([]float64)
		if index >= len(a) {
			if panics {
				panic(fmt.Sprintf("objx: Index %d is out of range because the []float64 only contains %d items.", index, len(a)))
			}
			return nil
		} else {
			return a[index]
		}

	case []complex64:
		a := object.([]complex64)
		if index >= len(a) {
			if panics {
				panic(fmt.Sprintf("objx: Index %d is out of range because the []complex64 only contains %d items.", index, len(a)))
			}
			return nil
		} else {
			return a[index]
		}

	case []complex128:
		a := object.([]complex128)
		if index >= len(a) {
			if panics {
				panic(fmt.Sprintf("objx: Index %d is out of range because the []complex128 only contains %d items.", index, len(a)))
			}
			return nil
		} else {
			return a[index]
		}

	default:
		// let reflect deal with this
		panic(fmt.Sprintf("objx: Cannot index into a non-array type.  %v is not supported.", reflect.TypeOf(object)))
	}

	return nil
}

// intFromInterface converts an interface object to the largest
// representation of an unsigned integer using a type switch and
// assertions
func intFromInterface(selector interface{}) int {
	var value int
	switch selector.(type) {
	case int:
		value = selector.(int)
	case int8:
		value = int(selector.(int8))
	case int16:
		value = int(selector.(int16))
	case int32:
		value = int(selector.(int32))
	case int64:
		value = int(selector.(int64))
	case uint:
		value = int(selector.(uint))
	case uint8:
		value = int(selector.(uint8))
	case uint16:
		value = int(selector.(uint16))
	case uint32:
		value = int(selector.(uint32))
	case uint64:
		value = int(selector.(uint64))
	default:
		panic("objx: array access argument is not an integer type (this should never happen)")
	}

	return value
}
