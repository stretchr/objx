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
// First you use the Get method to select the
//
// Example: Int(), Str(), Bool(), Float64(), UInt64() etc.
//
// If the data is missing or of the wrong type, the methods will
// return a default value, or optionally, you can specify the default value.
//
// Objx also providers stricter Must methods for each supported type (like MustStr(),
// MustBool(), MustInt(), MustFloat64() etc.)
package objx
