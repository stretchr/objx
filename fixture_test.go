package objx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var getFixtures = []struct {
	// name is the name of the fixture (used for reporting
	// failures)
	name string
	// data is the JSON data to be worked on
	data string
	// set provides a map of arguments that are
	// passed to the Set method before the get test is performed.
	set map[string]interface{}
	// get is the argument(s) to pass to Get
	get interface{}
	// output is the expected output
	output interface{}
}{
	{
		name:   "Simple get",
		data:   `{"name": "Mat"}`,
		get:    "name",
		output: "Mat",
	},
	{
		name:   "Get with dot notation",
		data:   `{"address": {"city": "Boulder"}}`,
		get:    "address.city",
		output: "Boulder",
	},
	{
		name:   "Deep get with dot notation",
		data:   `{"one": {"two": {"three": {"four": "hello"}}}}`,
		get:    "one.two.three.four",
		output: "hello",
	},
	{
		name:   "Get missing with dot notation",
		data:   `{"one": {"two": {"three": {"four": "hello"}}}}`,
		get:    "one.ten",
		output: nil,
	},
	{
		name:   "Get with array notation",
		data:   `{"tags": ["one", "two", "three"]}`,
		get:    "tags[1]",
		output: "two",
	},
	{
		name:   "Get with array and dot notation",
		data:   `{"types": { "tags": ["one", "two", "three"]}}`,
		get:    "types.tags[1]",
		output: "two",
	},
	{
		name:   "Get with array and dot notation - field after array",
		data:   `{"tags": [{"name":"one"}, {"name":"two"}, {"name":"three"}]}`,
		get:    "tags[1].name",
		output: "two",
	},
	{
		name:   "Complex get with array and dot notation",
		data:   `{"tags": [{"list": [{"one":"pizza"}]}]}`,
		get:    "tags[0].list[0].one",
		output: "pizza",
	},
	{
		name:   "Get with integer argument",
		data:   `["one", "two", "three"]`,
		get:    1,
		output: "two",
	},
	{
		name:   "Get field from within string should be nil",
		data:   `{"name":"Tyler"}`,
		get:    "name.something",
		output: nil,
	},
	{
		name:   "Get field from within string (using array accessor) should be nil",
		data:   `{"numbers":["one", "two", "three"]}`,
		get:    "numbers[0].nope",
		output: nil,
	},
	{
		name:   "Set inside map",
		data:   `{"address": {"city": "Boulder"}}`,
		set:    map[string]interface{}{"address.city": "Denver"},
		get:    "address.city",
		output: "Denver",
	},
}

func TestFixtures(t *testing.T) {

	for _, fixture := range getFixtures {

		o := MustFromJSON(fixture.data)

		// do we need to do a set?
		if fixture.set != nil {
			for k, v := range fixture.set {
				o.Set(k, v)
			}
		}

		// get the value
		t.Logf("Running get fixture: \"%s\" (%v)", fixture.name, fixture)
		newO := o.Get(fixture.get)

		// make sure it matches
		assert.Equal(t, fixture.output, newO.Obj(),
			"Get fixture \"%s\" failed: %v", fixture.name, fixture,
		)

	}

}
