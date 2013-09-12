package objx

import (
	"reflect"
)

// kind gets the reflect.Kind of the data stored in
// this object.
func (o *Obj) Kind() reflect.Kind {

	// use the cache if we have it
	if o.kind != reflect.Invalid {
		return o.kind
	}

	switch o.Obj().(type) {
	case bool:
		o.kind = reflect.Bool
	case int:
		o.kind = reflect.Int
	case int8:
		o.kind = reflect.Int8
	case int16:
		o.kind = reflect.Int16
	case int32:
		o.kind = reflect.Int32
	case int64:
		o.kind = reflect.Int64
	case uint:
		o.kind = reflect.Uint
	case uint8:
		o.kind = reflect.Uint8
	case uint16:
		o.kind = reflect.Uint16
	case uint32:
		o.kind = reflect.Uint32
	case uint64:
		o.kind = reflect.Uint64
	case float32:
		o.kind = reflect.Float32
	case float64:
		o.kind = reflect.Float64
	case complex64:
		o.kind = reflect.Complex64
	case complex128:
		o.kind = reflect.Complex128
	case string:
		o.kind = reflect.String
	case uintptr:
		o.kind = reflect.Uintptr
	default:
		// let reflect deal with this
		o.kind = reflect.ValueOf(o.Obj()).Kind()
	}

	return o.kind

}
