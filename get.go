package objx

import (
	"strings"
)

// Get gets a value from this object.
func (o *O) Get(selector interface{}) *O {

	switch selector.(type) {
	case string: // "address.postcode.inner"

		selStr := selector.(string)
		segs := strings.Split(selStr, ".")

		current := o.Obj

		for _, field := range segs {

			if m, ok := current.(map[string]interface{}); ok {
				current = m[field]
			}

		}

		return New(current)

	}

	panic("objx: Get with invalid selector.")

}
