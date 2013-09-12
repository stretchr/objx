// Objx is a package for dealing with maps, slices, JSON and other data in Go.
//
// Reading data
//
// Objx follows a fluent design allowing you to quickly and intuitively access
// data from within otherwise difficult or deep nested places.
//
//     // parse some JSON
//     data := objx.MustFromJSON(`{
//       "name": "Tyler",
//       "age": 29,
//       "tags": ["musician","coder","singer"],
//       "deep": {
//         "nested": {
//           "item": "easy"
//         }
//       }
//     }`)
//
//     // get some fields
//     name := data.Get("name").Str()
//     age := data.Get("age").Int()
//
//     // access an array element
//     secondTag := data.Get("tags[2]").Str()
//
//     // get at some deep data using dot notation
//     deepItem := data.Get("deep.nested.item").Str()
//
// Type strong
//
// Objx is type strong, meaning that it attempts to remove the need to
// work so loosely with interface{}.  In that spirit, it providers
// strongly typed accessors that make reading data easy.
//
// First you use the Get method to select the value you're interested in,
// then use some of the strongly typed accessors:
//
// Example: Int(), Str(), Bool(), Float64(), UInt64() etc.
//
// If the data is missing or of the wrong type, the methods will
// return a default value, or optionally, you can specify the default value.
//
// Objx also providers stricter Must methods for each supported type (like MustStr(),
// MustBool(), MustInt(), MustFloat64() etc.)
//
// Slices
//
// Objx provides extra support to make working with arrays and slices easy, and in a
// fluent way.  Each method comes with strongly typed versions.
//
// Each* - Each methods iterate over a all items in a slice.  The callback func does the
// work, and takes the index and object like when ranging.
//
//     // log each string in a []string
//     obj.EachStr(func(index int, value string) bool {
//
//       // log each value
//       log.Println(value)
//
//       // return true (means carry on)
//       return true
//     })
//
// Where* - Where methods allow you to select a subset of items using a simple selector
// func.
//
//     // collect questions from a []string
//     obj.WhereStr(func(index int, value string) bool {
//
//       // does the string contain a question mark?
//       return strings.Contains(value, "?")
//
//     })
//
// Group* - Group methods (e.g. GroupStr, GroupInt etc.) allow you to generate a map[string]T
// where the key is the name of the group, and the values is a slice of the items in that group.
// The group name comes from the return of the func argument.
//
//     // group strings by the first character
//     listings := obj.GroupStr(func(index int, value string) string {
//
//       // return first character
//       return value[0]
//
//     })
//
// Replace* - Replace methods allow you to swap each value in a slice for another of the same
// type.  The type of the data will not change.
//
//     // increase all numbers in []int by one
//     increased := obj.ReplaceInt(func(index int, value int) int {
//       return value + 1
//     })
//
// Collect* - Collect methods allow you to collect new values (of type interface{}) by
// iterating over each item.  The type of the data returned will be []interface{}.
//
//     // Collect descriptions of the numbers in []float64
//     descriptions := obj.CollectFloat64(func(index int, value float64) interface{} {
//       return fmt.Sprintf("Value %d is %v", index, float64)
//     })
//
// Optimistic
//
// Objx is optimistic, which means it won't panic if things go wrong.  Instead, default values
// are used.  If you want panics to ensure the right data is in place, you can use the Must prefix
// on most methods.
package objx
