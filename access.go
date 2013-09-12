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
// following actions:
// 1. if value is nil, retrieve the object at selector
// 2. if value is not nil, set the selector to value
//
// panics indicates whether the function will panic on failure
// or just use a nil value instead.
func (o *Obj) access(selector interface{}, value interface{}, isSet, panics bool) *Obj {

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
					panic("objx: Array index is not an integer.  Must use array[int].")
				}

				if m, ok := current.(map[string]interface{}); ok {
					arrayObj := o.arrayAccess(m[mName], index, value, isSet, panics)
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
						if current, ok = m[field]; !ok && panics {
							panic("objx: '" + field + "' is not defined.")
						}
					}
				} else {
					if panics {
						panic("objx: Cannot get deep within objects of type " + reflect.TypeOf(current).Name())
					} else {
						current = nil
					}
				}
			}

		}

		return New(current)

	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		selInt := intFromInterface(selector)
		return o.arrayAccess(o.obj, selInt, value, isSet, panics)

	}

	panic("objx: access with invalid selector.")

}

func (o *Obj) arrayAccess(object interface{}, index int, value interface{}, isSet, panics bool) *Obj {

	switch object.(type) {
	case []bool:
		a := object.([]bool)
		if isSet {
			a[index] = value.(bool)
			return o
		} else {
			if index >= len(a) {
				if panics {
					panic(fmt.Sprintf("objx: Index %d is out of range because the array only contains %d items.", index, len(a)))
				}
				return Nil
			} else {
				return New(a[index])
			}
		}
	case []int:
		a := object.([]int)
		if isSet {
			a[index] = value.(int)
			return o
		} else {
			if index >= len(a) {
				if panics {
					panic(fmt.Sprintf("objx: Index %d is out of range because the array only contains %d items.", index, len(a)))
				}
				return Nil
			} else {
				return New(a[index])
			}
		}
	case []int8:
		a := object.([]int8)
		if isSet {
			a[index] = value.(int8)
			return o
		} else {
			if index >= len(a) {
				if panics {
					panic(fmt.Sprintf("objx: Index %d is out of range because the array only contains %d items.", index, len(a)))
				}
				return Nil
			} else {
				return New(a[index])
			}
		}
	case []int16:
		a := object.([]int16)
		if isSet {
			a[index] = value.(int16)
			return o
		} else {
			if index >= len(a) {
				if panics {
					panic(fmt.Sprintf("objx: Index %d is out of range because the array only contains %d items.", index, len(a)))
				}
				return Nil
			} else {
				return New(a[index])
			}
		}
	case []int32:
		a := object.([]int32)
		if isSet {
			a[index] = value.(int32)
			return o
		} else {
			if index >= len(a) {
				if panics {
					panic(fmt.Sprintf("objx: Index %d is out of range because the array only contains %d items.", index, len(a)))
				}
				return Nil
			} else {
				return New(a[index])
			}
		}
	case []int64:
		a := object.([]int64)
		if isSet {
			a[index] = value.(int64)
			return o
		} else {
			if index >= len(a) {
				if panics {
					panic(fmt.Sprintf("objx: Index %d is out of range because the array only contains %d items.", index, len(a)))
				}
				return Nil
			} else {
				return New(a[index])
			}
		}
	case []uint:
		a := object.([]uint)
		if isSet {
			a[index] = value.(uint)
			return o
		} else {
			if index >= len(a) {
				if panics {
					panic(fmt.Sprintf("objx: Index %d is out of range because the array only contains %d items.", index, len(a)))
				}
				return Nil
			} else {
				return New(a[index])
			}
		}
	case []uint8:
		a := object.([]uint8)
		if isSet {
			a[index] = value.(uint8)
			return o
		} else {
			if index >= len(a) {
				if panics {
					panic(fmt.Sprintf("objx: Index %d is out of range because the array only contains %d items.", index, len(a)))
				}
				return Nil
			} else {
				return New(a[index])
			}
		}
	case []uint16:
		a := object.([]uint16)
		if isSet {
			a[index] = value.(uint16)
			return o
		} else {
			if index >= len(a) {
				if panics {
					panic(fmt.Sprintf("objx: Index %d is out of range because the array only contains %d items.", index, len(a)))
				}
				return Nil
			} else {
				return New(a[index])
			}
		}
	case []uint32:
		a := object.([]uint32)
		if isSet {
			a[index] = value.(uint32)
			return o
		} else {
			if index >= len(a) {
				if panics {
					panic(fmt.Sprintf("objx: Index %d is out of range because the array only contains %d items.", index, len(a)))
				}
				return Nil
			} else {
				return New(a[index])
			}
		}
	case []uint64:
		a := object.([]uint64)
		if isSet {
			a[index] = value.(uint64)
			return o
		} else {
			if index >= len(a) {
				if panics {
					panic(fmt.Sprintf("objx: Index %d is out of range because the array only contains %d items.", index, len(a)))
				}
				return Nil
			} else {
				return New(a[index])
			}
		}
	case []float32:
		a := object.([]float32)
		if isSet {
			a[index] = value.(float32)
			return o
		} else {
			if index >= len(a) {
				if panics {
					panic(fmt.Sprintf("objx: Index %d is out of range because the array only contains %d items.", index, len(a)))
				}
				return Nil
			} else {
				return New(a[index])
			}
		}
	case []float64:
		a := object.([]float64)
		if isSet {
			a[index] = value.(float64)
			return o
		} else {
			if index >= len(a) {
				if panics {
					panic(fmt.Sprintf("objx: Index %d is out of range because the array only contains %d items.", index, len(a)))
				}
				return Nil
			} else {
				return New(a[index])
			}
		}
	case []complex64:
		a := object.([]complex64)
		if isSet {
			a[index] = value.(complex64)
			return o
		} else {
			if index >= len(a) {
				if panics {
					panic(fmt.Sprintf("objx: Index %d is out of range because the array only contains %d items.", index, len(a)))
				}
				return Nil
			} else {
				return New(a[index])
			}
		}
	case []complex128:
		a := object.([]complex128)
		if isSet {
			a[index] = value.(complex128)
			return o
		} else {
			if index >= len(a) {
				if panics {
					panic(fmt.Sprintf("objx: Index %d is out of range because the array only contains %d items.", index, len(a)))
				}
				return Nil
			} else {
				return New(a[index])
			}
		}
	case []string:
		a := object.([]string)
		if isSet {
			a[index] = value.(string)
			return o
		} else {
			if index >= len(a) {
				if panics {
					panic(fmt.Sprintf("objx: Index %d is out of range because the array only contains %d items.", index, len(a)))
				}
				return Nil
			} else {
				return New(a[index])
			}
		}
	case []uintptr:
		a := object.([]uintptr)
		if isSet {
			a[index] = value.(uintptr)
			return o
		} else {
			if index >= len(a) {
				if panics {
					panic(fmt.Sprintf("objx: Index %d is out of range because the array only contains %d items.", index, len(a)))
				}
				return Nil
			} else {
				return New(a[index])
			}
		}
	case []interface{}:
		a := object.([]interface{})
		if isSet {
			a[index] = value.(interface{})
			return o
		} else {
			if index >= len(a) {
				if panics {
					panic(fmt.Sprintf("objx: Index %d is out of range because the array only contains %d items.", index, len(a)))
				}
				return Nil
			} else {
				return New(a[index])
			}
		}
	default:
		// let reflect deal with this
		panic("objx: Cannot index into a non-array type.  " + reflect.TypeOf(object).Name() + " is not supported.")
	}

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
