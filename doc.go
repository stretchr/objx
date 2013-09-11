// Objx is a package for dealing with maps, slices, JSON and other data in Go.
//
// Reading data
//
//     obj, err := objx.FromJSON(`{"name":"Tyler","age":29,"tags":["musician","coder","singer"]}`)
//     name := obj.Get("name").Str()
//     age := obj.Get("age").Int()
//     secondTag := obj.Get("tags[2]").Str()
package objx
