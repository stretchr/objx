package objx

import (
	"reflect"
)

/*
	Inter (interface{} and []interface{})
	--------------------------------------------------
*/

// Inter gets the value as a interface{}, returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *Obj) Inter(optionalDefault ...interface{}) interface{} {
	if s, ok := o.obj.(interface{}); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustInter gets the value as a interface{}.
//
// Panics if the object is not a interface{}.
func (o *Obj) MustInter() interface{} {
	return o.obj.(interface{})
}

// InterSlice gets the value as a []interface{}, returns the optionalDefault
// value or nil if the value is not a []interface{}.
func (o *Obj) InterSlice(optionalDefault ...[]interface{}) []interface{} {
	if s, ok := o.obj.([]interface{}); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustInterSlice gets the value as a []interface{}.
//
// Panics if the object is not a []interface{}.
func (o *Obj) MustInterSlice() []interface{} {
	return o.obj.([]interface{})
}

// IsInter gets whether the object contained is a interface{} or not.
func (o *Obj) IsInter() bool {
	return o.IsKind(reflect.Interface)
}

// IsInterSlice gets whether the object contained is a []interface{} or not.
func (o *Obj) IsInterSlice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]interface{})
	return ok
}

// EachInter calls the specified callback for each object
// in the []interface{}.
//
// Panics if the object is the wrong type.
func (o *Obj) EachInter(callback func(int, interface{}) bool) *Obj {

	for index, val := range o.MustInterSlice() {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

// WhereInter uses the specified decider function to select items
// from the []interface{}.  The object contained in the result will contain
// only the selected items.
func (o *Obj) WhereInter(decider func(int, interface{}) bool) *Obj {

	var selected []interface{}

	o.EachInter(func(index int, val interface{}) bool {
		shouldSelect := decider(index, val)
		if shouldSelect == false {
			selected = append(selected, val)
		}
		return true
	})

	return New(selected)

}

// GroupInter uses the specified grouper function to group the items
// keyed by the return of the grouper.  The object contained in the
// result will contain a map[string][]interface{}.
func (o *Obj) GroupInter(grouper func(int, interface{}) string) *Obj {

	groups := make(map[string][]interface{})

	o.EachInter(func(index int, val interface{}) bool {
		group := grouper(index, val)
		if _, ok := groups[group]; !ok {
			groups[group] = make([]interface{}, 0)
		}
		groups[group] = append(groups[group], val)
		return true
	})

	return New(groups)

}

// ReplaceInter uses the specified function to replace each interface{}s
// by iterating each item.  The data in the returned result will be a
// []interface{} containing the replaced items.
func (o *Obj) ReplaceInter(replacer func(int, interface{}) interface{}) *Obj {

	arr := o.MustInterSlice()
	replaced := make([]interface{}, len(arr))

	o.EachInter(func(index int, val interface{}) bool {
		replaced[index] = replacer(index, val)
		return true
	})

	return New(replaced)

}

// CollectInter uses the specified collector function to collect a value
// for each of the interface{}s in the slice.  The data returned will be a
// []interface{}.
func (o *Obj) CollectInter(collector func(int, interface{}) interface{}) *Obj {

	arr := o.MustInterSlice()
	collected := make([]interface{}, len(arr))

	o.EachInter(func(index int, val interface{}) bool {
		collected[index] = collector(index, val)
		return true
	})

	return New(collected)
}

/*
	Bool (bool and []bool)
	--------------------------------------------------
*/

// Bool gets the value as a bool, returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *Obj) Bool(optionalDefault ...bool) bool {
	if s, ok := o.obj.(bool); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return false
}

// MustBool gets the value as a bool.
//
// Panics if the object is not a bool.
func (o *Obj) MustBool() bool {
	return o.obj.(bool)
}

// BoolSlice gets the value as a []bool, returns the optionalDefault
// value or nil if the value is not a []bool.
func (o *Obj) BoolSlice(optionalDefault ...[]bool) []bool {
	if s, ok := o.obj.([]bool); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustBoolSlice gets the value as a []bool.
//
// Panics if the object is not a []bool.
func (o *Obj) MustBoolSlice() []bool {
	return o.obj.([]bool)
}

// IsBool gets whether the object contained is a bool or not.
func (o *Obj) IsBool() bool {
	return o.IsKind(reflect.Bool)
}

// IsBoolSlice gets whether the object contained is a []bool or not.
func (o *Obj) IsBoolSlice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]bool)
	return ok
}

// EachBool calls the specified callback for each object
// in the []bool.
//
// Panics if the object is the wrong type.
func (o *Obj) EachBool(callback func(int, bool) bool) *Obj {

	for index, val := range o.MustBoolSlice() {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

// WhereBool uses the specified decider function to select items
// from the []bool.  The object contained in the result will contain
// only the selected items.
func (o *Obj) WhereBool(decider func(int, bool) bool) *Obj {

	var selected []bool

	o.EachBool(func(index int, val bool) bool {
		shouldSelect := decider(index, val)
		if shouldSelect == false {
			selected = append(selected, val)
		}
		return true
	})

	return New(selected)

}

// GroupBool uses the specified grouper function to group the items
// keyed by the return of the grouper.  The object contained in the
// result will contain a map[string][]bool.
func (o *Obj) GroupBool(grouper func(int, bool) string) *Obj {

	groups := make(map[string][]bool)

	o.EachBool(func(index int, val bool) bool {
		group := grouper(index, val)
		if _, ok := groups[group]; !ok {
			groups[group] = make([]bool, 0)
		}
		groups[group] = append(groups[group], val)
		return true
	})

	return New(groups)

}

// ReplaceBool uses the specified function to replace each bools
// by iterating each item.  The data in the returned result will be a
// []bool containing the replaced items.
func (o *Obj) ReplaceBool(replacer func(int, bool) bool) *Obj {

	arr := o.MustBoolSlice()
	replaced := make([]bool, len(arr))

	o.EachBool(func(index int, val bool) bool {
		replaced[index] = replacer(index, val)
		return true
	})

	return New(replaced)

}

// CollectBool uses the specified collector function to collect a value
// for each of the bools in the slice.  The data returned will be a
// []interface{}.
func (o *Obj) CollectBool(collector func(int, bool) interface{}) *Obj {

	arr := o.MustBoolSlice()
	collected := make([]interface{}, len(arr))

	o.EachBool(func(index int, val bool) bool {
		collected[index] = collector(index, val)
		return true
	})

	return New(collected)
}

/*
	Str (string and []string)
	--------------------------------------------------
*/

// Str gets the value as a string, returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *Obj) Str(optionalDefault ...string) string {
	if s, ok := o.obj.(string); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return ""
}

// MustStr gets the value as a string.
//
// Panics if the object is not a string.
func (o *Obj) MustStr() string {
	return o.obj.(string)
}

// StrSlice gets the value as a []string, returns the optionalDefault
// value or nil if the value is not a []string.
func (o *Obj) StrSlice(optionalDefault ...[]string) []string {
	if s, ok := o.obj.([]string); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustStrSlice gets the value as a []string.
//
// Panics if the object is not a []string.
func (o *Obj) MustStrSlice() []string {
	return o.obj.([]string)
}

// IsStr gets whether the object contained is a string or not.
func (o *Obj) IsStr() bool {
	return o.IsKind(reflect.String)
}

// IsStrSlice gets whether the object contained is a []string or not.
func (o *Obj) IsStrSlice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]string)
	return ok
}

// EachStr calls the specified callback for each object
// in the []string.
//
// Panics if the object is the wrong type.
func (o *Obj) EachStr(callback func(int, string) bool) *Obj {

	for index, val := range o.MustStrSlice() {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

// WhereStr uses the specified decider function to select items
// from the []string.  The object contained in the result will contain
// only the selected items.
func (o *Obj) WhereStr(decider func(int, string) bool) *Obj {

	var selected []string

	o.EachStr(func(index int, val string) bool {
		shouldSelect := decider(index, val)
		if shouldSelect == false {
			selected = append(selected, val)
		}
		return true
	})

	return New(selected)

}

// GroupStr uses the specified grouper function to group the items
// keyed by the return of the grouper.  The object contained in the
// result will contain a map[string][]string.
func (o *Obj) GroupStr(grouper func(int, string) string) *Obj {

	groups := make(map[string][]string)

	o.EachStr(func(index int, val string) bool {
		group := grouper(index, val)
		if _, ok := groups[group]; !ok {
			groups[group] = make([]string, 0)
		}
		groups[group] = append(groups[group], val)
		return true
	})

	return New(groups)

}

// ReplaceStr uses the specified function to replace each strings
// by iterating each item.  The data in the returned result will be a
// []string containing the replaced items.
func (o *Obj) ReplaceStr(replacer func(int, string) string) *Obj {

	arr := o.MustStrSlice()
	replaced := make([]string, len(arr))

	o.EachStr(func(index int, val string) bool {
		replaced[index] = replacer(index, val)
		return true
	})

	return New(replaced)

}

// CollectStr uses the specified collector function to collect a value
// for each of the strings in the slice.  The data returned will be a
// []interface{}.
func (o *Obj) CollectStr(collector func(int, string) interface{}) *Obj {

	arr := o.MustStrSlice()
	collected := make([]interface{}, len(arr))

	o.EachStr(func(index int, val string) bool {
		collected[index] = collector(index, val)
		return true
	})

	return New(collected)
}

/*
	Int (int and []int)
	--------------------------------------------------
*/

// Int gets the value as a int, returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *Obj) Int(optionalDefault ...int) int {
	if s, ok := o.obj.(int); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// MustInt gets the value as a int.
//
// Panics if the object is not a int.
func (o *Obj) MustInt() int {
	return o.obj.(int)
}

// IntSlice gets the value as a []int, returns the optionalDefault
// value or nil if the value is not a []int.
func (o *Obj) IntSlice(optionalDefault ...[]int) []int {
	if s, ok := o.obj.([]int); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustIntSlice gets the value as a []int.
//
// Panics if the object is not a []int.
func (o *Obj) MustIntSlice() []int {
	return o.obj.([]int)
}

// IsInt gets whether the object contained is a int or not.
func (o *Obj) IsInt() bool {
	return o.IsKind(reflect.Int)
}

// IsIntSlice gets whether the object contained is a []int or not.
func (o *Obj) IsIntSlice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]int)
	return ok
}

// EachInt calls the specified callback for each object
// in the []int.
//
// Panics if the object is the wrong type.
func (o *Obj) EachInt(callback func(int, int) bool) *Obj {

	for index, val := range o.MustIntSlice() {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

// WhereInt uses the specified decider function to select items
// from the []int.  The object contained in the result will contain
// only the selected items.
func (o *Obj) WhereInt(decider func(int, int) bool) *Obj {

	var selected []int

	o.EachInt(func(index int, val int) bool {
		shouldSelect := decider(index, val)
		if shouldSelect == false {
			selected = append(selected, val)
		}
		return true
	})

	return New(selected)

}

// GroupInt uses the specified grouper function to group the items
// keyed by the return of the grouper.  The object contained in the
// result will contain a map[string][]int.
func (o *Obj) GroupInt(grouper func(int, int) string) *Obj {

	groups := make(map[string][]int)

	o.EachInt(func(index int, val int) bool {
		group := grouper(index, val)
		if _, ok := groups[group]; !ok {
			groups[group] = make([]int, 0)
		}
		groups[group] = append(groups[group], val)
		return true
	})

	return New(groups)

}

// ReplaceInt uses the specified function to replace each ints
// by iterating each item.  The data in the returned result will be a
// []int containing the replaced items.
func (o *Obj) ReplaceInt(replacer func(int, int) int) *Obj {

	arr := o.MustIntSlice()
	replaced := make([]int, len(arr))

	o.EachInt(func(index int, val int) bool {
		replaced[index] = replacer(index, val)
		return true
	})

	return New(replaced)

}

// CollectInt uses the specified collector function to collect a value
// for each of the ints in the slice.  The data returned will be a
// []interface{}.
func (o *Obj) CollectInt(collector func(int, int) interface{}) *Obj {

	arr := o.MustIntSlice()
	collected := make([]interface{}, len(arr))

	o.EachInt(func(index int, val int) bool {
		collected[index] = collector(index, val)
		return true
	})

	return New(collected)
}

/*
	Int8 (int8 and []int8)
	--------------------------------------------------
*/

// Int8 gets the value as a int8, returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *Obj) Int8(optionalDefault ...int8) int8 {
	if s, ok := o.obj.(int8); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// MustInt8 gets the value as a int8.
//
// Panics if the object is not a int8.
func (o *Obj) MustInt8() int8 {
	return o.obj.(int8)
}

// Int8Slice gets the value as a []int8, returns the optionalDefault
// value or nil if the value is not a []int8.
func (o *Obj) Int8Slice(optionalDefault ...[]int8) []int8 {
	if s, ok := o.obj.([]int8); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustInt8Slice gets the value as a []int8.
//
// Panics if the object is not a []int8.
func (o *Obj) MustInt8Slice() []int8 {
	return o.obj.([]int8)
}

// IsInt8 gets whether the object contained is a int8 or not.
func (o *Obj) IsInt8() bool {
	return o.IsKind(reflect.Int8)
}

// IsInt8Slice gets whether the object contained is a []int8 or not.
func (o *Obj) IsInt8Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]int8)
	return ok
}

// EachInt8 calls the specified callback for each object
// in the []int8.
//
// Panics if the object is the wrong type.
func (o *Obj) EachInt8(callback func(int, int8) bool) *Obj {

	for index, val := range o.MustInt8Slice() {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

// WhereInt8 uses the specified decider function to select items
// from the []int8.  The object contained in the result will contain
// only the selected items.
func (o *Obj) WhereInt8(decider func(int, int8) bool) *Obj {

	var selected []int8

	o.EachInt8(func(index int, val int8) bool {
		shouldSelect := decider(index, val)
		if shouldSelect == false {
			selected = append(selected, val)
		}
		return true
	})

	return New(selected)

}

// GroupInt8 uses the specified grouper function to group the items
// keyed by the return of the grouper.  The object contained in the
// result will contain a map[string][]int8.
func (o *Obj) GroupInt8(grouper func(int, int8) string) *Obj {

	groups := make(map[string][]int8)

	o.EachInt8(func(index int, val int8) bool {
		group := grouper(index, val)
		if _, ok := groups[group]; !ok {
			groups[group] = make([]int8, 0)
		}
		groups[group] = append(groups[group], val)
		return true
	})

	return New(groups)

}

// ReplaceInt8 uses the specified function to replace each int8s
// by iterating each item.  The data in the returned result will be a
// []int8 containing the replaced items.
func (o *Obj) ReplaceInt8(replacer func(int, int8) int8) *Obj {

	arr := o.MustInt8Slice()
	replaced := make([]int8, len(arr))

	o.EachInt8(func(index int, val int8) bool {
		replaced[index] = replacer(index, val)
		return true
	})

	return New(replaced)

}

// CollectInt8 uses the specified collector function to collect a value
// for each of the int8s in the slice.  The data returned will be a
// []interface{}.
func (o *Obj) CollectInt8(collector func(int, int8) interface{}) *Obj {

	arr := o.MustInt8Slice()
	collected := make([]interface{}, len(arr))

	o.EachInt8(func(index int, val int8) bool {
		collected[index] = collector(index, val)
		return true
	})

	return New(collected)
}

/*
	Int16 (int16 and []int16)
	--------------------------------------------------
*/

// Int16 gets the value as a int16, returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *Obj) Int16(optionalDefault ...int16) int16 {
	if s, ok := o.obj.(int16); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// MustInt16 gets the value as a int16.
//
// Panics if the object is not a int16.
func (o *Obj) MustInt16() int16 {
	return o.obj.(int16)
}

// Int16Slice gets the value as a []int16, returns the optionalDefault
// value or nil if the value is not a []int16.
func (o *Obj) Int16Slice(optionalDefault ...[]int16) []int16 {
	if s, ok := o.obj.([]int16); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustInt16Slice gets the value as a []int16.
//
// Panics if the object is not a []int16.
func (o *Obj) MustInt16Slice() []int16 {
	return o.obj.([]int16)
}

// IsInt16 gets whether the object contained is a int16 or not.
func (o *Obj) IsInt16() bool {
	return o.IsKind(reflect.Int16)
}

// IsInt16Slice gets whether the object contained is a []int16 or not.
func (o *Obj) IsInt16Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]int16)
	return ok
}

// EachInt16 calls the specified callback for each object
// in the []int16.
//
// Panics if the object is the wrong type.
func (o *Obj) EachInt16(callback func(int, int16) bool) *Obj {

	for index, val := range o.MustInt16Slice() {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

// WhereInt16 uses the specified decider function to select items
// from the []int16.  The object contained in the result will contain
// only the selected items.
func (o *Obj) WhereInt16(decider func(int, int16) bool) *Obj {

	var selected []int16

	o.EachInt16(func(index int, val int16) bool {
		shouldSelect := decider(index, val)
		if shouldSelect == false {
			selected = append(selected, val)
		}
		return true
	})

	return New(selected)

}

// GroupInt16 uses the specified grouper function to group the items
// keyed by the return of the grouper.  The object contained in the
// result will contain a map[string][]int16.
func (o *Obj) GroupInt16(grouper func(int, int16) string) *Obj {

	groups := make(map[string][]int16)

	o.EachInt16(func(index int, val int16) bool {
		group := grouper(index, val)
		if _, ok := groups[group]; !ok {
			groups[group] = make([]int16, 0)
		}
		groups[group] = append(groups[group], val)
		return true
	})

	return New(groups)

}

// ReplaceInt16 uses the specified function to replace each int16s
// by iterating each item.  The data in the returned result will be a
// []int16 containing the replaced items.
func (o *Obj) ReplaceInt16(replacer func(int, int16) int16) *Obj {

	arr := o.MustInt16Slice()
	replaced := make([]int16, len(arr))

	o.EachInt16(func(index int, val int16) bool {
		replaced[index] = replacer(index, val)
		return true
	})

	return New(replaced)

}

// CollectInt16 uses the specified collector function to collect a value
// for each of the int16s in the slice.  The data returned will be a
// []interface{}.
func (o *Obj) CollectInt16(collector func(int, int16) interface{}) *Obj {

	arr := o.MustInt16Slice()
	collected := make([]interface{}, len(arr))

	o.EachInt16(func(index int, val int16) bool {
		collected[index] = collector(index, val)
		return true
	})

	return New(collected)
}

/*
	Int32 (int32 and []int32)
	--------------------------------------------------
*/

// Int32 gets the value as a int32, returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *Obj) Int32(optionalDefault ...int32) int32 {
	if s, ok := o.obj.(int32); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// MustInt32 gets the value as a int32.
//
// Panics if the object is not a int32.
func (o *Obj) MustInt32() int32 {
	return o.obj.(int32)
}

// Int32Slice gets the value as a []int32, returns the optionalDefault
// value or nil if the value is not a []int32.
func (o *Obj) Int32Slice(optionalDefault ...[]int32) []int32 {
	if s, ok := o.obj.([]int32); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustInt32Slice gets the value as a []int32.
//
// Panics if the object is not a []int32.
func (o *Obj) MustInt32Slice() []int32 {
	return o.obj.([]int32)
}

// IsInt32 gets whether the object contained is a int32 or not.
func (o *Obj) IsInt32() bool {
	return o.IsKind(reflect.Int32)
}

// IsInt32Slice gets whether the object contained is a []int32 or not.
func (o *Obj) IsInt32Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]int32)
	return ok
}

// EachInt32 calls the specified callback for each object
// in the []int32.
//
// Panics if the object is the wrong type.
func (o *Obj) EachInt32(callback func(int, int32) bool) *Obj {

	for index, val := range o.MustInt32Slice() {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

// WhereInt32 uses the specified decider function to select items
// from the []int32.  The object contained in the result will contain
// only the selected items.
func (o *Obj) WhereInt32(decider func(int, int32) bool) *Obj {

	var selected []int32

	o.EachInt32(func(index int, val int32) bool {
		shouldSelect := decider(index, val)
		if shouldSelect == false {
			selected = append(selected, val)
		}
		return true
	})

	return New(selected)

}

// GroupInt32 uses the specified grouper function to group the items
// keyed by the return of the grouper.  The object contained in the
// result will contain a map[string][]int32.
func (o *Obj) GroupInt32(grouper func(int, int32) string) *Obj {

	groups := make(map[string][]int32)

	o.EachInt32(func(index int, val int32) bool {
		group := grouper(index, val)
		if _, ok := groups[group]; !ok {
			groups[group] = make([]int32, 0)
		}
		groups[group] = append(groups[group], val)
		return true
	})

	return New(groups)

}

// ReplaceInt32 uses the specified function to replace each int32s
// by iterating each item.  The data in the returned result will be a
// []int32 containing the replaced items.
func (o *Obj) ReplaceInt32(replacer func(int, int32) int32) *Obj {

	arr := o.MustInt32Slice()
	replaced := make([]int32, len(arr))

	o.EachInt32(func(index int, val int32) bool {
		replaced[index] = replacer(index, val)
		return true
	})

	return New(replaced)

}

// CollectInt32 uses the specified collector function to collect a value
// for each of the int32s in the slice.  The data returned will be a
// []interface{}.
func (o *Obj) CollectInt32(collector func(int, int32) interface{}) *Obj {

	arr := o.MustInt32Slice()
	collected := make([]interface{}, len(arr))

	o.EachInt32(func(index int, val int32) bool {
		collected[index] = collector(index, val)
		return true
	})

	return New(collected)
}

/*
	Int64 (int64 and []int64)
	--------------------------------------------------
*/

// Int64 gets the value as a int64, returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *Obj) Int64(optionalDefault ...int64) int64 {
	if s, ok := o.obj.(int64); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// MustInt64 gets the value as a int64.
//
// Panics if the object is not a int64.
func (o *Obj) MustInt64() int64 {
	return o.obj.(int64)
}

// Int64Slice gets the value as a []int64, returns the optionalDefault
// value or nil if the value is not a []int64.
func (o *Obj) Int64Slice(optionalDefault ...[]int64) []int64 {
	if s, ok := o.obj.([]int64); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustInt64Slice gets the value as a []int64.
//
// Panics if the object is not a []int64.
func (o *Obj) MustInt64Slice() []int64 {
	return o.obj.([]int64)
}

// IsInt64 gets whether the object contained is a int64 or not.
func (o *Obj) IsInt64() bool {
	return o.IsKind(reflect.Int64)
}

// IsInt64Slice gets whether the object contained is a []int64 or not.
func (o *Obj) IsInt64Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]int64)
	return ok
}

// EachInt64 calls the specified callback for each object
// in the []int64.
//
// Panics if the object is the wrong type.
func (o *Obj) EachInt64(callback func(int, int64) bool) *Obj {

	for index, val := range o.MustInt64Slice() {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

// WhereInt64 uses the specified decider function to select items
// from the []int64.  The object contained in the result will contain
// only the selected items.
func (o *Obj) WhereInt64(decider func(int, int64) bool) *Obj {

	var selected []int64

	o.EachInt64(func(index int, val int64) bool {
		shouldSelect := decider(index, val)
		if shouldSelect == false {
			selected = append(selected, val)
		}
		return true
	})

	return New(selected)

}

// GroupInt64 uses the specified grouper function to group the items
// keyed by the return of the grouper.  The object contained in the
// result will contain a map[string][]int64.
func (o *Obj) GroupInt64(grouper func(int, int64) string) *Obj {

	groups := make(map[string][]int64)

	o.EachInt64(func(index int, val int64) bool {
		group := grouper(index, val)
		if _, ok := groups[group]; !ok {
			groups[group] = make([]int64, 0)
		}
		groups[group] = append(groups[group], val)
		return true
	})

	return New(groups)

}

// ReplaceInt64 uses the specified function to replace each int64s
// by iterating each item.  The data in the returned result will be a
// []int64 containing the replaced items.
func (o *Obj) ReplaceInt64(replacer func(int, int64) int64) *Obj {

	arr := o.MustInt64Slice()
	replaced := make([]int64, len(arr))

	o.EachInt64(func(index int, val int64) bool {
		replaced[index] = replacer(index, val)
		return true
	})

	return New(replaced)

}

// CollectInt64 uses the specified collector function to collect a value
// for each of the int64s in the slice.  The data returned will be a
// []interface{}.
func (o *Obj) CollectInt64(collector func(int, int64) interface{}) *Obj {

	arr := o.MustInt64Slice()
	collected := make([]interface{}, len(arr))

	o.EachInt64(func(index int, val int64) bool {
		collected[index] = collector(index, val)
		return true
	})

	return New(collected)
}

/*
	Uint (uint and []uint)
	--------------------------------------------------
*/

// Uint gets the value as a uint, returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *Obj) Uint(optionalDefault ...uint) uint {
	if s, ok := o.obj.(uint); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// MustUint gets the value as a uint.
//
// Panics if the object is not a uint.
func (o *Obj) MustUint() uint {
	return o.obj.(uint)
}

// UintSlice gets the value as a []uint, returns the optionalDefault
// value or nil if the value is not a []uint.
func (o *Obj) UintSlice(optionalDefault ...[]uint) []uint {
	if s, ok := o.obj.([]uint); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustUintSlice gets the value as a []uint.
//
// Panics if the object is not a []uint.
func (o *Obj) MustUintSlice() []uint {
	return o.obj.([]uint)
}

// IsUint gets whether the object contained is a uint or not.
func (o *Obj) IsUint() bool {
	return o.IsKind(reflect.Uint)
}

// IsUintSlice gets whether the object contained is a []uint or not.
func (o *Obj) IsUintSlice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]uint)
	return ok
}

// EachUint calls the specified callback for each object
// in the []uint.
//
// Panics if the object is the wrong type.
func (o *Obj) EachUint(callback func(int, uint) bool) *Obj {

	for index, val := range o.MustUintSlice() {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

// WhereUint uses the specified decider function to select items
// from the []uint.  The object contained in the result will contain
// only the selected items.
func (o *Obj) WhereUint(decider func(int, uint) bool) *Obj {

	var selected []uint

	o.EachUint(func(index int, val uint) bool {
		shouldSelect := decider(index, val)
		if shouldSelect == false {
			selected = append(selected, val)
		}
		return true
	})

	return New(selected)

}

// GroupUint uses the specified grouper function to group the items
// keyed by the return of the grouper.  The object contained in the
// result will contain a map[string][]uint.
func (o *Obj) GroupUint(grouper func(int, uint) string) *Obj {

	groups := make(map[string][]uint)

	o.EachUint(func(index int, val uint) bool {
		group := grouper(index, val)
		if _, ok := groups[group]; !ok {
			groups[group] = make([]uint, 0)
		}
		groups[group] = append(groups[group], val)
		return true
	})

	return New(groups)

}

// ReplaceUint uses the specified function to replace each uints
// by iterating each item.  The data in the returned result will be a
// []uint containing the replaced items.
func (o *Obj) ReplaceUint(replacer func(int, uint) uint) *Obj {

	arr := o.MustUintSlice()
	replaced := make([]uint, len(arr))

	o.EachUint(func(index int, val uint) bool {
		replaced[index] = replacer(index, val)
		return true
	})

	return New(replaced)

}

// CollectUint uses the specified collector function to collect a value
// for each of the uints in the slice.  The data returned will be a
// []interface{}.
func (o *Obj) CollectUint(collector func(int, uint) interface{}) *Obj {

	arr := o.MustUintSlice()
	collected := make([]interface{}, len(arr))

	o.EachUint(func(index int, val uint) bool {
		collected[index] = collector(index, val)
		return true
	})

	return New(collected)
}

/*
	Uint8 (uint8 and []uint8)
	--------------------------------------------------
*/

// Uint8 gets the value as a uint8, returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *Obj) Uint8(optionalDefault ...uint8) uint8 {
	if s, ok := o.obj.(uint8); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// MustUint8 gets the value as a uint8.
//
// Panics if the object is not a uint8.
func (o *Obj) MustUint8() uint8 {
	return o.obj.(uint8)
}

// Uint8Slice gets the value as a []uint8, returns the optionalDefault
// value or nil if the value is not a []uint8.
func (o *Obj) Uint8Slice(optionalDefault ...[]uint8) []uint8 {
	if s, ok := o.obj.([]uint8); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustUint8Slice gets the value as a []uint8.
//
// Panics if the object is not a []uint8.
func (o *Obj) MustUint8Slice() []uint8 {
	return o.obj.([]uint8)
}

// IsUint8 gets whether the object contained is a uint8 or not.
func (o *Obj) IsUint8() bool {
	return o.IsKind(reflect.Uint8)
}

// IsUint8Slice gets whether the object contained is a []uint8 or not.
func (o *Obj) IsUint8Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]uint8)
	return ok
}

// EachUint8 calls the specified callback for each object
// in the []uint8.
//
// Panics if the object is the wrong type.
func (o *Obj) EachUint8(callback func(int, uint8) bool) *Obj {

	for index, val := range o.MustUint8Slice() {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

// WhereUint8 uses the specified decider function to select items
// from the []uint8.  The object contained in the result will contain
// only the selected items.
func (o *Obj) WhereUint8(decider func(int, uint8) bool) *Obj {

	var selected []uint8

	o.EachUint8(func(index int, val uint8) bool {
		shouldSelect := decider(index, val)
		if shouldSelect == false {
			selected = append(selected, val)
		}
		return true
	})

	return New(selected)

}

// GroupUint8 uses the specified grouper function to group the items
// keyed by the return of the grouper.  The object contained in the
// result will contain a map[string][]uint8.
func (o *Obj) GroupUint8(grouper func(int, uint8) string) *Obj {

	groups := make(map[string][]uint8)

	o.EachUint8(func(index int, val uint8) bool {
		group := grouper(index, val)
		if _, ok := groups[group]; !ok {
			groups[group] = make([]uint8, 0)
		}
		groups[group] = append(groups[group], val)
		return true
	})

	return New(groups)

}

// ReplaceUint8 uses the specified function to replace each uint8s
// by iterating each item.  The data in the returned result will be a
// []uint8 containing the replaced items.
func (o *Obj) ReplaceUint8(replacer func(int, uint8) uint8) *Obj {

	arr := o.MustUint8Slice()
	replaced := make([]uint8, len(arr))

	o.EachUint8(func(index int, val uint8) bool {
		replaced[index] = replacer(index, val)
		return true
	})

	return New(replaced)

}

// CollectUint8 uses the specified collector function to collect a value
// for each of the uint8s in the slice.  The data returned will be a
// []interface{}.
func (o *Obj) CollectUint8(collector func(int, uint8) interface{}) *Obj {

	arr := o.MustUint8Slice()
	collected := make([]interface{}, len(arr))

	o.EachUint8(func(index int, val uint8) bool {
		collected[index] = collector(index, val)
		return true
	})

	return New(collected)
}

/*
	Uint16 (uint16 and []uint16)
	--------------------------------------------------
*/

// Uint16 gets the value as a uint16, returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *Obj) Uint16(optionalDefault ...uint16) uint16 {
	if s, ok := o.obj.(uint16); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// MustUint16 gets the value as a uint16.
//
// Panics if the object is not a uint16.
func (o *Obj) MustUint16() uint16 {
	return o.obj.(uint16)
}

// Uint16Slice gets the value as a []uint16, returns the optionalDefault
// value or nil if the value is not a []uint16.
func (o *Obj) Uint16Slice(optionalDefault ...[]uint16) []uint16 {
	if s, ok := o.obj.([]uint16); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustUint16Slice gets the value as a []uint16.
//
// Panics if the object is not a []uint16.
func (o *Obj) MustUint16Slice() []uint16 {
	return o.obj.([]uint16)
}

// IsUint16 gets whether the object contained is a uint16 or not.
func (o *Obj) IsUint16() bool {
	return o.IsKind(reflect.Uint16)
}

// IsUint16Slice gets whether the object contained is a []uint16 or not.
func (o *Obj) IsUint16Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]uint16)
	return ok
}

// EachUint16 calls the specified callback for each object
// in the []uint16.
//
// Panics if the object is the wrong type.
func (o *Obj) EachUint16(callback func(int, uint16) bool) *Obj {

	for index, val := range o.MustUint16Slice() {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

// WhereUint16 uses the specified decider function to select items
// from the []uint16.  The object contained in the result will contain
// only the selected items.
func (o *Obj) WhereUint16(decider func(int, uint16) bool) *Obj {

	var selected []uint16

	o.EachUint16(func(index int, val uint16) bool {
		shouldSelect := decider(index, val)
		if shouldSelect == false {
			selected = append(selected, val)
		}
		return true
	})

	return New(selected)

}

// GroupUint16 uses the specified grouper function to group the items
// keyed by the return of the grouper.  The object contained in the
// result will contain a map[string][]uint16.
func (o *Obj) GroupUint16(grouper func(int, uint16) string) *Obj {

	groups := make(map[string][]uint16)

	o.EachUint16(func(index int, val uint16) bool {
		group := grouper(index, val)
		if _, ok := groups[group]; !ok {
			groups[group] = make([]uint16, 0)
		}
		groups[group] = append(groups[group], val)
		return true
	})

	return New(groups)

}

// ReplaceUint16 uses the specified function to replace each uint16s
// by iterating each item.  The data in the returned result will be a
// []uint16 containing the replaced items.
func (o *Obj) ReplaceUint16(replacer func(int, uint16) uint16) *Obj {

	arr := o.MustUint16Slice()
	replaced := make([]uint16, len(arr))

	o.EachUint16(func(index int, val uint16) bool {
		replaced[index] = replacer(index, val)
		return true
	})

	return New(replaced)

}

// CollectUint16 uses the specified collector function to collect a value
// for each of the uint16s in the slice.  The data returned will be a
// []interface{}.
func (o *Obj) CollectUint16(collector func(int, uint16) interface{}) *Obj {

	arr := o.MustUint16Slice()
	collected := make([]interface{}, len(arr))

	o.EachUint16(func(index int, val uint16) bool {
		collected[index] = collector(index, val)
		return true
	})

	return New(collected)
}

/*
	Uint32 (uint32 and []uint32)
	--------------------------------------------------
*/

// Uint32 gets the value as a uint32, returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *Obj) Uint32(optionalDefault ...uint32) uint32 {
	if s, ok := o.obj.(uint32); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// MustUint32 gets the value as a uint32.
//
// Panics if the object is not a uint32.
func (o *Obj) MustUint32() uint32 {
	return o.obj.(uint32)
}

// Uint32Slice gets the value as a []uint32, returns the optionalDefault
// value or nil if the value is not a []uint32.
func (o *Obj) Uint32Slice(optionalDefault ...[]uint32) []uint32 {
	if s, ok := o.obj.([]uint32); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustUint32Slice gets the value as a []uint32.
//
// Panics if the object is not a []uint32.
func (o *Obj) MustUint32Slice() []uint32 {
	return o.obj.([]uint32)
}

// IsUint32 gets whether the object contained is a uint32 or not.
func (o *Obj) IsUint32() bool {
	return o.IsKind(reflect.Uint32)
}

// IsUint32Slice gets whether the object contained is a []uint32 or not.
func (o *Obj) IsUint32Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]uint32)
	return ok
}

// EachUint32 calls the specified callback for each object
// in the []uint32.
//
// Panics if the object is the wrong type.
func (o *Obj) EachUint32(callback func(int, uint32) bool) *Obj {

	for index, val := range o.MustUint32Slice() {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

// WhereUint32 uses the specified decider function to select items
// from the []uint32.  The object contained in the result will contain
// only the selected items.
func (o *Obj) WhereUint32(decider func(int, uint32) bool) *Obj {

	var selected []uint32

	o.EachUint32(func(index int, val uint32) bool {
		shouldSelect := decider(index, val)
		if shouldSelect == false {
			selected = append(selected, val)
		}
		return true
	})

	return New(selected)

}

// GroupUint32 uses the specified grouper function to group the items
// keyed by the return of the grouper.  The object contained in the
// result will contain a map[string][]uint32.
func (o *Obj) GroupUint32(grouper func(int, uint32) string) *Obj {

	groups := make(map[string][]uint32)

	o.EachUint32(func(index int, val uint32) bool {
		group := grouper(index, val)
		if _, ok := groups[group]; !ok {
			groups[group] = make([]uint32, 0)
		}
		groups[group] = append(groups[group], val)
		return true
	})

	return New(groups)

}

// ReplaceUint32 uses the specified function to replace each uint32s
// by iterating each item.  The data in the returned result will be a
// []uint32 containing the replaced items.
func (o *Obj) ReplaceUint32(replacer func(int, uint32) uint32) *Obj {

	arr := o.MustUint32Slice()
	replaced := make([]uint32, len(arr))

	o.EachUint32(func(index int, val uint32) bool {
		replaced[index] = replacer(index, val)
		return true
	})

	return New(replaced)

}

// CollectUint32 uses the specified collector function to collect a value
// for each of the uint32s in the slice.  The data returned will be a
// []interface{}.
func (o *Obj) CollectUint32(collector func(int, uint32) interface{}) *Obj {

	arr := o.MustUint32Slice()
	collected := make([]interface{}, len(arr))

	o.EachUint32(func(index int, val uint32) bool {
		collected[index] = collector(index, val)
		return true
	})

	return New(collected)
}

/*
	Uint64 (uint64 and []uint64)
	--------------------------------------------------
*/

// Uint64 gets the value as a uint64, returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *Obj) Uint64(optionalDefault ...uint64) uint64 {
	if s, ok := o.obj.(uint64); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// MustUint64 gets the value as a uint64.
//
// Panics if the object is not a uint64.
func (o *Obj) MustUint64() uint64 {
	return o.obj.(uint64)
}

// Uint64Slice gets the value as a []uint64, returns the optionalDefault
// value or nil if the value is not a []uint64.
func (o *Obj) Uint64Slice(optionalDefault ...[]uint64) []uint64 {
	if s, ok := o.obj.([]uint64); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustUint64Slice gets the value as a []uint64.
//
// Panics if the object is not a []uint64.
func (o *Obj) MustUint64Slice() []uint64 {
	return o.obj.([]uint64)
}

// IsUint64 gets whether the object contained is a uint64 or not.
func (o *Obj) IsUint64() bool {
	return o.IsKind(reflect.Uint64)
}

// IsUint64Slice gets whether the object contained is a []uint64 or not.
func (o *Obj) IsUint64Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]uint64)
	return ok
}

// EachUint64 calls the specified callback for each object
// in the []uint64.
//
// Panics if the object is the wrong type.
func (o *Obj) EachUint64(callback func(int, uint64) bool) *Obj {

	for index, val := range o.MustUint64Slice() {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

// WhereUint64 uses the specified decider function to select items
// from the []uint64.  The object contained in the result will contain
// only the selected items.
func (o *Obj) WhereUint64(decider func(int, uint64) bool) *Obj {

	var selected []uint64

	o.EachUint64(func(index int, val uint64) bool {
		shouldSelect := decider(index, val)
		if shouldSelect == false {
			selected = append(selected, val)
		}
		return true
	})

	return New(selected)

}

// GroupUint64 uses the specified grouper function to group the items
// keyed by the return of the grouper.  The object contained in the
// result will contain a map[string][]uint64.
func (o *Obj) GroupUint64(grouper func(int, uint64) string) *Obj {

	groups := make(map[string][]uint64)

	o.EachUint64(func(index int, val uint64) bool {
		group := grouper(index, val)
		if _, ok := groups[group]; !ok {
			groups[group] = make([]uint64, 0)
		}
		groups[group] = append(groups[group], val)
		return true
	})

	return New(groups)

}

// ReplaceUint64 uses the specified function to replace each uint64s
// by iterating each item.  The data in the returned result will be a
// []uint64 containing the replaced items.
func (o *Obj) ReplaceUint64(replacer func(int, uint64) uint64) *Obj {

	arr := o.MustUint64Slice()
	replaced := make([]uint64, len(arr))

	o.EachUint64(func(index int, val uint64) bool {
		replaced[index] = replacer(index, val)
		return true
	})

	return New(replaced)

}

// CollectUint64 uses the specified collector function to collect a value
// for each of the uint64s in the slice.  The data returned will be a
// []interface{}.
func (o *Obj) CollectUint64(collector func(int, uint64) interface{}) *Obj {

	arr := o.MustUint64Slice()
	collected := make([]interface{}, len(arr))

	o.EachUint64(func(index int, val uint64) bool {
		collected[index] = collector(index, val)
		return true
	})

	return New(collected)
}

/*
	Uintptr (uintptr and []uintptr)
	--------------------------------------------------
*/

// Uintptr gets the value as a uintptr, returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *Obj) Uintptr(optionalDefault ...uintptr) uintptr {
	if s, ok := o.obj.(uintptr); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// MustUintptr gets the value as a uintptr.
//
// Panics if the object is not a uintptr.
func (o *Obj) MustUintptr() uintptr {
	return o.obj.(uintptr)
}

// UintptrSlice gets the value as a []uintptr, returns the optionalDefault
// value or nil if the value is not a []uintptr.
func (o *Obj) UintptrSlice(optionalDefault ...[]uintptr) []uintptr {
	if s, ok := o.obj.([]uintptr); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustUintptrSlice gets the value as a []uintptr.
//
// Panics if the object is not a []uintptr.
func (o *Obj) MustUintptrSlice() []uintptr {
	return o.obj.([]uintptr)
}

// IsUintptr gets whether the object contained is a uintptr or not.
func (o *Obj) IsUintptr() bool {
	return o.IsKind(reflect.Uintptr)
}

// IsUintptrSlice gets whether the object contained is a []uintptr or not.
func (o *Obj) IsUintptrSlice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]uintptr)
	return ok
}

// EachUintptr calls the specified callback for each object
// in the []uintptr.
//
// Panics if the object is the wrong type.
func (o *Obj) EachUintptr(callback func(int, uintptr) bool) *Obj {

	for index, val := range o.MustUintptrSlice() {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

// WhereUintptr uses the specified decider function to select items
// from the []uintptr.  The object contained in the result will contain
// only the selected items.
func (o *Obj) WhereUintptr(decider func(int, uintptr) bool) *Obj {

	var selected []uintptr

	o.EachUintptr(func(index int, val uintptr) bool {
		shouldSelect := decider(index, val)
		if shouldSelect == false {
			selected = append(selected, val)
		}
		return true
	})

	return New(selected)

}

// GroupUintptr uses the specified grouper function to group the items
// keyed by the return of the grouper.  The object contained in the
// result will contain a map[string][]uintptr.
func (o *Obj) GroupUintptr(grouper func(int, uintptr) string) *Obj {

	groups := make(map[string][]uintptr)

	o.EachUintptr(func(index int, val uintptr) bool {
		group := grouper(index, val)
		if _, ok := groups[group]; !ok {
			groups[group] = make([]uintptr, 0)
		}
		groups[group] = append(groups[group], val)
		return true
	})

	return New(groups)

}

// ReplaceUintptr uses the specified function to replace each uintptrs
// by iterating each item.  The data in the returned result will be a
// []uintptr containing the replaced items.
func (o *Obj) ReplaceUintptr(replacer func(int, uintptr) uintptr) *Obj {

	arr := o.MustUintptrSlice()
	replaced := make([]uintptr, len(arr))

	o.EachUintptr(func(index int, val uintptr) bool {
		replaced[index] = replacer(index, val)
		return true
	})

	return New(replaced)

}

// CollectUintptr uses the specified collector function to collect a value
// for each of the uintptrs in the slice.  The data returned will be a
// []interface{}.
func (o *Obj) CollectUintptr(collector func(int, uintptr) interface{}) *Obj {

	arr := o.MustUintptrSlice()
	collected := make([]interface{}, len(arr))

	o.EachUintptr(func(index int, val uintptr) bool {
		collected[index] = collector(index, val)
		return true
	})

	return New(collected)
}

/*
	Float32 (float32 and []float32)
	--------------------------------------------------
*/

// Float32 gets the value as a float32, returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *Obj) Float32(optionalDefault ...float32) float32 {
	if s, ok := o.obj.(float32); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// MustFloat32 gets the value as a float32.
//
// Panics if the object is not a float32.
func (o *Obj) MustFloat32() float32 {
	return o.obj.(float32)
}

// Float32Slice gets the value as a []float32, returns the optionalDefault
// value or nil if the value is not a []float32.
func (o *Obj) Float32Slice(optionalDefault ...[]float32) []float32 {
	if s, ok := o.obj.([]float32); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustFloat32Slice gets the value as a []float32.
//
// Panics if the object is not a []float32.
func (o *Obj) MustFloat32Slice() []float32 {
	return o.obj.([]float32)
}

// IsFloat32 gets whether the object contained is a float32 or not.
func (o *Obj) IsFloat32() bool {
	return o.IsKind(reflect.Float32)
}

// IsFloat32Slice gets whether the object contained is a []float32 or not.
func (o *Obj) IsFloat32Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]float32)
	return ok
}

// EachFloat32 calls the specified callback for each object
// in the []float32.
//
// Panics if the object is the wrong type.
func (o *Obj) EachFloat32(callback func(int, float32) bool) *Obj {

	for index, val := range o.MustFloat32Slice() {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

// WhereFloat32 uses the specified decider function to select items
// from the []float32.  The object contained in the result will contain
// only the selected items.
func (o *Obj) WhereFloat32(decider func(int, float32) bool) *Obj {

	var selected []float32

	o.EachFloat32(func(index int, val float32) bool {
		shouldSelect := decider(index, val)
		if shouldSelect == false {
			selected = append(selected, val)
		}
		return true
	})

	return New(selected)

}

// GroupFloat32 uses the specified grouper function to group the items
// keyed by the return of the grouper.  The object contained in the
// result will contain a map[string][]float32.
func (o *Obj) GroupFloat32(grouper func(int, float32) string) *Obj {

	groups := make(map[string][]float32)

	o.EachFloat32(func(index int, val float32) bool {
		group := grouper(index, val)
		if _, ok := groups[group]; !ok {
			groups[group] = make([]float32, 0)
		}
		groups[group] = append(groups[group], val)
		return true
	})

	return New(groups)

}

// ReplaceFloat32 uses the specified function to replace each float32s
// by iterating each item.  The data in the returned result will be a
// []float32 containing the replaced items.
func (o *Obj) ReplaceFloat32(replacer func(int, float32) float32) *Obj {

	arr := o.MustFloat32Slice()
	replaced := make([]float32, len(arr))

	o.EachFloat32(func(index int, val float32) bool {
		replaced[index] = replacer(index, val)
		return true
	})

	return New(replaced)

}

// CollectFloat32 uses the specified collector function to collect a value
// for each of the float32s in the slice.  The data returned will be a
// []interface{}.
func (o *Obj) CollectFloat32(collector func(int, float32) interface{}) *Obj {

	arr := o.MustFloat32Slice()
	collected := make([]interface{}, len(arr))

	o.EachFloat32(func(index int, val float32) bool {
		collected[index] = collector(index, val)
		return true
	})

	return New(collected)
}

/*
	Float64 (float64 and []float64)
	--------------------------------------------------
*/

// Float64 gets the value as a float64, returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *Obj) Float64(optionalDefault ...float64) float64 {
	if s, ok := o.obj.(float64); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// MustFloat64 gets the value as a float64.
//
// Panics if the object is not a float64.
func (o *Obj) MustFloat64() float64 {
	return o.obj.(float64)
}

// Float64Slice gets the value as a []float64, returns the optionalDefault
// value or nil if the value is not a []float64.
func (o *Obj) Float64Slice(optionalDefault ...[]float64) []float64 {
	if s, ok := o.obj.([]float64); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustFloat64Slice gets the value as a []float64.
//
// Panics if the object is not a []float64.
func (o *Obj) MustFloat64Slice() []float64 {
	return o.obj.([]float64)
}

// IsFloat64 gets whether the object contained is a float64 or not.
func (o *Obj) IsFloat64() bool {
	return o.IsKind(reflect.Float64)
}

// IsFloat64Slice gets whether the object contained is a []float64 or not.
func (o *Obj) IsFloat64Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]float64)
	return ok
}

// EachFloat64 calls the specified callback for each object
// in the []float64.
//
// Panics if the object is the wrong type.
func (o *Obj) EachFloat64(callback func(int, float64) bool) *Obj {

	for index, val := range o.MustFloat64Slice() {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

// WhereFloat64 uses the specified decider function to select items
// from the []float64.  The object contained in the result will contain
// only the selected items.
func (o *Obj) WhereFloat64(decider func(int, float64) bool) *Obj {

	var selected []float64

	o.EachFloat64(func(index int, val float64) bool {
		shouldSelect := decider(index, val)
		if shouldSelect == false {
			selected = append(selected, val)
		}
		return true
	})

	return New(selected)

}

// GroupFloat64 uses the specified grouper function to group the items
// keyed by the return of the grouper.  The object contained in the
// result will contain a map[string][]float64.
func (o *Obj) GroupFloat64(grouper func(int, float64) string) *Obj {

	groups := make(map[string][]float64)

	o.EachFloat64(func(index int, val float64) bool {
		group := grouper(index, val)
		if _, ok := groups[group]; !ok {
			groups[group] = make([]float64, 0)
		}
		groups[group] = append(groups[group], val)
		return true
	})

	return New(groups)

}

// ReplaceFloat64 uses the specified function to replace each float64s
// by iterating each item.  The data in the returned result will be a
// []float64 containing the replaced items.
func (o *Obj) ReplaceFloat64(replacer func(int, float64) float64) *Obj {

	arr := o.MustFloat64Slice()
	replaced := make([]float64, len(arr))

	o.EachFloat64(func(index int, val float64) bool {
		replaced[index] = replacer(index, val)
		return true
	})

	return New(replaced)

}

// CollectFloat64 uses the specified collector function to collect a value
// for each of the float64s in the slice.  The data returned will be a
// []interface{}.
func (o *Obj) CollectFloat64(collector func(int, float64) interface{}) *Obj {

	arr := o.MustFloat64Slice()
	collected := make([]interface{}, len(arr))

	o.EachFloat64(func(index int, val float64) bool {
		collected[index] = collector(index, val)
		return true
	})

	return New(collected)
}

/*
	Complex64 (complex64 and []complex64)
	--------------------------------------------------
*/

// Complex64 gets the value as a complex64, returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *Obj) Complex64(optionalDefault ...complex64) complex64 {
	if s, ok := o.obj.(complex64); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// MustComplex64 gets the value as a complex64.
//
// Panics if the object is not a complex64.
func (o *Obj) MustComplex64() complex64 {
	return o.obj.(complex64)
}

// Complex64Slice gets the value as a []complex64, returns the optionalDefault
// value or nil if the value is not a []complex64.
func (o *Obj) Complex64Slice(optionalDefault ...[]complex64) []complex64 {
	if s, ok := o.obj.([]complex64); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustComplex64Slice gets the value as a []complex64.
//
// Panics if the object is not a []complex64.
func (o *Obj) MustComplex64Slice() []complex64 {
	return o.obj.([]complex64)
}

// IsComplex64 gets whether the object contained is a complex64 or not.
func (o *Obj) IsComplex64() bool {
	return o.IsKind(reflect.Complex64)
}

// IsComplex64Slice gets whether the object contained is a []complex64 or not.
func (o *Obj) IsComplex64Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]complex64)
	return ok
}

// EachComplex64 calls the specified callback for each object
// in the []complex64.
//
// Panics if the object is the wrong type.
func (o *Obj) EachComplex64(callback func(int, complex64) bool) *Obj {

	for index, val := range o.MustComplex64Slice() {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

// WhereComplex64 uses the specified decider function to select items
// from the []complex64.  The object contained in the result will contain
// only the selected items.
func (o *Obj) WhereComplex64(decider func(int, complex64) bool) *Obj {

	var selected []complex64

	o.EachComplex64(func(index int, val complex64) bool {
		shouldSelect := decider(index, val)
		if shouldSelect == false {
			selected = append(selected, val)
		}
		return true
	})

	return New(selected)

}

// GroupComplex64 uses the specified grouper function to group the items
// keyed by the return of the grouper.  The object contained in the
// result will contain a map[string][]complex64.
func (o *Obj) GroupComplex64(grouper func(int, complex64) string) *Obj {

	groups := make(map[string][]complex64)

	o.EachComplex64(func(index int, val complex64) bool {
		group := grouper(index, val)
		if _, ok := groups[group]; !ok {
			groups[group] = make([]complex64, 0)
		}
		groups[group] = append(groups[group], val)
		return true
	})

	return New(groups)

}

// ReplaceComplex64 uses the specified function to replace each complex64s
// by iterating each item.  The data in the returned result will be a
// []complex64 containing the replaced items.
func (o *Obj) ReplaceComplex64(replacer func(int, complex64) complex64) *Obj {

	arr := o.MustComplex64Slice()
	replaced := make([]complex64, len(arr))

	o.EachComplex64(func(index int, val complex64) bool {
		replaced[index] = replacer(index, val)
		return true
	})

	return New(replaced)

}

// CollectComplex64 uses the specified collector function to collect a value
// for each of the complex64s in the slice.  The data returned will be a
// []interface{}.
func (o *Obj) CollectComplex64(collector func(int, complex64) interface{}) *Obj {

	arr := o.MustComplex64Slice()
	collected := make([]interface{}, len(arr))

	o.EachComplex64(func(index int, val complex64) bool {
		collected[index] = collector(index, val)
		return true
	})

	return New(collected)
}

/*
	Complex128 (complex128 and []complex128)
	--------------------------------------------------
*/

// Complex128 gets the value as a complex128, returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *Obj) Complex128(optionalDefault ...complex128) complex128 {
	if s, ok := o.obj.(complex128); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return 0
}

// MustComplex128 gets the value as a complex128.
//
// Panics if the object is not a complex128.
func (o *Obj) MustComplex128() complex128 {
	return o.obj.(complex128)
}

// Complex128Slice gets the value as a []complex128, returns the optionalDefault
// value or nil if the value is not a []complex128.
func (o *Obj) Complex128Slice(optionalDefault ...[]complex128) []complex128 {
	if s, ok := o.obj.([]complex128); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustComplex128Slice gets the value as a []complex128.
//
// Panics if the object is not a []complex128.
func (o *Obj) MustComplex128Slice() []complex128 {
	return o.obj.([]complex128)
}

// IsComplex128 gets whether the object contained is a complex128 or not.
func (o *Obj) IsComplex128() bool {
	return o.IsKind(reflect.Complex128)
}

// IsComplex128Slice gets whether the object contained is a []complex128 or not.
func (o *Obj) IsComplex128Slice() bool {
	if !o.IsSlice() {
		return false
	}
	_, ok := o.obj.([]complex128)
	return ok
}

// EachComplex128 calls the specified callback for each object
// in the []complex128.
//
// Panics if the object is the wrong type.
func (o *Obj) EachComplex128(callback func(int, complex128) bool) *Obj {

	for index, val := range o.MustComplex128Slice() {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

// WhereComplex128 uses the specified decider function to select items
// from the []complex128.  The object contained in the result will contain
// only the selected items.
func (o *Obj) WhereComplex128(decider func(int, complex128) bool) *Obj {

	var selected []complex128

	o.EachComplex128(func(index int, val complex128) bool {
		shouldSelect := decider(index, val)
		if shouldSelect == false {
			selected = append(selected, val)
		}
		return true
	})

	return New(selected)

}

// GroupComplex128 uses the specified grouper function to group the items
// keyed by the return of the grouper.  The object contained in the
// result will contain a map[string][]complex128.
func (o *Obj) GroupComplex128(grouper func(int, complex128) string) *Obj {

	groups := make(map[string][]complex128)

	o.EachComplex128(func(index int, val complex128) bool {
		group := grouper(index, val)
		if _, ok := groups[group]; !ok {
			groups[group] = make([]complex128, 0)
		}
		groups[group] = append(groups[group], val)
		return true
	})

	return New(groups)

}

// ReplaceComplex128 uses the specified function to replace each complex128s
// by iterating each item.  The data in the returned result will be a
// []complex128 containing the replaced items.
func (o *Obj) ReplaceComplex128(replacer func(int, complex128) complex128) *Obj {

	arr := o.MustComplex128Slice()
	replaced := make([]complex128, len(arr))

	o.EachComplex128(func(index int, val complex128) bool {
		replaced[index] = replacer(index, val)
		return true
	})

	return New(replaced)

}

// CollectComplex128 uses the specified collector function to collect a value
// for each of the complex128s in the slice.  The data returned will be a
// []interface{}.
func (o *Obj) CollectComplex128(collector func(int, complex128) interface{}) *Obj {

	arr := o.MustComplex128Slice()
	collected := make([]interface{}, len(arr))

	o.EachComplex128(func(index int, val complex128) bool {
		collected[index] = collector(index, val)
		return true
	})

	return New(collected)
}
