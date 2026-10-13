package objx_test

import (
	"testing"

	"github.com/stretchr/objx"
)

func TestDeleteNestedIndexesPreservesOuterSlice(t *testing.T) {
	m := objx.Map{"nested": []any{[]any{"a", "b", "c"}, "x"}}
	m.Delete("nested[0][1]")
	assert.Equal(t, []any{[]any{"a", "c"}, "x"}, m["nested"])
}

func TestDeletePreservesCallerSlice(t *testing.T) {
	orig := []any{"a", "b", "c"}
	m := objx.Map{"tags": orig}
	m.Delete("tags[0]")
	assert.Equal(t, []any{"b", "c"}, m["tags"])
	assert.Equal(t, []any{"a", "b", "c"}, orig)
}

func TestDeletePreservesTypedSlice(t *testing.T) {
	m := objx.Map{"tags": []string{"a", "b", "c"}}
	m.Delete("tags[0]")
	assert.Equal(t, []string{"b", "c"}, m["tags"])
	assert.Equal(t, []string{"b", "c"}, m.Get("tags").StrSlice())
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

	result := m.Get("tags").StrSlice()
	result[0] = "changed"
	result = append(result, "new")
	assert.Equal(t, []string{"changed", "c", "new"}, result)
	assert.Equal(t, []string{"a", "b", "c", "spare"}, backing)
}

func TestDeleteMissingPaths(t *testing.T) {
	for _, selector := range []string{
		"missing", "missing.nested.key", "missing[0]", "scalar.key", "scalar[0]",
		"tags[99]", "tags[-1]", "tags[0].key", "tags[0][0]", "tags[0][99]",
		"nested[99][0]", "nested[0][99]", "nested[0][0][0]",
		"nil[0]", "empty[0]", "array[0]", "nilmap.key",
	} {
		t.Run(selector, func(t *testing.T) {
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
