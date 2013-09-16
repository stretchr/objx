package objx

/*
	Inter (interface{} and []interface{})
	--------------------------------------------------
*/

// Inter gets the value as a interface{}, returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *Objx) Inter(optionalDefault ...interface{}) interface{} {
	if s, ok := o.data.(interface{}); ok {
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
func (o *Objx) MustInter() interface{} {
	return o.data.(interface{})
}

// InterSlice gets the value as a []interface{}, returns the optionalDefault
// value or nil if the value is not a []interface{}.
func (o *Objx) InterSlice(optionalDefault ...[]interface{}) []interface{} {
	if s, ok := o.data.([]interface{}); ok {
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
func (o *Objx) MustInterSlice() []interface{} {
	return o.data.([]interface{})
}

// IsInter gets whether the object contained is a interface{} or not.
func (o *Objx) IsInter() bool {
	_, ok := o.data.(interface{})
	return ok
}

// IsInterSlice gets whether the object contained is a []interface{} or not.
func (o *Objx) IsInterSlice() bool {
	_, ok := o.data.([]interface{})
	return ok
}

// EachInter calls the specified callback for each object
// in the []interface{}.
//
// Panics if the object is the wrong type.
func (o *Objx) EachInter(callback func(int, interface{}) bool) *Objx {

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
func (o *Objx) WhereInter(decider func(int, interface{}) bool) *Objx {

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
func (o *Objx) GroupInter(grouper func(int, interface{}) string) *Objx {

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
func (o *Objx) ReplaceInter(replacer func(int, interface{}) interface{}) *Objx {

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
func (o *Objx) CollectInter(collector func(int, interface{}) interface{}) *Objx {

	arr := o.MustInterSlice()
	collected := make([]interface{}, len(arr))

	o.EachInter(func(index int, val interface{}) bool {
		collected[index] = collector(index, val)
		return true
	})

	return New(collected)
}

/*
	MSI (map[string]interface{} and []map[string]interface{})
	--------------------------------------------------
*/

// MSI gets the value as a map[string]interface{}, returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *Objx) MSI(optionalDefault ...map[string]interface{}) map[string]interface{} {
	if s, ok := o.data.(map[string]interface{}); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustMSI gets the value as a map[string]interface{}.
//
// Panics if the object is not a map[string]interface{}.
func (o *Objx) MustMSI() map[string]interface{} {
	return o.data.(map[string]interface{})
}

// MSISlice gets the value as a []map[string]interface{}, returns the optionalDefault
// value or nil if the value is not a []map[string]interface{}.
func (o *Objx) MSISlice(optionalDefault ...[]map[string]interface{}) []map[string]interface{} {
	if s, ok := o.data.([]map[string]interface{}); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustMSISlice gets the value as a []map[string]interface{}.
//
// Panics if the object is not a []map[string]interface{}.
func (o *Objx) MustMSISlice() []map[string]interface{} {
	return o.data.([]map[string]interface{})
}

// IsMSI gets whether the object contained is a map[string]interface{} or not.
func (o *Objx) IsMSI() bool {
	_, ok := o.data.(map[string]interface{})
	return ok
}

// IsMSISlice gets whether the object contained is a []map[string]interface{} or not.
func (o *Objx) IsMSISlice() bool {
	_, ok := o.data.([]map[string]interface{})
	return ok
}

// EachMSI calls the specified callback for each object
// in the []map[string]interface{}.
//
// Panics if the object is the wrong type.
func (o *Objx) EachMSI(callback func(int, map[string]interface{}) bool) *Objx {

	for index, val := range o.MustMSISlice() {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

// WhereMSI uses the specified decider function to select items
// from the []map[string]interface{}.  The object contained in the result will contain
// only the selected items.
func (o *Objx) WhereMSI(decider func(int, map[string]interface{}) bool) *Objx {

	var selected []map[string]interface{}

	o.EachMSI(func(index int, val map[string]interface{}) bool {
		shouldSelect := decider(index, val)
		if shouldSelect == false {
			selected = append(selected, val)
		}
		return true
	})

	return New(selected)

}

// GroupMSI uses the specified grouper function to group the items
// keyed by the return of the grouper.  The object contained in the
// result will contain a map[string][]map[string]interface{}.
func (o *Objx) GroupMSI(grouper func(int, map[string]interface{}) string) *Objx {

	groups := make(map[string][]map[string]interface{})

	o.EachMSI(func(index int, val map[string]interface{}) bool {
		group := grouper(index, val)
		if _, ok := groups[group]; !ok {
			groups[group] = make([]map[string]interface{}, 0)
		}
		groups[group] = append(groups[group], val)
		return true
	})

	return New(groups)

}

// ReplaceMSI uses the specified function to replace each map[string]interface{}s
// by iterating each item.  The data in the returned result will be a
// []map[string]interface{} containing the replaced items.
func (o *Objx) ReplaceMSI(replacer func(int, map[string]interface{}) map[string]interface{}) *Objx {

	arr := o.MustMSISlice()
	replaced := make([]map[string]interface{}, len(arr))

	o.EachMSI(func(index int, val map[string]interface{}) bool {
		replaced[index] = replacer(index, val)
		return true
	})

	return New(replaced)

}

// CollectMSI uses the specified collector function to collect a value
// for each of the map[string]interface{}s in the slice.  The data returned will be a
// []interface{}.
func (o *Objx) CollectMSI(collector func(int, map[string]interface{}) interface{}) *Objx {

	arr := o.MustMSISlice()
	collected := make([]interface{}, len(arr))

	o.EachMSI(func(index int, val map[string]interface{}) bool {
		collected[index] = collector(index, val)
		return true
	})

	return New(collected)
}

/*
	ObjxObjx ((*Objx) and [](*Objx))
	--------------------------------------------------
*/

// ObjxObjx gets the value as a (*Objx), returns the optionalDefault
// value or a system default object if the value is the wrong type.
func (o *Objx) ObjxObjx(optionalDefault ...(*Objx)) *Objx {
	if s, ok := o.data.((*Objx)); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return New(nil)
}

// MustObjxObjx gets the value as a (*Objx).
//
// Panics if the object is not a (*Objx).
func (o *Objx) MustObjxObjx() *Objx {
	return o.data.((*Objx))
}

// ObjxObjxSlice gets the value as a [](*Objx), returns the optionalDefault
// value or nil if the value is not a [](*Objx).
func (o *Objx) ObjxObjxSlice(optionalDefault ...[](*Objx)) [](*Objx) {
	if s, ok := o.data.([](*Objx)); ok {
		return s
	}
	if len(optionalDefault) == 1 {
		return optionalDefault[0]
	}
	return nil
}

// MustObjxObjxSlice gets the value as a [](*Objx).
//
// Panics if the object is not a [](*Objx).
func (o *Objx) MustObjxObjxSlice() [](*Objx) {
	return o.data.([](*Objx))
}

// IsObjxObjx gets whether the object contained is a (*Objx) or not.
func (o *Objx) IsObjxObjx() bool {
	_, ok := o.data.((*Objx))
	return ok
}

// IsObjxObjxSlice gets whether the object contained is a [](*Objx) or not.
func (o *Objx) IsObjxObjxSlice() bool {
	_, ok := o.data.([](*Objx))
	return ok
}

// EachObjxObjx calls the specified callback for each object
// in the [](*Objx).
//
// Panics if the object is the wrong type.
func (o *Objx) EachObjxObjx(callback func(int, *Objx) bool) *Objx {

	for index, val := range o.MustObjxObjxSlice() {
		carryon := callback(index, val)
		if carryon == false {
			break
		}
	}

	return o

}

// WhereObjxObjx uses the specified decider function to select items
// from the [](*Objx).  The object contained in the result will contain
// only the selected items.
func (o *Objx) WhereObjxObjx(decider func(int, *Objx) bool) *Objx {

	var selected [](*Objx)

	o.EachObjxObjx(func(index int, val *Objx) bool {
		shouldSelect := decider(index, val)
		if shouldSelect == false {
			selected = append(selected, val)
		}
		return true
	})

	return New(selected)

}

// GroupObjxObjx uses the specified grouper function to group the items
// keyed by the return of the grouper.  The object contained in the
// result will contain a map[string][](*Objx).
func (o *Objx) GroupObjxObjx(grouper func(int, *Objx) string) *Objx {

	groups := make(map[string][](*Objx))

	o.EachObjxObjx(func(index int, val *Objx) bool {
		group := grouper(index, val)
		if _, ok := groups[group]; !ok {
			groups[group] = make([](*Objx), 0)
		}
		groups[group] = append(groups[group], val)
		return true
	})

	return New(groups)

}

// ReplaceObjxObjx uses the specified function to replace each (*Objx)s
// by iterating each item.  The data in the returned result will be a
// [](*Objx) containing the replaced items.
func (o *Objx) ReplaceObjxObjx(replacer func(int, *Objx) *Objx) *Objx {

	arr := o.MustObjxObjxSlice()
	replaced := make([](*Objx), len(arr))

	o.EachObjxObjx(func(index int, val *Objx) bool {
		replaced[index] = replacer(index, val)
		return true
	})

	return New(replaced)

}

// CollectObjxObjx uses the specified collector function to collect a value
// for each of the (*Objx)s in the slice.  The data returned will be a
// []interface{}.
func (o *Objx) CollectObjxObjx(collector func(int, *Objx) interface{}) *Objx {

	arr := o.MustObjxObjxSlice()
	collected := make([]interface{}, len(arr))

	o.EachObjxObjx(func(index int, val *Objx) bool {
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
func (o *Objx) Bool(optionalDefault ...bool) bool {
	if s, ok := o.data.(bool); ok {
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
func (o *Objx) MustBool() bool {
	return o.data.(bool)
}

// BoolSlice gets the value as a []bool, returns the optionalDefault
// value or nil if the value is not a []bool.
func (o *Objx) BoolSlice(optionalDefault ...[]bool) []bool {
	if s, ok := o.data.([]bool); ok {
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
func (o *Objx) MustBoolSlice() []bool {
	return o.data.([]bool)
}

// IsBool gets whether the object contained is a bool or not.
func (o *Objx) IsBool() bool {
	_, ok := o.data.(bool)
	return ok
}

// IsBoolSlice gets whether the object contained is a []bool or not.
func (o *Objx) IsBoolSlice() bool {
	_, ok := o.data.([]bool)
	return ok
}

// EachBool calls the specified callback for each object
// in the []bool.
//
// Panics if the object is the wrong type.
func (o *Objx) EachBool(callback func(int, bool) bool) *Objx {

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
func (o *Objx) WhereBool(decider func(int, bool) bool) *Objx {

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
func (o *Objx) GroupBool(grouper func(int, bool) string) *Objx {

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
func (o *Objx) ReplaceBool(replacer func(int, bool) bool) *Objx {

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
func (o *Objx) CollectBool(collector func(int, bool) interface{}) *Objx {

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
func (o *Objx) Str(optionalDefault ...string) string {
	if s, ok := o.data.(string); ok {
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
func (o *Objx) MustStr() string {
	return o.data.(string)
}

// StrSlice gets the value as a []string, returns the optionalDefault
// value or nil if the value is not a []string.
func (o *Objx) StrSlice(optionalDefault ...[]string) []string {
	if s, ok := o.data.([]string); ok {
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
func (o *Objx) MustStrSlice() []string {
	return o.data.([]string)
}

// IsStr gets whether the object contained is a string or not.
func (o *Objx) IsStr() bool {
	_, ok := o.data.(string)
	return ok
}

// IsStrSlice gets whether the object contained is a []string or not.
func (o *Objx) IsStrSlice() bool {
	_, ok := o.data.([]string)
	return ok
}

// EachStr calls the specified callback for each object
// in the []string.
//
// Panics if the object is the wrong type.
func (o *Objx) EachStr(callback func(int, string) bool) *Objx {

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
func (o *Objx) WhereStr(decider func(int, string) bool) *Objx {

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
func (o *Objx) GroupStr(grouper func(int, string) string) *Objx {

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
func (o *Objx) ReplaceStr(replacer func(int, string) string) *Objx {

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
func (o *Objx) CollectStr(collector func(int, string) interface{}) *Objx {

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
func (o *Objx) Int(optionalDefault ...int) int {
	if s, ok := o.data.(int); ok {
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
func (o *Objx) MustInt() int {
	return o.data.(int)
}

// IntSlice gets the value as a []int, returns the optionalDefault
// value or nil if the value is not a []int.
func (o *Objx) IntSlice(optionalDefault ...[]int) []int {
	if s, ok := o.data.([]int); ok {
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
func (o *Objx) MustIntSlice() []int {
	return o.data.([]int)
}

// IsInt gets whether the object contained is a int or not.
func (o *Objx) IsInt() bool {
	_, ok := o.data.(int)
	return ok
}

// IsIntSlice gets whether the object contained is a []int or not.
func (o *Objx) IsIntSlice() bool {
	_, ok := o.data.([]int)
	return ok
}

// EachInt calls the specified callback for each object
// in the []int.
//
// Panics if the object is the wrong type.
func (o *Objx) EachInt(callback func(int, int) bool) *Objx {

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
func (o *Objx) WhereInt(decider func(int, int) bool) *Objx {

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
func (o *Objx) GroupInt(grouper func(int, int) string) *Objx {

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
func (o *Objx) ReplaceInt(replacer func(int, int) int) *Objx {

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
func (o *Objx) CollectInt(collector func(int, int) interface{}) *Objx {

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
func (o *Objx) Int8(optionalDefault ...int8) int8 {
	if s, ok := o.data.(int8); ok {
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
func (o *Objx) MustInt8() int8 {
	return o.data.(int8)
}

// Int8Slice gets the value as a []int8, returns the optionalDefault
// value or nil if the value is not a []int8.
func (o *Objx) Int8Slice(optionalDefault ...[]int8) []int8 {
	if s, ok := o.data.([]int8); ok {
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
func (o *Objx) MustInt8Slice() []int8 {
	return o.data.([]int8)
}

// IsInt8 gets whether the object contained is a int8 or not.
func (o *Objx) IsInt8() bool {
	_, ok := o.data.(int8)
	return ok
}

// IsInt8Slice gets whether the object contained is a []int8 or not.
func (o *Objx) IsInt8Slice() bool {
	_, ok := o.data.([]int8)
	return ok
}

// EachInt8 calls the specified callback for each object
// in the []int8.
//
// Panics if the object is the wrong type.
func (o *Objx) EachInt8(callback func(int, int8) bool) *Objx {

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
func (o *Objx) WhereInt8(decider func(int, int8) bool) *Objx {

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
func (o *Objx) GroupInt8(grouper func(int, int8) string) *Objx {

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
func (o *Objx) ReplaceInt8(replacer func(int, int8) int8) *Objx {

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
func (o *Objx) CollectInt8(collector func(int, int8) interface{}) *Objx {

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
func (o *Objx) Int16(optionalDefault ...int16) int16 {
	if s, ok := o.data.(int16); ok {
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
func (o *Objx) MustInt16() int16 {
	return o.data.(int16)
}

// Int16Slice gets the value as a []int16, returns the optionalDefault
// value or nil if the value is not a []int16.
func (o *Objx) Int16Slice(optionalDefault ...[]int16) []int16 {
	if s, ok := o.data.([]int16); ok {
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
func (o *Objx) MustInt16Slice() []int16 {
	return o.data.([]int16)
}

// IsInt16 gets whether the object contained is a int16 or not.
func (o *Objx) IsInt16() bool {
	_, ok := o.data.(int16)
	return ok
}

// IsInt16Slice gets whether the object contained is a []int16 or not.
func (o *Objx) IsInt16Slice() bool {
	_, ok := o.data.([]int16)
	return ok
}

// EachInt16 calls the specified callback for each object
// in the []int16.
//
// Panics if the object is the wrong type.
func (o *Objx) EachInt16(callback func(int, int16) bool) *Objx {

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
func (o *Objx) WhereInt16(decider func(int, int16) bool) *Objx {

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
func (o *Objx) GroupInt16(grouper func(int, int16) string) *Objx {

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
func (o *Objx) ReplaceInt16(replacer func(int, int16) int16) *Objx {

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
func (o *Objx) CollectInt16(collector func(int, int16) interface{}) *Objx {

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
func (o *Objx) Int32(optionalDefault ...int32) int32 {
	if s, ok := o.data.(int32); ok {
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
func (o *Objx) MustInt32() int32 {
	return o.data.(int32)
}

// Int32Slice gets the value as a []int32, returns the optionalDefault
// value or nil if the value is not a []int32.
func (o *Objx) Int32Slice(optionalDefault ...[]int32) []int32 {
	if s, ok := o.data.([]int32); ok {
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
func (o *Objx) MustInt32Slice() []int32 {
	return o.data.([]int32)
}

// IsInt32 gets whether the object contained is a int32 or not.
func (o *Objx) IsInt32() bool {
	_, ok := o.data.(int32)
	return ok
}

// IsInt32Slice gets whether the object contained is a []int32 or not.
func (o *Objx) IsInt32Slice() bool {
	_, ok := o.data.([]int32)
	return ok
}

// EachInt32 calls the specified callback for each object
// in the []int32.
//
// Panics if the object is the wrong type.
func (o *Objx) EachInt32(callback func(int, int32) bool) *Objx {

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
func (o *Objx) WhereInt32(decider func(int, int32) bool) *Objx {

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
func (o *Objx) GroupInt32(grouper func(int, int32) string) *Objx {

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
func (o *Objx) ReplaceInt32(replacer func(int, int32) int32) *Objx {

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
func (o *Objx) CollectInt32(collector func(int, int32) interface{}) *Objx {

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
func (o *Objx) Int64(optionalDefault ...int64) int64 {
	if s, ok := o.data.(int64); ok {
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
func (o *Objx) MustInt64() int64 {
	return o.data.(int64)
}

// Int64Slice gets the value as a []int64, returns the optionalDefault
// value or nil if the value is not a []int64.
func (o *Objx) Int64Slice(optionalDefault ...[]int64) []int64 {
	if s, ok := o.data.([]int64); ok {
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
func (o *Objx) MustInt64Slice() []int64 {
	return o.data.([]int64)
}

// IsInt64 gets whether the object contained is a int64 or not.
func (o *Objx) IsInt64() bool {
	_, ok := o.data.(int64)
	return ok
}

// IsInt64Slice gets whether the object contained is a []int64 or not.
func (o *Objx) IsInt64Slice() bool {
	_, ok := o.data.([]int64)
	return ok
}

// EachInt64 calls the specified callback for each object
// in the []int64.
//
// Panics if the object is the wrong type.
func (o *Objx) EachInt64(callback func(int, int64) bool) *Objx {

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
func (o *Objx) WhereInt64(decider func(int, int64) bool) *Objx {

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
func (o *Objx) GroupInt64(grouper func(int, int64) string) *Objx {

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
func (o *Objx) ReplaceInt64(replacer func(int, int64) int64) *Objx {

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
func (o *Objx) CollectInt64(collector func(int, int64) interface{}) *Objx {

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
func (o *Objx) Uint(optionalDefault ...uint) uint {
	if s, ok := o.data.(uint); ok {
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
func (o *Objx) MustUint() uint {
	return o.data.(uint)
}

// UintSlice gets the value as a []uint, returns the optionalDefault
// value or nil if the value is not a []uint.
func (o *Objx) UintSlice(optionalDefault ...[]uint) []uint {
	if s, ok := o.data.([]uint); ok {
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
func (o *Objx) MustUintSlice() []uint {
	return o.data.([]uint)
}

// IsUint gets whether the object contained is a uint or not.
func (o *Objx) IsUint() bool {
	_, ok := o.data.(uint)
	return ok
}

// IsUintSlice gets whether the object contained is a []uint or not.
func (o *Objx) IsUintSlice() bool {
	_, ok := o.data.([]uint)
	return ok
}

// EachUint calls the specified callback for each object
// in the []uint.
//
// Panics if the object is the wrong type.
func (o *Objx) EachUint(callback func(int, uint) bool) *Objx {

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
func (o *Objx) WhereUint(decider func(int, uint) bool) *Objx {

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
func (o *Objx) GroupUint(grouper func(int, uint) string) *Objx {

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
func (o *Objx) ReplaceUint(replacer func(int, uint) uint) *Objx {

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
func (o *Objx) CollectUint(collector func(int, uint) interface{}) *Objx {

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
func (o *Objx) Uint8(optionalDefault ...uint8) uint8 {
	if s, ok := o.data.(uint8); ok {
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
func (o *Objx) MustUint8() uint8 {
	return o.data.(uint8)
}

// Uint8Slice gets the value as a []uint8, returns the optionalDefault
// value or nil if the value is not a []uint8.
func (o *Objx) Uint8Slice(optionalDefault ...[]uint8) []uint8 {
	if s, ok := o.data.([]uint8); ok {
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
func (o *Objx) MustUint8Slice() []uint8 {
	return o.data.([]uint8)
}

// IsUint8 gets whether the object contained is a uint8 or not.
func (o *Objx) IsUint8() bool {
	_, ok := o.data.(uint8)
	return ok
}

// IsUint8Slice gets whether the object contained is a []uint8 or not.
func (o *Objx) IsUint8Slice() bool {
	_, ok := o.data.([]uint8)
	return ok
}

// EachUint8 calls the specified callback for each object
// in the []uint8.
//
// Panics if the object is the wrong type.
func (o *Objx) EachUint8(callback func(int, uint8) bool) *Objx {

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
func (o *Objx) WhereUint8(decider func(int, uint8) bool) *Objx {

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
func (o *Objx) GroupUint8(grouper func(int, uint8) string) *Objx {

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
func (o *Objx) ReplaceUint8(replacer func(int, uint8) uint8) *Objx {

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
func (o *Objx) CollectUint8(collector func(int, uint8) interface{}) *Objx {

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
func (o *Objx) Uint16(optionalDefault ...uint16) uint16 {
	if s, ok := o.data.(uint16); ok {
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
func (o *Objx) MustUint16() uint16 {
	return o.data.(uint16)
}

// Uint16Slice gets the value as a []uint16, returns the optionalDefault
// value or nil if the value is not a []uint16.
func (o *Objx) Uint16Slice(optionalDefault ...[]uint16) []uint16 {
	if s, ok := o.data.([]uint16); ok {
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
func (o *Objx) MustUint16Slice() []uint16 {
	return o.data.([]uint16)
}

// IsUint16 gets whether the object contained is a uint16 or not.
func (o *Objx) IsUint16() bool {
	_, ok := o.data.(uint16)
	return ok
}

// IsUint16Slice gets whether the object contained is a []uint16 or not.
func (o *Objx) IsUint16Slice() bool {
	_, ok := o.data.([]uint16)
	return ok
}

// EachUint16 calls the specified callback for each object
// in the []uint16.
//
// Panics if the object is the wrong type.
func (o *Objx) EachUint16(callback func(int, uint16) bool) *Objx {

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
func (o *Objx) WhereUint16(decider func(int, uint16) bool) *Objx {

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
func (o *Objx) GroupUint16(grouper func(int, uint16) string) *Objx {

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
func (o *Objx) ReplaceUint16(replacer func(int, uint16) uint16) *Objx {

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
func (o *Objx) CollectUint16(collector func(int, uint16) interface{}) *Objx {

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
func (o *Objx) Uint32(optionalDefault ...uint32) uint32 {
	if s, ok := o.data.(uint32); ok {
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
func (o *Objx) MustUint32() uint32 {
	return o.data.(uint32)
}

// Uint32Slice gets the value as a []uint32, returns the optionalDefault
// value or nil if the value is not a []uint32.
func (o *Objx) Uint32Slice(optionalDefault ...[]uint32) []uint32 {
	if s, ok := o.data.([]uint32); ok {
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
func (o *Objx) MustUint32Slice() []uint32 {
	return o.data.([]uint32)
}

// IsUint32 gets whether the object contained is a uint32 or not.
func (o *Objx) IsUint32() bool {
	_, ok := o.data.(uint32)
	return ok
}

// IsUint32Slice gets whether the object contained is a []uint32 or not.
func (o *Objx) IsUint32Slice() bool {
	_, ok := o.data.([]uint32)
	return ok
}

// EachUint32 calls the specified callback for each object
// in the []uint32.
//
// Panics if the object is the wrong type.
func (o *Objx) EachUint32(callback func(int, uint32) bool) *Objx {

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
func (o *Objx) WhereUint32(decider func(int, uint32) bool) *Objx {

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
func (o *Objx) GroupUint32(grouper func(int, uint32) string) *Objx {

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
func (o *Objx) ReplaceUint32(replacer func(int, uint32) uint32) *Objx {

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
func (o *Objx) CollectUint32(collector func(int, uint32) interface{}) *Objx {

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
func (o *Objx) Uint64(optionalDefault ...uint64) uint64 {
	if s, ok := o.data.(uint64); ok {
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
func (o *Objx) MustUint64() uint64 {
	return o.data.(uint64)
}

// Uint64Slice gets the value as a []uint64, returns the optionalDefault
// value or nil if the value is not a []uint64.
func (o *Objx) Uint64Slice(optionalDefault ...[]uint64) []uint64 {
	if s, ok := o.data.([]uint64); ok {
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
func (o *Objx) MustUint64Slice() []uint64 {
	return o.data.([]uint64)
}

// IsUint64 gets whether the object contained is a uint64 or not.
func (o *Objx) IsUint64() bool {
	_, ok := o.data.(uint64)
	return ok
}

// IsUint64Slice gets whether the object contained is a []uint64 or not.
func (o *Objx) IsUint64Slice() bool {
	_, ok := o.data.([]uint64)
	return ok
}

// EachUint64 calls the specified callback for each object
// in the []uint64.
//
// Panics if the object is the wrong type.
func (o *Objx) EachUint64(callback func(int, uint64) bool) *Objx {

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
func (o *Objx) WhereUint64(decider func(int, uint64) bool) *Objx {

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
func (o *Objx) GroupUint64(grouper func(int, uint64) string) *Objx {

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
func (o *Objx) ReplaceUint64(replacer func(int, uint64) uint64) *Objx {

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
func (o *Objx) CollectUint64(collector func(int, uint64) interface{}) *Objx {

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
func (o *Objx) Uintptr(optionalDefault ...uintptr) uintptr {
	if s, ok := o.data.(uintptr); ok {
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
func (o *Objx) MustUintptr() uintptr {
	return o.data.(uintptr)
}

// UintptrSlice gets the value as a []uintptr, returns the optionalDefault
// value or nil if the value is not a []uintptr.
func (o *Objx) UintptrSlice(optionalDefault ...[]uintptr) []uintptr {
	if s, ok := o.data.([]uintptr); ok {
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
func (o *Objx) MustUintptrSlice() []uintptr {
	return o.data.([]uintptr)
}

// IsUintptr gets whether the object contained is a uintptr or not.
func (o *Objx) IsUintptr() bool {
	_, ok := o.data.(uintptr)
	return ok
}

// IsUintptrSlice gets whether the object contained is a []uintptr or not.
func (o *Objx) IsUintptrSlice() bool {
	_, ok := o.data.([]uintptr)
	return ok
}

// EachUintptr calls the specified callback for each object
// in the []uintptr.
//
// Panics if the object is the wrong type.
func (o *Objx) EachUintptr(callback func(int, uintptr) bool) *Objx {

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
func (o *Objx) WhereUintptr(decider func(int, uintptr) bool) *Objx {

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
func (o *Objx) GroupUintptr(grouper func(int, uintptr) string) *Objx {

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
func (o *Objx) ReplaceUintptr(replacer func(int, uintptr) uintptr) *Objx {

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
func (o *Objx) CollectUintptr(collector func(int, uintptr) interface{}) *Objx {

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
func (o *Objx) Float32(optionalDefault ...float32) float32 {
	if s, ok := o.data.(float32); ok {
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
func (o *Objx) MustFloat32() float32 {
	return o.data.(float32)
}

// Float32Slice gets the value as a []float32, returns the optionalDefault
// value or nil if the value is not a []float32.
func (o *Objx) Float32Slice(optionalDefault ...[]float32) []float32 {
	if s, ok := o.data.([]float32); ok {
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
func (o *Objx) MustFloat32Slice() []float32 {
	return o.data.([]float32)
}

// IsFloat32 gets whether the object contained is a float32 or not.
func (o *Objx) IsFloat32() bool {
	_, ok := o.data.(float32)
	return ok
}

// IsFloat32Slice gets whether the object contained is a []float32 or not.
func (o *Objx) IsFloat32Slice() bool {
	_, ok := o.data.([]float32)
	return ok
}

// EachFloat32 calls the specified callback for each object
// in the []float32.
//
// Panics if the object is the wrong type.
func (o *Objx) EachFloat32(callback func(int, float32) bool) *Objx {

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
func (o *Objx) WhereFloat32(decider func(int, float32) bool) *Objx {

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
func (o *Objx) GroupFloat32(grouper func(int, float32) string) *Objx {

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
func (o *Objx) ReplaceFloat32(replacer func(int, float32) float32) *Objx {

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
func (o *Objx) CollectFloat32(collector func(int, float32) interface{}) *Objx {

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
func (o *Objx) Float64(optionalDefault ...float64) float64 {
	if s, ok := o.data.(float64); ok {
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
func (o *Objx) MustFloat64() float64 {
	return o.data.(float64)
}

// Float64Slice gets the value as a []float64, returns the optionalDefault
// value or nil if the value is not a []float64.
func (o *Objx) Float64Slice(optionalDefault ...[]float64) []float64 {
	if s, ok := o.data.([]float64); ok {
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
func (o *Objx) MustFloat64Slice() []float64 {
	return o.data.([]float64)
}

// IsFloat64 gets whether the object contained is a float64 or not.
func (o *Objx) IsFloat64() bool {
	_, ok := o.data.(float64)
	return ok
}

// IsFloat64Slice gets whether the object contained is a []float64 or not.
func (o *Objx) IsFloat64Slice() bool {
	_, ok := o.data.([]float64)
	return ok
}

// EachFloat64 calls the specified callback for each object
// in the []float64.
//
// Panics if the object is the wrong type.
func (o *Objx) EachFloat64(callback func(int, float64) bool) *Objx {

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
func (o *Objx) WhereFloat64(decider func(int, float64) bool) *Objx {

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
func (o *Objx) GroupFloat64(grouper func(int, float64) string) *Objx {

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
func (o *Objx) ReplaceFloat64(replacer func(int, float64) float64) *Objx {

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
func (o *Objx) CollectFloat64(collector func(int, float64) interface{}) *Objx {

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
func (o *Objx) Complex64(optionalDefault ...complex64) complex64 {
	if s, ok := o.data.(complex64); ok {
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
func (o *Objx) MustComplex64() complex64 {
	return o.data.(complex64)
}

// Complex64Slice gets the value as a []complex64, returns the optionalDefault
// value or nil if the value is not a []complex64.
func (o *Objx) Complex64Slice(optionalDefault ...[]complex64) []complex64 {
	if s, ok := o.data.([]complex64); ok {
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
func (o *Objx) MustComplex64Slice() []complex64 {
	return o.data.([]complex64)
}

// IsComplex64 gets whether the object contained is a complex64 or not.
func (o *Objx) IsComplex64() bool {
	_, ok := o.data.(complex64)
	return ok
}

// IsComplex64Slice gets whether the object contained is a []complex64 or not.
func (o *Objx) IsComplex64Slice() bool {
	_, ok := o.data.([]complex64)
	return ok
}

// EachComplex64 calls the specified callback for each object
// in the []complex64.
//
// Panics if the object is the wrong type.
func (o *Objx) EachComplex64(callback func(int, complex64) bool) *Objx {

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
func (o *Objx) WhereComplex64(decider func(int, complex64) bool) *Objx {

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
func (o *Objx) GroupComplex64(grouper func(int, complex64) string) *Objx {

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
func (o *Objx) ReplaceComplex64(replacer func(int, complex64) complex64) *Objx {

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
func (o *Objx) CollectComplex64(collector func(int, complex64) interface{}) *Objx {

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
func (o *Objx) Complex128(optionalDefault ...complex128) complex128 {
	if s, ok := o.data.(complex128); ok {
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
func (o *Objx) MustComplex128() complex128 {
	return o.data.(complex128)
}

// Complex128Slice gets the value as a []complex128, returns the optionalDefault
// value or nil if the value is not a []complex128.
func (o *Objx) Complex128Slice(optionalDefault ...[]complex128) []complex128 {
	if s, ok := o.data.([]complex128); ok {
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
func (o *Objx) MustComplex128Slice() []complex128 {
	return o.data.([]complex128)
}

// IsComplex128 gets whether the object contained is a complex128 or not.
func (o *Objx) IsComplex128() bool {
	_, ok := o.data.(complex128)
	return ok
}

// IsComplex128Slice gets whether the object contained is a []complex128 or not.
func (o *Objx) IsComplex128Slice() bool {
	_, ok := o.data.([]complex128)
	return ok
}

// EachComplex128 calls the specified callback for each object
// in the []complex128.
//
// Panics if the object is the wrong type.
func (o *Objx) EachComplex128(callback func(int, complex128) bool) *Objx {

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
func (o *Objx) WhereComplex128(decider func(int, complex128) bool) *Objx {

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
func (o *Objx) GroupComplex128(grouper func(int, complex128) string) *Objx {

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
func (o *Objx) ReplaceComplex128(replacer func(int, complex128) complex128) *Objx {

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
func (o *Objx) CollectComplex128(collector func(int, complex128) interface{}) *Objx {

	arr := o.MustComplex128Slice()
	collected := make([]interface{}, len(arr))

	o.EachComplex128(func(index int, val complex128) bool {
		collected[index] = collector(index, val)
		return true
	})

	return New(collected)
}
