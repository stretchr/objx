package objx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccessGetSingleField(t *testing.T) {

	current := map[string]interface{}{"name": "Tyler"}
	assert.Equal(t, "Tyler", access(current, "name", true))

}
func TestAccessGetDeep(t *testing.T) {

	current := map[string]interface{}{"name": map[string]interface{}{"first": "Tyler", "last": "Bunnell"}}
	assert.Equal(t, "Tyler", access(current, "name.first", true))
	assert.Equal(t, "Bunnell", access(current, "name.last", true))

}
func TestAccessGetDeepDeep(t *testing.T) {

	current := map[string]interface{}{"one": map[string]interface{}{"two": map[string]interface{}{"three": map[string]interface{}{"four": 4}}}}
	assert.Equal(t, 4, access(current, "one.two.three.four", true))

}
func TestAccessGetInsideArray(t *testing.T) {

	current := map[string]interface{}{"names": []map[string]interface{}{map[string]interface{}{"first": "Tyler", "last": "Bunnell"}, map[string]interface{}{"first": "Capitol", "last": "Bollocks"}}}
	assert.Equal(t, "Tyler", access(current, "names[0].first", true))
	assert.Equal(t, "Bunnell", access(current, "names[0].last", true))
	assert.Equal(t, "Capitol", access(current, "names[1].first", true))
	assert.Equal(t, "Bollocks", access(current, "names[1].last", true))

}
func TestAccessGetFromArrayWithInt(t *testing.T) {

	current := []map[string]interface{}{map[string]interface{}{"first": "Tyler", "last": "Bunnell"}, map[string]interface{}{"first": "Capitol", "last": "Bollocks"}}
	one := access(current, 0, false)
	two := access(current, 1, false)
	three := access(current, 2, false)

	assert.Equal(t, "Tyler", one.(map[string]interface{})["first"])
	assert.Equal(t, "Capitol", two.(map[string]interface{})["first"])
	assert.Nil(t, three)

}
