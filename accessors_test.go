package objx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccessorsAccessGetSingleField(t *testing.T) {

	current := map[string]interface{}{"name": "Tyler"}
	assert.Equal(t, "Tyler", access(current, "name", true))

}
func TestAccessorsAccessGetDeep(t *testing.T) {

	current := map[string]interface{}{"name": map[string]interface{}{"first": "Tyler", "last": "Bunnell"}}
	assert.Equal(t, "Tyler", access(current, "name.first", true))
	assert.Equal(t, "Bunnell", access(current, "name.last", true))

}
func TestAccessorsAccessGetDeepDeep(t *testing.T) {

	current := map[string]interface{}{"one": map[string]interface{}{"two": map[string]interface{}{"three": map[string]interface{}{"four": 4}}}}
	assert.Equal(t, 4, access(current, "one.two.three.four", true))

}
func TestAccessorsAccessGetInsideArray(t *testing.T) {

	current := map[string]interface{}{"names": []interface{}{map[string]interface{}{"first": "Tyler", "last": "Bunnell"}, map[string]interface{}{"first": "Capitol", "last": "Bollocks"}}}
	assert.Equal(t, "Tyler", access(current, "names[0].first", true))
	assert.Equal(t, "Bunnell", access(current, "names[0].last", true))
	assert.Equal(t, "Capitol", access(current, "names[1].first", true))
	assert.Equal(t, "Bollocks", access(current, "names[1].last", true))

	assert.Panics(t, func() {
		access(current, "names[2]", true)
	})
	assert.Nil(t, access(current, "names[2]", false))

}

func TestAccessorsAccessGetFromArrayWithInt(t *testing.T) {

	current := []interface{}{map[string]interface{}{"first": "Tyler", "last": "Bunnell"}, map[string]interface{}{"first": "Capitol", "last": "Bollocks"}}
	one := access(current, 0, false)
	two := access(current, 1, false)
	three := access(current, 2, false)

	assert.Equal(t, "Tyler", one.(map[string]interface{})["first"])
	assert.Equal(t, "Capitol", two.(map[string]interface{})["first"])
	assert.Nil(t, three)

}

func TestAccessorsGet(t *testing.T) {

	current := New(map[string]interface{}{"name": "Tyler"})
	assert.Equal(t, "Tyler", current.Get("name").data)

}
