package objx_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/objx"
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
			"ifriends": []any{
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

func BenchmarkGetNestedBracketKeys(b *testing.B) {
	m := objx.Map{"root": objx.Map{"first": objx.Map{"second": objx.Map{"third": "leaf"}}}}
	b.ReportAllocs()
	for b.Loop() {
		m.Get("root[first][second][third]")
	}
}

func TestAccessorsAccessGetInsideArray(t *testing.T) {
	m := objx.Map{
		"names": []any{
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

func TestAccessorsAccessSetNilNestedMap(t *testing.T) {
	for name, nested := range map[string]any{
		"nil":     nil,
		"map":     map[string]any(nil),
		"objxMap": objx.Map(nil),
	} {
		t.Run(name, func(t *testing.T) {
			m := objx.Map{"nested": nested}
			m.Set("nested.answer", 42)
			assert.Equal(t, 42, m.Get("nested.answer").Data())
		})
	}
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
		"names": []any{"Tyler"},
	}
	m.Set("names[0]", "Mat")

	assert.Equal(t, "Mat", m.Get("names[0]").Data())
}

func TestAccessorsAccessSetInsideArray(t *testing.T) {
	m := objx.Map{
		"names": []any{
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

	m = objx.New(map[string]any{
		"a": map[string]any{
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

func TestGetIndexOnNonSlice(t *testing.T) {
	for _, value := range []any{42, "text", objx.Map{"name": "value"}, nil} {
		m := objx.Map{"value": value}
		assert.Nil(t, m.Get("value[0]").Data())
		assert.Nil(t, m.Get("value[0].name").Data())
	}
	m := objx.Map{"value": []any{42}}
	assert.Nil(t, m.Get("value[0][0]").Data())
	assert.Equal(t, 42, m.Get("value[0]").Data())
}

func TestSetIndexPreservesSlice(t *testing.T) {
	m := objx.Map{"names": []any{"Tyler", "Sam"}}
	m.Set("names[0]", "Mat")
	assert.Equal(t, []any{"Mat", "Sam"}, m["names"])
	m.Set("names[2]", "outside")
	assert.Equal(t, []any{"Mat", "Sam"}, m["names"])
	m.Set("names[1]", nil)
	assert.Equal(t, []any{"Mat", nil}, m["names"])

	m = objx.Map{"nested": []any{[]any{"first", "second"}}}
	m.Set("nested[0][1]", "changed")
	assert.Equal(t, []any{[]any{"first", "changed"}}, m["nested"])
}

func TestSetIndexTypedSlices(t *testing.T) {
	type tags []string
	type id int
	for _, tt := range []struct {
		name     string
		slice    any
		selector string
		value    any
		want     any
	}{
		{"strings", []string{"a", "b"}, "items[0]", "x", []string{"x", "b"}},
		{"ints", []int{1, 2}, "items[0]", 9, []int{9, 2}},
		{"any", []any{"a", "b"}, "items[0]", "x", []any{"x", "b"}},
		{"named slice", tags{"a", "b"}, "items[0]", "x", tags{"x", "b"}},
		{"named element", []id{1, 2}, "items[0]", id(9), []id{9, 2}},
		{"nested", [][]string{{"a", "b"}}, "items[0][1]", "x", [][]string{{"a", "x"}}},
		{"mixed", []any{[]string{"a", "b"}}, "items[0][1]", "x", []any{[]string{"a", "x"}}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := objx.Map{"items": tt.slice}
			m.Set(tt.selector, tt.value)
			assert.Equal(t, tt.want, m["items"])
			// Set must also update the caller's slice, including nested slices.
			assert.Equal(t, tt.want, tt.slice)
		})
	}
}

func TestSetIndexTypedSliceInMap(t *testing.T) {
	tags := []string{"a", "b"}
	m := objx.Map{"groups": []objx.Map{{"tags": tags}}}
	m.Set("groups[0].tags[1]", "x")
	assert.Equal(t, []objx.Map{{"tags": []string{"a", "x"}}}, m["groups"])
	assert.Equal(t, []string{"a", "x"}, tags)
}

func TestSetIndexAssignableValues(t *testing.T) {
	old, replacement := 1, 9
	err := errors.New("replacement")
	for _, tt := range []struct {
		name  string
		slice any
		value any
		want  any
	}{
		{"pointer", []*int{&old}, &replacement, []*int{&replacement}},
		{"nil pointer", []*int{&old}, nil, []*int{nil}},
		{"typed nil pointer", []*int{&old}, (*int)(nil), []*int{nil}},
		{"nil map", []map[string]any{{"a": 1}}, nil, []map[string]any{nil}},
		{"nil slice", [][]string{{"a"}}, nil, [][]string{nil}},
		{"nil channel", []chan int{make(chan int)}, nil, []chan int{nil}},
		{"nil function", []func(){func() {}}, nil, []func(){nil}},
		{"nil interface", []any{"a"}, nil, []any{nil}},
		{"nil typed interface", []error{err}, nil, []error{nil}},
		{"typed nil in interface", []error{err}, (*strconv.NumError)(nil), []error{(*strconv.NumError)(nil)}},
		{"interface", []error{nil}, err, []error{err}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := objx.Map{"items": tt.slice}
			m.Set("items[0]", tt.value)
			assert.Equal(t, tt.want, m["items"])
			assert.Equal(t, tt.want, tt.slice)
		})
	}
}

func TestSetIndexUnsupportedValues(t *testing.T) {
	type id int
	for _, tt := range []struct {
		name     string
		slice    any
		selector string
		value    any
		want     any
	}{
		{"wrong type", []int{1, 2}, "items[0]", "x", []int{1, 2}},
		{"nested wrong type", [][]int{{1, 2}}, "items[0][1]", "x", [][]int{{1, 2}}},
		{"no conversion", []int{1, 2}, "items[0]", int64(9), []int{1, 2}},
		{"named element", []id{1, 2}, "items[0]", 9, []id{1, 2}},
		{"nil scalar", []int{1, 2}, "items[0]", nil, []int{1, 2}},
		{"wrong typed nil", []int{1, 2}, "items[0]", (*int)(nil), []int{1, 2}},
		{"wrong interface", []error{nil}, "items[0]", "x", []error{nil}},
		{"out of range", []int{1, 2}, "items[2]", 9, []int{1, 2}},
		{"nested out of range", [][]int{{1, 2}}, "items[0][2]", 9, [][]int{{1, 2}}},
		{"nil slice", []int(nil), "items[0]", 9, []int(nil)},
		{"nil inner slice", [][]int{nil}, "items[0][0]", 9, [][]int{nil}},
		{"nil value", nil, "items[0]", 9, nil},
		{"non-slice", 1, "items[0]", 9, 1},
		{"nested non-slice", []any{1}, "items[0][0]", 9, []any{1}},
		{"array", [2]int{1, 2}, "items[0]", 9, [2]int{1, 2}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := objx.Map{"items": tt.slice}
			m.Set(tt.selector, tt.value)
			assert.Equal(t, tt.want, m["items"])
		})
	}
}

func TestSetOverflowingIndexPreservesSlice(t *testing.T) {
	const index = "999999999999999999999999"
	type tags []string
	for _, slice := range []any{[]any{"a", "b"}, []string{"a", "b"}, tags{"a", "b"}, []any(nil)} {
		for _, selector := range []string{
			"items[" + index + "]",
			"items[" + index + "].name",
			"items[" + index + "][0]",
			"items[+" + index + "]",
			"items[-" + index + "]",
		} {
			t.Run(fmt.Sprintf("%T/%s", slice, selector), func(t *testing.T) {
				m := objx.Map{"items": slice, "other": "kept"}
				m.Set(selector, "x")
				assert.Equal(t, objx.Map{"items": slice, "other": "kept"}, m)
			})
		}
	}

	slice := []any{"a", "b"}
	m := objx.Map{"groups": []any{objx.Map{"items": slice}}}
	m.Set("groups[0].items["+index+"]", "x")
	assert.Equal(t, objx.Map{"groups": []any{objx.Map{"items": []any{"a", "b"}}}}, m)

	for _, key := range []string{"name", index} {
		m := objx.Map{"keys": objx.Map{key: []any{"a", "b"}}}
		m.Set("keys["+key+"]["+index+"]", "x")
		assert.Equal(t, objx.Map{"keys": objx.Map{key: []any{"a", "b"}}}, m)
	}
}

func TestAccessorsBracketMapKeys(t *testing.T) {
	const index = "999999999999999999999999"
	for _, key := range []string{
		"name", "example.com", "123abc", index, "+" + index, "-" + index,
		index + "x", "+" + index + "x", "-" + index + "x", index + ".com",
	} {
		for _, useObjxMap := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/Map=%t", key, useObjxMap), func(t *testing.T) {
				keys := map[string]any{key: "old", "other": "kept"}
				m := objx.Map{"keys": keys}
				if useObjxMap {
					m["keys"] = objx.Map(keys)
				}
				selector := "keys[" + key + "]"
				assert.Equal(t, "old", m.Get(selector).Data())
				assert.True(t, m.Has(selector))
				m.Set(selector, "new")
				assert.Equal(t, map[string]any{key: "new", "other": "kept"}, keys)
				m.Delete(selector)
				assert.Equal(t, map[string]any{"other": "kept"}, keys)

				m = objx.Map{}
				m.Set(selector+"[value]", "created")
				assert.Equal(t, objx.Map{"keys": map[string]any{key: map[string]any{"value": "created"}}}, m)
			})
		}
	}
}

// Nonnumeric bracket keys retain the existing slice-to-map replacement behavior.
func TestSetNonNumericBracketKeyAfterSlice(t *testing.T) {
	const index = "999999999999999999999999"
	type tags []string
	for _, key := range []string{
		index + "x", "+" + index + "x", "-" + index + "x",
		index + ".com", index + "_tag", index + " ", index + "é",
	} {
		for _, slice := range []any{[]any{"a", "b"}, []string{"a", "b"}, tags{"a", "b"}, []any(nil)} {
			for _, selector := range []string{"items[" + key + "]", "[items][" + key + "]"} {
				t.Run(fmt.Sprintf("%T/%s", slice, selector), func(t *testing.T) {
					m := objx.Map{"items": slice, "other": "kept"}
					m.Set(selector, "x")
					assert.Equal(t, objx.Map{"items": map[string]any{key: "x"}, "other": "kept"}, m)
				})
			}
		}

		for _, parent := range []string{"name", index} {
			m := objx.Map{"keys": objx.Map{parent: []any{"a", "b"}}}
			m.Set("keys["+parent+"]["+key+"]", "x")
			assert.Equal(t, objx.Map{"keys": objx.Map{parent: map[string]any{key: "x"}}}, m)
		}
	}
}

func TestAccessorsDelete(t *testing.T) {
	m := objx.Map{
		"name": objx.Map{
			"first": "Tyler",
			"last":  "Bunnell",
		},
		"age": 30,
		"domains": objx.Map{
			"example.com": objx.Map{
				"apex": "example",
			},
		},
	}

	// Delete single top-level field
	m.Delete("age")
	assert.False(t, m.Has("age"))
	assert.Nil(t, m.Get("age").Data())

	// Delete nested field
	m.Delete("name.first")
	assert.False(t, m.Has("name.first"))
	assert.Equal(t, "Bunnell", m.Get("name.last").Data())

	// Delete with complex bracket key
	m.Delete("domains[example.com][apex]")
	assert.False(t, m.Has("domains[example.com][apex]"))
}

func TestDeleteCopiesEnclosingSlices(t *testing.T) {
	type tags []string
	type rows []tags
	type groups []rows

	orig := groups{{{"a", "b", "c"}, {"x"}}, {{"y"}}}
	m := objx.Map{"nested": orig}
	m.Delete("nested[0][0][1]")
	assert.Equal(t, groups{{{"a", "c"}, {"x"}}, {{"y"}}}, m["nested"])
	assert.Equal(t, groups{{{"a", "b", "c"}, {"x"}}, {{"y"}}}, orig)

	outer := []any{[]any{"a", "b", "c"}, "x"}
	m = objx.Map{"nested": outer}
	m.Delete("nested[0][1]")
	assert.Equal(t, []any{[]any{"a", "c"}, "x"}, m["nested"])
	assert.Equal(t, []any{[]any{"a", "b", "c"}, "x"}, outer)
}

func TestDeleteSliceElements(t *testing.T) {
	type tags []string
	type values []any

	for _, tt := range []struct {
		name     string
		input    any
		selector string
		want     any
	}{
		{"first", []int{1, 2, 3}, "items[0]", []int{2, 3}},
		{"middle", []int{1, 2, 3}, "items[1]", []int{1, 3}},
		{"last", []int{1, 2, 3}, "items[2]", []int{1, 2}},
		{"only", []string{"a"}, "items[0]", []string{}},
		{"nil element", []any{nil, "b"}, "items[0]", []any{"b"}},
		{"named strings", tags{"a", "b", "c"}, "items[1]", tags{"a", "c"}},
		{"named interfaces", values{"a", "b", "c"}, "items[1]", values{"a", "c"}},
		{"nested only", [][]string{{"a"}, {"b"}}, "items[0][0]", [][]string{{}, {"b"}}},
		{"mixed slices", []any{tags{"a", "b"}, "x"}, "items[0][0]", []any{tags{"b"}, "x"}},
		{"outer element", [][]string{{"a"}, {"b"}}, "items[0]", [][]string{{"b"}}},
		{"map elements", []objx.Map{{"a": 1}, {"b": 2}}, "items[0]", []objx.Map{{"b": 2}}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := objx.Map{"items": tt.input, "keep": true}
			m.Delete(tt.selector)
			assert.Equal(t, objx.Map{"items": tt.want, "keep": true}, m)
		})
	}
}

func TestDeletePreservesBackingArrayWithSpareCapacity(t *testing.T) {
	backing := []string{"a", "b", "c", "spare"}
	m := objx.Map{"tags": backing[:3]}
	m.Delete("tags[1]")
	assert.Equal(t, []string{"a", "c"}, m["tags"])
	assert.Equal(t, []string{"a", "b", "c", "spare"}, backing)
}

func TestDeleteMissingPaths(t *testing.T) {
	makeMap := func() objx.Map {
		return objx.Map{
			"scalar": 42,
			"tags":   []string{"a", "b"},
			"nested": [][]string{{"a", "b"}},
			"nil":    []string(nil),
			"empty":  []string{},
			"array":  [2]string{"a", "b"},
			"nilmap": map[string]any(nil),
		}
	}
	for _, selector := range []string{
		"missing", "missing.nested.key", "missing[0]", "scalar.key", "scalar[0]",
		"tags[99]", "tags[-1]", "tags[0].key", "tags[0][0]", "tags[0][99]",
		"nested[99][0]", "nested[0][99]", "nested[0][0][0]",
		"nil[0]", "empty[0]", "array[0]", "nilmap.key",
	} {
		t.Run(selector, func(t *testing.T) {
			m := makeMap()
			m.Delete(selector)
			assert.Equal(t, makeMap(), m)
		})
	}
}

func TestDeleteSelectorsInOrderAndChaining(t *testing.T) {
	m := objx.Map{"tags": []string{"a", "b", "c", "d"}}
	m.Delete("tags[0]", "tags[1]").Set("chained", true)
	assert.Equal(t, objx.Map{"tags": []string{"b", "d"}, "chained": true}, m)
	m.Delete("tags[0]", "tags[0]")
	assert.Equal(t, []string{}, m["tags"])
	m.Delete()
	assert.Equal(t, objx.Map{"tags": []string{}, "chained": true}, m)
	var nilMap objx.Map
	assert.Nil(t, nilMap.Delete())
	assert.Nil(t, nilMap.Delete("missing.key", "missing[0]"))
}

func TestDeleteWithinTypedMapSlices(t *testing.T) {
	m := objx.Map{
		"groups": []map[string]any{{
			"users": []objx.Map{{"name": nil, "tags": []string{"a", "b", "c"}}},
		}},
	}
	assert.True(t, m.Has("groups[0].users[0].name"))
	m.Delete("groups[0].users[0].name", "groups[0].users[0].tags[1]")
	assert.False(t, m.Has("groups[0].users[0].name"))
	assert.Equal(t, objx.Map{
		"groups": []map[string]any{{"users": []objx.Map{{"tags": []string{"a", "c"}}}}},
	}, m)
}
