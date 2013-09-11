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
//	       "nested": {
//	         item: "easy"
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
package objx
