package objx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet(t *testing.T) {

	m := map[string]interface{}{"name": "Mat"}
	o := New(m)
	assert.Equal(t, o, o.Set("name", "Tyler"), "Set should chain")
	assert.Equal(t, "Tyler", m["name"])

}

func TestSetInDeepMap(t *testing.T) {

	m := map[string]interface{}{"one": map[string]interface{}{"two": map[string]interface{}{"three": 3}}}
	o := New(m)
	assert.Equal(t, o, o.Set("one.two.three", "THREE"), "Set should chain")
	assert.Equal(t, "THREE", m["one"].(map[string]interface{})["two"].(map[string]interface{})["three"])

}

func TestSetNilInDeepMap(t *testing.T) {

	m := map[string]interface{}{"one": map[string]interface{}{"two": map[string]interface{}{"three": 3}}}
	o := New(m)
	assert.Equal(t, o, o.Set("one.two.three", nil), "Set should chain")
	assert.Equal(t, nil, m["one"].(map[string]interface{})["two"].(map[string]interface{})["three"])

}

func TestSetInArray(t *testing.T) {

	a := []string{"one", "two", "three"}
	o := New(a)

	o.Set(0, "1").Set(1, "2").Set(2, "3")

	assert.Equal(t, "1", a[0])
	assert.Equal(t, "2", a[1])
	assert.Equal(t, "3", a[2])

}
