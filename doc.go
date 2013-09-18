// objx - Go package for dealing with maps, slices, JSON and other data.
//
// Pattern
//
// Objx uses a preditable pattern to make access data from within `map[string]interface{}'s
// easy.
//
//     1. Use `Get` to access the value you're interested in.  You can use dot and array
//        notation too, like this: `Get("places[0].latlng")`
//
//     2. Once you have saught the `Value` you're interested in, you can use the `Is*` methods
//        to determine its type.  Like this: `if Get("code").IsStr()`
//
//     3. Or you can just assume the type, and use one of the strong type methods to
//        extract the real value.  Like this: `Get("code").Int()`
//
//     4. If there's no value there (or if it's the wrong type) then a default value
//        will be returned, or you can be explicit about the default value.
//        Like this: `Get("code").Int(-1)`
//
// Example
//
// A simple example of how to use Objx:
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
