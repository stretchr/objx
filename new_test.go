package objx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {

	o := New(TestMap)

	if assert.NotNil(t, o) {
		assert.Equal(t, TestMap, o.Obj)
	}

}
