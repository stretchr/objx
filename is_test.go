package objx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsNil(t *testing.T) {

	n := New(nil)

	assert.True(t, n.IsNil())

	n.obj = "something"

	assert.False(t, n.IsNil())

}
