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

func TestSetInArray(t *testing.T) {

	a := []string{"one", "two", "three"}
	o := New(a)

	o.Set(0, "1").Set(1, "2").Set(2, "3")

	assert.Equal(t, "1", a[0])
	assert.Equal(t, "2", a[1])
	assert.Equal(t, "3", a[2])

}
