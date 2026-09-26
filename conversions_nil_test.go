package objx_test

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/objx"
)

func TestConversionJSONPreservesNilContainers(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		json  string
	}{
		{"nil interface slice", []interface{}(nil), "null"},
		{"nil map slice", []objx.Map(nil), "null"},
		{"nil string map slice", []map[string]interface{}(nil), "null"},
		{"nil map", objx.Map(nil), "null"},
		{"nil interface map", map[interface{}]interface{}(nil), "null"},
		{"nil string map", map[string]interface{}(nil), "null"},
		{"nil string slice", []string(nil), "null"},
		{"nil interface", nil, "null"},
		{"empty interface slice", []interface{}{}, "[]"},
		{"empty map slice", []objx.Map{}, "[]"},
		{"empty string map slice", []map[string]interface{}{}, "[]"},
		{"empty map", objx.Map{}, "{}"},
		{"empty interface map", map[interface{}]interface{}{}, "{}"},
		{"populated interface slice", []interface{}{1}, "[1]"},
		{"populated map slice", []objx.Map{{"key": "value"}}, `[{"key":"value"}]`},
		{"populated string map slice", []map[string]interface{}{{"key": "value"}}, `[{"key":"value"}]`},
		{"populated map", objx.Map{"key": "value"}, `{"key":"value"}`},
		{"populated interface map", map[interface{}]interface{}{1: "value"}, `{"1":"value"}`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, context := range []struct {
				name  string
				input objx.Map
				want  string
			}{
				{"direct", objx.Map{"value": test.value}, `{"value":` + test.json + `}`},
				{"in slice", objx.Map{"value": []interface{}{test.value}}, `{"value":[` + test.json + `]}`},
				{"in map", objx.Map{"value": objx.Map{"nested": test.value}}, `{"value":{"nested":` + test.json + `}}`},
			} {
				t.Run(context.name, func(t *testing.T) {
					got, err := context.input.JSON()
					require.NoError(t, err)
					assert.Equal(t, context.want, got)
					assert.Equal(t, context.want, context.input.MustJSON())
					encoded, err := context.input.Base64()
					require.NoError(t, err)
					decoded, err := base64.StdEncoding.DecodeString(encoded)
					require.NoError(t, err)
					assert.Equal(t, context.want, string(decoded))
				})
			}
		})
	}
}
