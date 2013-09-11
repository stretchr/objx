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
