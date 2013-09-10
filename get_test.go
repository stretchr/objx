package objx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var getFixtures = []struct {
	data   string
	get    interface{}
	output interface{}
}{
	// get
	{
		data:   `{"name": "Mat"}`,
		get:    "name",
		output: "Mat",
	},
	// dot notation
	{
		data:   `{"address": {"city": "Boulder"}}`,
		get:    "address.city",
		output: "Boulder",
	},
	{
		data:   `{"one": {"two": {"three": {"four": "hello"}}}}`,
		get:    "one.two.three.four",
		output: "hello",
	},
	// no data
	{
		data:   `{"one": {"two": {"three": {"four": "hello"}}}}`,
		get:    "one.ten",
		output: nil,
	},
	// array [] notation
	{
		data:   `{"tags": ["one", "two", "three"]}`,
		get:    "tags[1]",
		output: "two",
	},
	// array number
	{
		data:   `["one", "two", "three"]`,
		get:    1,
		output: "two",
	},
}

func TestGetFixtures(t *testing.T) {

	for index, fixture := range getFixtures {

		o := MustJson(fixture.data)

		// call Get
		newO := o.Get(fixture.get)

		// ensure
		assert.Equal(t, fixture.output, newO.Obj, "Get fixture %d failed: %v", index+1, fixture)

	}

}
