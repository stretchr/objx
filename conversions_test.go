package objx_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/objx"
)

func TestConversionJSON(t *testing.T) {
	jsonString := `{"name":"Mat"}`
	o := objx.MustFromJSON(jsonString)

	result, err := o.JSON()

	require.NoError(t, err)
	assert.Equal(t, jsonString, result)
	assert.Equal(t, jsonString, o.MustJSON())

	i := objx.Map{
		"a": map[any]any{
			"b": objx.Map{"c": map[any]any{"d": "e"}},
			"f": []objx.Map{{"g": map[any]any{"h": "i"}}},
			"j": []map[string]any{{"k": map[any]any{"l": "m"}}},
			"n": []any{objx.Map{"o": "p"}},
		},
	}

	jsonString = `{"a":{"b":{"c":{"d":"e"}},"f":[{"g":{"h":"i"}}],"j":[{"k":{"l":"m"}}],"n":[{"o":"p"}]}}`
	result, err = i.JSON()
	require.NoError(t, err)
	assert.Equal(t, jsonString, result)
	assert.Equal(t, jsonString, i.MustJSON())
}

func TestConversionJSONWithError(t *testing.T) {
	o := objx.MSI()
	o["test"] = func() {}

	assert.Panics(t, func() {
		o.MustJSON()
	})

	_, err := o.JSON()

	assert.Error(t, err)
}

func TestConversionJSONNestedStringMaps(t *testing.T) {
	m := objx.Map{
		"config": map[string]any{"nested": map[any]any{true: 1}},
		"items":  []any{map[string]any{"nested": map[any]any{false: 2}}},
	}
	result, err := m.JSON()
	require.NoError(t, err)
	assert.Equal(t, `{"config":{"nested":{"true":1}},"items":[{"nested":{"false":2}}]}`, result)
}

func TestConversionJSONDoesNotMutateReceiver(t *testing.T) {
	for name, encode := range map[string]func(objx.Map) (string, error){
		"JSON":   objx.Map.JSON,
		"Base64": objx.Map.Base64,
		"MustJSON": func(m objx.Map) (string, error) {
			return m.MustJSON(), nil
		},
	} {
		t.Run(name, func(t *testing.T) {
			m := objx.Map{
				"config": map[string]any{"nested": map[any]any{true: 1}},
				"items":  []any{map[string]any{"nested": map[any]any{false: 2}}},
			}

			_, err := encode(m)
			require.NoError(t, err)
			assert.Equal(t, objx.Map{
				"config": map[string]any{"nested": map[any]any{true: 1}},
				"items":  []any{map[string]any{"nested": map[any]any{false: 2}}},
			}, m)
		})
	}
}

func TestConversionBase64(t *testing.T) {
	o := objx.Map{"name": "Mat"}

	result, err := o.Base64()

	require.NoError(t, err)
	assert.Equal(t, "eyJuYW1lIjoiTWF0In0=", result)
	assert.Equal(t, "eyJuYW1lIjoiTWF0In0=", o.MustBase64())
}

func TestConversionBase64WithError(t *testing.T) {
	o := objx.MSI()
	o["test"] = func() {}

	assert.Panics(t, func() {
		o.MustBase64()
	})

	_, err := o.Base64()

	assert.Error(t, err)
}

func TestConversionSignedBase64(t *testing.T) {
	o := objx.Map{"name": "Mat"}

	result, err := o.SignedBase64("key")

	require.NoError(t, err)
	assert.Equal(t, "eyJuYW1lIjoiTWF0In0=_67ee82916f90b2c0d68c903266e8998c9ef0c3d6", result)
	assert.Equal(t, "eyJuYW1lIjoiTWF0In0=_67ee82916f90b2c0d68c903266e8998c9ef0c3d6", o.MustSignedBase64("key"))
}

func TestConversionSignedBase64WithError(t *testing.T) {
	o := objx.MSI()
	o["test"] = func() {}

	assert.Panics(t, func() {
		o.MustSignedBase64("key")
	})

	_, err := o.SignedBase64("key")

	assert.Error(t, err)
}

func TestConversionURLValues(t *testing.T) {
	m := getURLQueryMap()
	u := m.URLValues()

	assert.Equal(t, url.Values{
		"abc":                []string{"123"},
		"name":               []string{"Mat"},
		"data[age]":          []string{"30"},
		"data[height]":       []string{"162"},
		"data[arr][]":        []string{"1", "2"},
		"stats[]":            []string{"1", "2"},
		"bools[]":            []string{"true", "false"},
		"mapSlice[][age]":    []string{"40"},
		"mapSlice[][height]": []string{"152"},
		"msiData[age]":       []string{"30"},
		"msiData[height]":    []string{"162"},
		"msiData[arr][]":     []string{"1", "2"},
		"msiSlice[][age]":    []string{"40"},
		"msiSlice[][height]": []string{"152"},
	}, u)
}

func TestConversionURLQuery(t *testing.T) {
	m := getURLQueryMap()
	u, err := m.URLQuery()

	assert.Nil(t, err)
	require.NotNil(t, u)

	ue, err := url.QueryUnescape(u)
	assert.Nil(t, err)
	require.NotNil(t, ue)

	assert.Equal(t, "abc=123&bools[]=true&bools[]=false&data[age]=30&data[arr][]=1&data[arr][]=2&data[height]=162&mapSlice[][age]=40&mapSlice[][height]=152&msiData[age]=30&msiData[arr][]=1&msiData[arr][]=2&msiData[height]=162&msiSlice[][age]=40&msiSlice[][height]=152&name=Mat&stats[]=1&stats[]=2", ue)
}

func TestConversionURLQueryNoSliceKeySuffix(t *testing.T) {
	t.Cleanup(func() { _ = objx.SetURLValuesSliceKeySuffix(objx.URLValuesSliceKeySuffixArray) })
	m := getURLQueryMap()
	err := objx.SetURLValuesSliceKeySuffix(objx.URLValuesSliceKeySuffixEmpty)
	require.Nil(t, err)
	u, err := m.URLQuery()

	assert.Nil(t, err)
	require.NotNil(t, u)

	ue, err := url.QueryUnescape(u)
	assert.Nil(t, err)
	require.NotNil(t, ue)

	assert.Equal(t, "abc=123&bools=true&bools=false&data[age]=30&data[arr]=1&data[arr]=2&data[height]=162&mapSlice[age]=40&mapSlice[height]=152&msiData[age]=30&msiData[arr]=1&msiData[arr]=2&msiData[height]=162&msiSlice[age]=40&msiSlice[height]=152&name=Mat&stats=1&stats=2", ue)
}

func TestConversionURLQueryIndexSliceKeySuffix(t *testing.T) {
	t.Cleanup(func() { _ = objx.SetURLValuesSliceKeySuffix(objx.URLValuesSliceKeySuffixArray) })
	m := getURLQueryMap()
	m.Set("mapSlice", []objx.Map{{"age": 40, "sex": "male"}, {"height": 152}})
	err := objx.SetURLValuesSliceKeySuffix(objx.URLValuesSliceKeySuffixIndex)
	require.Nil(t, err)
	u, err := m.URLQuery()

	assert.Nil(t, err)
	require.NotNil(t, u)

	ue, err := url.QueryUnescape(u)
	assert.Nil(t, err)
	require.NotNil(t, ue)

	assert.Equal(t, "abc=123&bools[0]=true&bools[1]=false&data[age]=30&data[arr][0]=1&data[arr][1]=2&data[height]=162&mapSlice[0][age]=40&mapSlice[0][sex]=male&mapSlice[1][height]=152&msiData[age]=30&msiData[arr][0]=1&msiData[arr][1]=2&msiData[height]=162&msiSlice[0][age]=40&msiSlice[1][height]=152&name=Mat&stats[0]=1&stats[1]=2", ue)
}

func TestConversionURLValuesNested(t *testing.T) {
	m := objx.Map{
		"a": objx.Map{
			"b": map[string]any{"c": 1},
			"s": []map[string]any{{"x": 1, "arr": []int{1, 2}}, {"x": 2, "arr": []int{3}}},
			"o": []any{objx.Map{"y": 1}, map[string]any{"y": 2}},
		},
		"top": []objx.Map{{"x": 1}, {"x": 2}},
		"j":   []any{1, "two", nil, map[string]any{"z": 3}, []any{4}},
	}

	// Without an index suffix, values from different slice elements are
	// merged under shared keys (e.g. a[s][][arr][] below); only the "[i]"
	// suffix keeps them separable.

	tests := []struct {
		suffix   string
		expected url.Values
	}{
		{
			suffix: objx.URLValuesSliceKeySuffixArray,
			expected: url.Values{
				"a[b][c]":       []string{"1"},
				"a[s][][x]":     []string{"1", "2"},
				"a[s][][arr][]": []string{"1", "2", "3"},
				"a[o][][y]":     []string{"1", "2"},
				"top[][x]":      []string{"1", "2"},
				"j[]":           []string{"1", "two"},
				"j[][z]":        []string{"3"},
				"j[][]":         []string{"4"},
			},
		},
		{
			suffix: objx.URLValuesSliceKeySuffixEmpty,
			expected: url.Values{
				"a[b][c]":   []string{"1"},
				"a[s][x]":   []string{"1", "2"},
				"a[s][arr]": []string{"1", "2", "3"},
				"a[o][y]":   []string{"1", "2"},
				"top[x]":    []string{"1", "2"},
				"j":         []string{"1", "two", "4"},
				"j[z]":      []string{"3"},
			},
		},
		{
			suffix: objx.URLValuesSliceKeySuffixIndex,
			expected: url.Values{
				"a[b][c]":         []string{"1"},
				"a[s][0][x]":      []string{"1"},
				"a[s][0][arr][0]": []string{"1"},
				"a[s][0][arr][1]": []string{"2"},
				"a[s][1][x]":      []string{"2"},
				"a[s][1][arr][0]": []string{"3"},
				"a[o][0][y]":      []string{"1"},
				"a[o][1][y]":      []string{"2"},
				"top[0][x]":       []string{"1"},
				"top[1][x]":       []string{"2"},
				"j[0]":            []string{"1"},
				"j[1]":            []string{"two"},
				"j[3][z]":         []string{"3"},
				"j[4][0]":         []string{"4"},
			},
		},
	}

	t.Cleanup(func() { _ = objx.SetURLValuesSliceKeySuffix(objx.URLValuesSliceKeySuffixArray) })
	for _, tt := range tests {
		t.Run(tt.suffix, func(t *testing.T) {
			require.Nil(t, objx.SetURLValuesSliceKeySuffix(tt.suffix))
			assert.Equal(t, tt.expected, m.URLValues())
		})
	}
}

func TestValidityURLQuerySliceKeySuffix(t *testing.T) {
	t.Cleanup(func() { _ = objx.SetURLValuesSliceKeySuffix(objx.URLValuesSliceKeySuffixArray) })
	err := objx.SetURLValuesSliceKeySuffix("")
	assert.Nil(t, err)
	err = objx.SetURLValuesSliceKeySuffix("[]")
	assert.Nil(t, err)
	err = objx.SetURLValuesSliceKeySuffix("[i]")
	assert.Nil(t, err)
	err = objx.SetURLValuesSliceKeySuffix("{}")
	assert.Error(t, err)
}

func getURLQueryMap() objx.Map {
	return objx.Map{
		"abc":      123,
		"name":     "Mat",
		"data":     objx.Map{"age": 30, "height": 162, "arr": []int{1, 2}},
		"mapSlice": []objx.Map{{"age": 40}, {"height": 152}},
		"msiData":  map[string]any{"age": 30, "height": 162, "arr": []int{1, 2}},
		"msiSlice": []map[string]any{{"age": 40}, {"height": 152}},
		"stats":    []string{"1", "2"},
		"bools":    []bool{true, false},
	}
}
