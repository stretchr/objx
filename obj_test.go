package objx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNil(t *testing.T) {
	assert.NotNil(t, Nil)
	assert.Nil(t, Nil.Obj())
}
