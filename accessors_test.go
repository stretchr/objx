package objx_test

import (
	"testing"

	"github.com/stretchr/objx"
	"github.com/stretchr/testify/assert"
)

func TestAccessorsAccessGetSingleField(t *testing.T) {
	m := objx.Map{"name": "Tyler"}

	assert.Equal(t, "Tyler", m.Get("name").Data())
}

func TestAccessorsAccessGetSingleFieldInt(t *testing.T) {
	m := objx.Map{"name": 10}

	assert.Equal(t, 10, m.Get("name").Data())
}

func TestAccessorsAccessGetDeep(t *testing.T) {
	m := objx.Map{
		"name": objx.Map{
			"first": "Tyler",
			"last":  "Bunnell",
		},
	}

	assert.Equal(t, "Tyler", m.Get("name.first").Data())
	assert.Equal(t, "Bunnell", m.Get("name.last").Data())
}

func TestAccessorsAccessGetDeepDeep(t *testing.T) {
	m := objx.Map{
		"one": objx.Map{
			"two": objx.Map{
				"three": objx.Map{
					"four": 4,
				},
			},
		},
	}

	assert.Equal(t, 4, m.Get("one.two.three.four").Data())
}

func TestAccessorsAccessGetInsideArray(t *testing.T) {
	m := objx.Map{
		"names": []interface{}{
			objx.Map{
				"first": "Tyler",
				"last":  "Bunnell",
			},
			objx.Map{
				"first": "Capitol",
				"last":  "Bollocks",
			},
		},
	}

	assert.Equal(t, "Tyler", m.Get("names[0].first").Data())
	assert.Equal(t, "Bunnell", m.Get("names[0].last").Data())
	assert.Equal(t, "Capitol", m.Get("names[1].first").Data())
	assert.Equal(t, "Bollocks", m.Get("names[1].last").Data())

	assert.Nil(t, m.Get("names[2]").Data())
}

func TestAccessorsGet(t *testing.T) {
	m := objx.Map{"name": "Tyler"}

	assert.Equal(t, "Tyler", m.Get("name").Data())
}

func TestAccessorsAccessSetSingleField(t *testing.T) {
	m := objx.Map{"name": "Tyler"}

	m.Set("name", "Mat")
	m.Set("age", 29)

	assert.Equal(t, m["name"], "Mat")
	assert.Equal(t, m["age"], 29)
}

func TestAccessorsAccessSetSingleFieldNotExisting(t *testing.T) {
	m := objx.Map{
		"first": "Tyler",
		"last":  "Bunnell",
	}

	m.Set("name", "Mat")

	assert.Equal(t, m["name"], "Mat")
}

func TestAccessorsAccessSetDeep(t *testing.T) {
	m := objx.Map{
		"name": objx.Map{
			"first": "Tyler",
			"last":  "Bunnell",
		},
	}

	m.Set("name.first", "Mat")
	m.Set("name.last", "Ryer")

	assert.Equal(t, "Mat", m.Get("name.first").Data())
	assert.Equal(t, "Ryer", m.Get("name.last").Data())
}

func TestAccessorsAccessSetDeepDeep(t *testing.T) {
	m := objx.Map{
		"one": objx.Map{
			"two": objx.Map{
				"three": objx.Map{
					"four": 4,
				},
			},
		},
	}

	m.Set("one.two.three.four", 5)

	assert.Equal(t, 5, m.Get("one.two.three.four").Data())
}

func TestAccessorsAccessSetArray(t *testing.T) {
	m := objx.Map{
		"names": []interface{}{"Tyler"},
	}
	m.Set("names[0]", "Mat")

	assert.Equal(t, "Mat", m.Get("names[0]").Data())
}

func TestAccessorsAccessSetInsideArray(t *testing.T) {
	m := objx.Map{
		"names": []interface{}{
			objx.Map{
				"first": "Tyler",
				"last":  "Bunnell",
			},
			objx.Map{
				"first": "Capitol",
				"last":  "Bollocks",
			},
		},
	}

	m.Set("names[0].first", "Mat")
	m.Set("names[0].last", "Ryer")
	m.Set("names[1].first", "Captain")
	m.Set("names[1].last", "Underpants")

	assert.Equal(t, "Mat", m.Get("names[0].first").Data())
	assert.Equal(t, "Ryer", m.Get("names[0].last").Data())
	assert.Equal(t, "Captain", m.Get("names[1].first").Data())
	assert.Equal(t, "Underpants", m.Get("names[1].last").Data())
}

func TestAccessorsSet(t *testing.T) {
	m := objx.Map{"name": "Tyler"}

	m.Set("name", "Mat")

	assert.Equal(t, "Mat", m.Get("name").Data())
}
