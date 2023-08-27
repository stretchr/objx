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
			"friends": []string{
				"Capitol",
				"Bollocks",
			},
			"ifriends": []interface{}{
				"Capitol",
				"Bollocks",
			},
		},
	}

	assert.Equal(t, "Tyler", m.Get("name.first").Data())
	assert.Equal(t, "Bunnell", m.Get("name.last").Data())
	assert.Equal(t, "Capitol", m.Get("name.friends[0]").Data())
	assert.Equal(t, "Capitol", m.Get("name.ifriends[0]").Data())
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
	assert.Equal(t, 4, m.Get("one[two][three][four]").Data())
}

func TestAccessorsGetWithComplexKey(t *testing.T) {
	m := objx.Map{
		"domains": objx.Map{
			"example-dot-com": objx.Map{
				"apex": "example",
			},
			"example.com": objx.Map{
				"apex": "example",
			},
		},
	}

	assert.Equal(t, "example", m.Get("domains.example-dot-com.apex").Data())

	assert.Equal(t, "example", m.Get("domains[example.com].apex").Data())
	assert.Equal(t, "example", m.Get("domains[example.com][apex]").Data())
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
	assert.Nil(t, m.Get("names[]").Data())
	assert.Nil(t, m.Get("names1]]").Data())
	assert.Nil(t, m.Get("names[1]]").Data())
	assert.Nil(t, m.Get("names[[1]]").Data())
	assert.Nil(t, m.Get("names[[1]").Data())
	assert.Nil(t, m.Get("names[[1").Data())
}

func TestAccessorsGet(t *testing.T) {
	m := objx.Map{"name": "Tyler"}

	assert.Equal(t, "Tyler", m.Get("name").Data())
}

func TestAccessorsAccessSetSingleField(t *testing.T) {
	m := objx.Map{"name": "Tyler"}

	m.Set("name", "Mat")
	m.Set("age", 29)

	assert.Equal(t, m.Get("name").Data(), "Mat")
	assert.Equal(t, m.Get("age").Data(), 29)
}

func TestAccessorsAccessSetSingleFieldNotExisting(t *testing.T) {
	m := objx.Map{
		"first": "Tyler",
		"last":  "Bunnell",
	}

	m.Set("name", "Mat")

	assert.Equal(t, m.Get("name").Data(), "Mat")
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

func TestAccessorsAccessSetDeepDeepWithoutExisting(t *testing.T) {
	m := objx.Map{}

	m.Set("one.two.three.four", 5)
	m.Set("one.two.three.five", 6)

	assert.Equal(t, 5, m.Get("one.two.three.four").Data())
	assert.Equal(t, 6, m.Get("one.two.three.five").Data())

	m.Set("one.two", 7)
	assert.Equal(t, 7, m.Get("one.two").Data())
	assert.Equal(t, nil, m.Get("one.two.three.four").Data())

	m.Set("one.two.three", 8)
	assert.Equal(t, 8, m.Get("one.two.three").Data())
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

func TestAccessorsSetWithinObjxMapChild(t *testing.T) {
	m, err := objx.FromJSON(`{"a": {"b": 1}}`)
	assert.NoError(t, err)

	m.Set("a.c", 2)
	jsonConverted, err := m.JSON()
	assert.NoError(t, err)

	m = objx.New(map[string]interface{}{
		"a": map[string]interface{}{
			"b": 1,
		},
	})
	m.Set("a.c", 2)
	jsonNewObj, err := m.JSON()

	assert.NoError(t, err)
	assert.Equal(t, jsonConverted, jsonNewObj)
}

func TestAccessorsNested(t *testing.T) {
	d := objx.MustFromJSON(`{"values":[["test", "test1"], ["test2", {"name":"Mat"}, {"names": ["Captain", "Mat"]}]]}`)

	value := d.Get("values[0][0]").String()
	assert.Equal(t, "test", value)

	value = d.Get("values[0][1]").String()
	assert.Equal(t, "test1", value)

	value = d.Get("values[1][0]").String()
	assert.Equal(t, "test2", value)

	value = d.Get("values[1][1].name").String()
	assert.Equal(t, "Mat", value)

	value = d.Get("values[1][2].names[0]").String()
	assert.Equal(t, "Captain", value)
}

func TestGenericDeepSetNested(t *testing.T) {
	m := objx.Map{}

	m.Set("other[0].x", "e0")
	m.Set("other[1].x", "e1")

	json, err := m.JSON()
	assert.NoError(t, err)
	assert.Equal(t, `{"other":[{"x":"e0"},{"x":"e1"}]}`, json)
}
func TestGenericDeepSetMixed(t *testing.T) {
	m := objx.Map{}

	m.Set("other[0].x", "e0")
	m.Set("other[1]", "e1")

	json, err := m.JSON()
	assert.NoError(t, err)
	assert.Equal(t, `{"other":[{"x":"e0"},"e1"]}`, json)
}
func TestGenericDeepSetOverride(t *testing.T) {
	m := objx.Map{}

	m.Set("other[0].x", "e0")
	m.Set("other[1].x", "e1")
	m.Set("other", "str")

	json, err := m.JSON()
	assert.NoError(t, err)
	assert.Equal(t, `{"other":"str"}`, json)
}
func TestGenericDeepSetArraySkip(t *testing.T) {
	m := objx.Map{}

	m.Set("other[1]", "e1")
	m.Set("other[3]", "e3")

	json, err := m.JSON()
	assert.NoError(t, err)
	assert.Equal(t, `{"other":[null,"e1",null,"e3"]}`, json)
}
func TestGenericDeepSetArrayAlloc(t *testing.T) {
	m := objx.Map{}

	m.Set("other[10]", "e10")
	m.Set("other[1]", "e1")

	json, err := m.JSON()
	assert.NoError(t, err)
	assert.Equal(t, `{"other":[null,"e1",null,null,null,null,null,null,null,null,"e10"]}`, json)
}

func TestGenericDeepGetNested(t *testing.T) {
	m := objx.Map{}

	m.Set("other[0].x", "e0")
	m.Set("other[1].x", "e1")

	assert.Equal(t, "e0", m.Get("other[0].x").Data())
	assert.Equal(t, "e1", m.Get("other[1].x").Data())
	assert.Equal(t, nil, m.Get("other[2].x").Data())
	assert.Equal(t, nil, m.Get("nope").Data())
	assert.Equal(t, nil, m.Get("other[0].nope").Data())
	assert.Equal(t, nil, m.Get("other[-1].nope").Data())

	json, err := m.JSON()
	assert.NoError(t, err)
	assert.Equal(t, `{"other":[{"x":"e0"},{"x":"e1"}]}`, json)
}

func TestGenericDeepGetSkip(t *testing.T) {
	m := objx.Map{}

	m.Set("other[1].x", "e1")

	assert.Equal(t, nil, m.Get("other[0]").Data())
	assert.Equal(t, nil, m.Get("other[0].x").Data())
	assert.Equal(t, "e1", m.Get("other[1].x").Data())

	json, err := m.JSON()
	assert.NoError(t, err)
	assert.Equal(t, `{"other":[null,{"x":"e1"}]}`, json)
}

func TestGenericDeepPrimitives(t *testing.T) {
	m := objx.Map{}

	m.Set("str", "str")
	m.Set("int", 1)
	m.Set("float", 1.1)
	m.Set("boolean", true)
	m.Set("obj", map[string]interface{}{})
	m.Set("arr", []interface{}{})
	m.Set("null", nil)

	assert.Equal(t, "str", m.Get("str").Data())
	assert.Equal(t, 1, m.Get("int").Data())
	assert.Equal(t, 1.1, m.Get("float").Data())
	assert.Equal(t, true, m.Get("boolean").Data())
	assert.Equal(t, map[string]interface{}{}, m.Get("obj").Data())
	assert.Equal(t, []interface{}{}, m.Get("arr").Data())
	assert.Equal(t, nil, m.Get("null").Data())

	json, err := m.JSON()
	assert.NoError(t, err)
	assert.Equal(t, `{"arr":[],"boolean":true,"float":1.1,"int":1,"null":null,"obj":{},"str":"str"}`, json)
}

func TestGenericDeepNull(t *testing.T) {
	m := objx.MustFromJSON(`{"nope":null}`)

	m.Set("null", nil)

	assert.Equal(t, true, m.Has("nope"))
	assert.Equal(t, true, m.Has("null"))
	assert.Equal(t, false, m.Has("no-exist"))

	assert.Equal(t, nil, m.Get("nope").Data())
	assert.Equal(t, nil, m.Get("null").Data())
	assert.Equal(t, nil, m.Get("no-exist").Data())

	json, err := m.JSON()
	assert.NoError(t, err)
	assert.Equal(t, `{"nope":null,"null":null}`, json)
}

