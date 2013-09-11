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

// access accesses the object using the selector and performs the
// following actions:
// 1. if value is nil, retrieve the object at selector
// 2. if value is not nil, set the selector to value
func (o *O) access(selector interface{}, value interface{}, isSet bool) *O {

	switch selector.(type) {
	case string: // "address.postcode.inner"

		selStr := selector.(string)
		segs := strings.Split(selStr, PathSeparator)
		segLen := len(segs)

		current := o.Obj()

		for index, field := range segs {

			arrayMatches := arrayAccesRegex.FindStringSubmatch(field)

			if len(arrayMatches) > 0 {
				// An array notation has been detected.

				// Get the key into the map
				mName := arrayMatches[1]

				// Get the index into the array at the key
				index, err := strconv.Atoi(arrayMatches[2])
				if err != nil {
					// This should never happen. If it does, something has gone
					// seriously wrong. Panic.
					panic("objx: access - array index is not an integer. This should never happen.")
				}

				if m, ok := current.(map[string]interface{}); ok {
					arrayObj := o.arrayAccess(m[mName], uint64(index), value, isSet)
					if segLen-1 == index && isSet {
						return arrayObj
					} else {
						current = arrayObj.obj
					}
				}
			} else {
				if m, ok := current.(map[string]interface{}); ok {
					if segLen-1 == index && isSet {
						m[field] = value
						return o
					} else {
						current = m[field]
					}
				}
			}

		}

		return New(current)

	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		selInt := uint64FromInterface(selector)
		return o.arrayAccess(o.obj, selInt, value, isSet)

	}

	panic("objx: access with invalid selector.")

}

func (o *O) arrayAccess(object interface{}, index uint64, value interface{}, isSet bool) *O {

	switch object.(type) {
	case []bool:
		a := object.([]bool)
		if isSet {
			a[index] = value.(bool)
			return o
		} else {
			return New(a[index])
		}
	case []int:
		a := object.([]int)
		if isSet {
			a[index] = value.(int)
			return o
		} else {
			return New(a[index])
		}
	case []int8:
		a := object.([]int8)
		if isSet {
			a[index] = value.(int8)
			return o
		} else {
			return New(a[index])
		}
	case []int16:
		a := object.([]int16)
		if isSet {
			a[index] = value.(int16)
			return o
		} else {
			return New(a[index])
		}
	case []int32:
		a := object.([]int32)
		if isSet {
			a[index] = value.(int32)
			return o
		} else {
			return New(a[index])
		}
	case []int64:
		a := object.([]int64)
		if isSet {
			a[index] = value.(int64)
			return o
		} else {
			return New(a[index])
		}
	case []uint:
		a := object.([]uint)
		if isSet {
			a[index] = value.(uint)
			return o
		} else {
			return New(a[index])
		}
	case []uint8:
		a := object.([]uint8)
		if isSet {
			a[index] = value.(uint8)
			return o
		} else {
			return New(a[index])
		}
	case []uint16:
		a := object.([]uint16)
		if isSet {
			a[index] = value.(uint16)
			return o
		} else {
			return New(a[index])
		}
	case []uint32:
		a := object.([]uint32)
		if isSet {
			a[index] = value.(uint32)
			return o
		} else {
			return New(a[index])
		}
	case []uint64:
		a := object.([]uint64)
		if isSet {
			a[index] = value.(uint64)
			return o
		} else {
			return New(a[index])
		}
	case []float32:
		a := object.([]float32)
		if isSet {
			a[index] = value.(float32)
			return o
		} else {
			return New(a[index])
		}
	case []float64:
		a := object.([]float64)
		if isSet {
			a[index] = value.(float64)
			return o
		} else {
			return New(a[index])
		}
	case []complex64:
		a := object.([]complex64)
		if isSet {
			a[index] = value.(complex64)
			return o
		} else {
			return New(a[index])
		}
	case []complex128:
		a := object.([]complex128)
		if isSet {
			a[index] = value.(complex128)
			return o
		} else {
			return New(a[index])
		}
	case []string:
		a := object.([]string)
		if isSet {
			a[index] = value.(string)
			return o
		} else {
			return New(a[index])
		}
	case []uintptr:
		a := object.([]uintptr)
		if isSet {
			a[index] = value.(uintptr)
			return o
		} else {
			return New(a[index])
		}
	case []interface{}:
		a := object.([]interface{})
		if isSet {
			a[index] = value.(interface{})
			return o
		} else {
			return New(a[index])
		}
	default:
		// let reflect deal with this
		panic("objx: access attempted to index into a non-array object")
	}

}

// uint64FromInterface converts an interface object to the largest
// representation of an unsigned integer using a type switch and
// assertions
func uint64FromInterface(selector interface{}) uint64 {
	var value uint64
	switch selector.(type) {
	case int:
		value = uint64(selector.(int))
	case int8:
		value = uint64(selector.(int8))
	case int16:
		value = uint64(selector.(int16))
	case int32:
		value = uint64(selector.(int32))
	case int64:
		value = uint64(selector.(int64))
	case uint:
		value = uint64(selector.(uint))
	case uint8:
		value = uint64(selector.(uint8))
	case uint16:
		value = uint64(selector.(uint16))
	case uint32:
		value = uint64(selector.(uint32))
	case uint64:
		value = selector.(uint64)
	default:
		panic("objx: array access argument is not an integer (this should never happen)")
	}

	return value
}
