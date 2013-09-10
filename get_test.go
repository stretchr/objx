package objx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var getFixtures = []struct {
	name   string
	data   string
	get    interface{}
	output interface{}
}{
	// get
	{
		name:   "Simple get",
		data:   `{"name": "Mat"}`,
		get:    "name",
		output: "Mat",
	},
	// dot notation
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
	// no data
	{
		name:   "Get missing with dot notation",
		data:   `{"one": {"two": {"three": {"four": "hello"}}}}`,
		get:    "one.ten",
		output: nil,
	},
	// array [] notation
	{
		name:   "Get with array notation",
		data:   `{"tags": ["one", "two", "three"]}`,
		get:    "tags[1]",
		output: "two",
	},
	// array [] notation in map
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
	// array number
	{
		name:   "Get with integer argument",
		data:   `["one", "two", "three"]`,
		get:    1,
		output: "two",
	},
}

func TestGetFixtures(t *testing.T) {

	for index, fixture := range getFixtures {
		o := MustJson(fixture.data)
		newO := o.Get(fixture.get)
		assert.Equal(t, fixture.output, newO.Obj, "Get fixture %d (%s) failed: %v", index+1, fixture.name, fixture)
	}

}
