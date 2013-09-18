// objx - Go package for dealing with maps, slices, JSON and other data.
//
//     // use MustFromJSON to make an objx.Map from some JSON
//     m := objx.MustFromJSON(`{"name": "Mat", "age": 30}`)
//
//     // get the details
//     name := m.Get("name").Str()
//     age := m.Get("age").Int()
//
//     // get their nickname (or use their name if they
//     // don't have one)
//     nickname := m.Get("nickname").Str(name)
package objx
