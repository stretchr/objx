// This file provides minimal assert/require helpers built on Go's standard library to
// remove the dependency on github.com/stretchr/testify from objx tests.
package objx_test

import (
	"fmt"
	"reflect"
	"testing"
)

type nonfatal struct{}
type fatal struct{}

var assert nonfatal
var require fatal

func (nonfatal) fail(t *testing.T, msg string, msgAndArgs ...interface{}) bool {
	t.Helper()
	if len(msgAndArgs) > 0 {
		msg = fmt.Sprintf(msg, msgAndArgs...)
	}
	t.Error(msg)
	return false
}

func (fatal) fail(t *testing.T, msg string, msgAndArgs ...interface{}) {
	t.Helper()
	if len(msgAndArgs) > 0 {
		msg = fmt.Sprintf(msg, msgAndArgs...)
	}
	t.Fatal(msg)
}

func (a nonfatal) Equal(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		return a.fail(t, "not equal:\nexpected: %#v\nactual:   %#v", expected, actual)
	}
	return true
}
func (a fatal) Equal(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		a.fail(t, "not equal:\nexpected: %#v\nactual:   %#v", expected, actual)
	}
}

func (a nonfatal) NotEqual(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()
	if reflect.DeepEqual(expected, actual) {
		return a.fail(t, "should not be equal: %#v", actual)
	}
	return true
}

func (a fatal) NotEqual(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if reflect.DeepEqual(expected, actual) {
		a.fail(t, "should not be equal: %#v", actual)
	}
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	}
	return false
}

func (a nonfatal) Nil(t *testing.T, object interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()
	if !isNil(object) {
		return a.fail(t, "expected nil, got: %#v", object)
	}
	return true
}

func (a fatal) Nil(t *testing.T, object interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if !isNil(object) {
		a.fail(t, "expected nil, got: %#v", object)
	}
}

func (a nonfatal) NotNil(t *testing.T, object interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()
	if isNil(object) {
		return a.fail(t, "expected not nil")
	}
	return true
}

func (a fatal) NotNil(t *testing.T, object interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if isNil(object) {
		a.fail(t, "expected not nil")
	}
}

func (a nonfatal) True(t *testing.T, value bool, msgAndArgs ...interface{}) bool {
	t.Helper()
	if !value {
		return a.fail(t, "expected true, got false")
	}
	return true
}
func (a fatal) True(t *testing.T, value bool, msgAndArgs ...interface{}) {
	t.Helper()
	if !value {
		a.fail(t, "expected true, got false")
	}
}

func (a nonfatal) False(t *testing.T, value bool, msgAndArgs ...interface{}) bool {
	t.Helper()
	if value {
		return a.fail(t, "expected false, got true")
	}
	return true
}
func (a fatal) False(t *testing.T, value bool, msgAndArgs ...interface{}) {
	t.Helper()
	if value {
		a.fail(t, "expected false, got true")
	}
}

func (a nonfatal) NoError(t *testing.T, err error, msgAndArgs ...interface{}) bool {
	t.Helper()
	if err != nil {
		return a.fail(t, "expected no error, got: %v", err)
	}
	return true
}
func (a fatal) NoError(t *testing.T, err error, msgAndArgs ...interface{}) {
	t.Helper()
	if err != nil {
		a.fail(t, "expected no error, got: %v", err)
	}
}

func (a nonfatal) Error(t *testing.T, err error, msgAndArgs ...interface{}) bool {
	t.Helper()
	if err == nil {
		return a.fail(t, "expected error, got nil")
	}
	return true
}
func (a fatal) Error(t *testing.T, err error, msgAndArgs ...interface{}) {
	t.Helper()
	if err == nil {
		a.fail(t, "expected error, got nil")
	}
}

func (a nonfatal) Panics(t *testing.T, f func(), msgAndArgs ...interface{}) bool {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			_ = a.fail(t, "expected panic, but function did not panic")
		}
	}()
	f()
	return true
}

func (a fatal) Panics(t *testing.T, f func(), msgAndArgs ...interface{}) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			a.fail(t, "expected panic, but function did not panic")
		}
	}()
	f()
}

func (a nonfatal) Empty(t *testing.T, object interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()
	if !isEmpty(object) {
		return a.fail(t, "expected empty, got: %#v", object)
	}
	return true
}

func (a fatal) Empty(t *testing.T, object interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if !isEmpty(object) {
		a.fail(t, "expected empty, got: %#v", object)
	}
}

func isEmpty(object interface{}) bool {
	if object == nil {
		return true
	}
	v := reflect.ValueOf(object)
	switch v.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan, reflect.String:
		return v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		if v.IsNil() {
			return true
		}
		return isEmpty(v.Elem().Interface())
	}
	// numbers and structs are never considered empty here
	return false
}

func (a nonfatal) Len(t *testing.T, object interface{}, length int, msgAndArgs ...interface{}) bool {
	t.Helper()
	v := reflect.ValueOf(object)
	switch v.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan, reflect.String:
		if v.Len() != length {
			return a.fail(t, "unexpected length, expected %d got %d", length, v.Len())
		}
		return true
	default:
		return a.fail(t, "Len not supported for kind %s", v.Kind())
	}
}

func (a fatal) Len(t *testing.T, object interface{}, length int, msgAndArgs ...interface{}) {
	t.Helper()
	v := reflect.ValueOf(object)
	switch v.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan, reflect.String:
		if v.Len() != length {
			a.fail(t, "unexpected length, expected %d got %d", length, v.Len())
		}
	default:
		a.fail(t, "Len not supported for kind %s", v.Kind())
	}
}
