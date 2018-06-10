package objx

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
)

// SignatureSeparator is the character that is used to
// separate the Base64 string from the security signature.
const SignatureSeparator = "_"

// URLValuesSliceKeySuffix is the character that is used to
// specify a suffic for slices parsed by URLValues.
// If the suffix is set to "[i]", then the index of the slice
// is used in place of i
// Ex: Suffix "[]" would have the form a[]=b&a[]=c
// OR Suffix "[i]" would have the form a[0]=b&a[1]=c
// OR Suffix "" would have the form a=b&a=c
var urlValuesSliceKeySuffix = "[]"

const (
	URLValuesSliceKeySuffixEmpty = ""
	URLValuesSliceKeySuffixArray = "[]"
	URLValuesSliceKeySuffixIndex = "[i]"
)

// SetURLValuesSliceKeySuffix sets the character that is used to
// specify a suffic for slices parsed by URLValues.
// If the suffix is set to "[i]", then the index of the slice
// is used in place of i
// Ex: Suffix "[]" would have the form a[]=b&a[]=c
// OR Suffix "[i]" would have the form a[0]=b&a[1]=c
// OR Suffix "" would have the form a=b&a=c
func SetURLValuesSliceKeySuffix(s string) {
	urlValuesSliceKeySuffix = s
}

// JSON converts the contained object to a JSON string
// representation
func (m Map) JSON() (string, error) {
	result, err := json.Marshal(m)
	if err != nil {
		err = errors.New("objx: JSON encode failed with: " + err.Error())
	}
	return string(result), err
}

// MustJSON converts the contained object to a JSON string
// representation and panics if there is an error
func (m Map) MustJSON() string {
	result, err := m.JSON()
	if err != nil {
		panic(err.Error())
	}
	return result
}

// Base64 converts the contained object to a Base64 string
// representation of the JSON string representation
func (m Map) Base64() (string, error) {
	var buf bytes.Buffer

	jsonData, err := m.JSON()
	if err != nil {
		return "", err
	}

	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	_, _ = encoder.Write([]byte(jsonData))
	_ = encoder.Close()

	return buf.String(), nil
}

// MustBase64 converts the contained object to a Base64 string
// representation of the JSON string representation and panics
// if there is an error
func (m Map) MustBase64() string {
	result, err := m.Base64()
	if err != nil {
		panic(err.Error())
	}
	return result
}

// SignedBase64 converts the contained object to a Base64 string
// representation of the JSON string representation and signs it
// using the provided key.
func (m Map) SignedBase64(key string) (string, error) {
	base64, err := m.Base64()
	if err != nil {
		return "", err
	}

	sig := HashWithKey(base64, key)
	return base64 + SignatureSeparator + sig, nil
}

// MustSignedBase64 converts the contained object to a Base64 string
// representation of the JSON string representation and signs it
// using the provided key and panics if there is an error
func (m Map) MustSignedBase64(key string) string {
	result, err := m.SignedBase64(key)
	if err != nil {
		panic(err.Error())
	}
	return result
}

/*
	URL Query
	------------------------------------------------
*/

// URLValues creates a url.Values object from an Obj. This
// function requires that the wrapped object be a map[string]interface{}
func (m Map) URLValues() url.Values {
	vals := make(url.Values)

	m.parseURLValues(m, vals, "")

	return vals
}

func (m Map) parseURLValues(queryMap Map, vals url.Values, key string) {
	useSliceIndex := false
	if urlValuesSliceKeySuffix == "[i]" {
		useSliceIndex = true
	}

	for k, v := range queryMap {
		val := &Value{data: v}
		switch {
		case val.IsObjxMap():
			if key == "" {
				m.parseURLValues(val.ObjxMap(), vals, k)
			} else {
				m.parseURLValues(val.ObjxMap(), vals, key+"["+k+"]")
			}
		case val.IsObjxMapSlice():
			sliceKey := k
			if key != "" {
				sliceKey = key + "[" + k + "]"
			}

			if useSliceIndex {
				for i, sv := range val.MustObjxMapSlice() {
					sk := sliceKey + "[" + strconv.FormatInt(int64(i), 10) + "]"
					m.parseURLValues(sv, vals, sk)
				}
			} else {
				sliceKey = sliceKey + urlValuesSliceKeySuffix
				for _, sv := range val.MustObjxMapSlice() {
					m.parseURLValues(sv, vals, sliceKey)
				}
			}
		case val.IsMSISlice():
			sliceKey := k
			if key != "" {
				sliceKey = key + "[" + k + "]"
			}

			if useSliceIndex {
				for i, sv := range val.MustMSISlice() {
					sk := sliceKey + "[" + strconv.FormatInt(int64(i), 10) + "]"
					m.parseURLValues(New(sv), vals, sk)
				}
			} else {
				sliceKey = sliceKey + urlValuesSliceKeySuffix
				for _, sv := range val.MustMSISlice() {
					m.parseURLValues(New(sv), vals, sliceKey)
				}
			}
		case val.IsStrSlice(), val.IsBoolSlice(),
			val.IsFloat32Slice(), val.IsFloat64Slice(),
			val.IsIntSlice(), val.IsInt8Slice(), val.IsInt16Slice(), val.IsInt32Slice(), val.IsInt64Slice(),
			val.IsUintSlice(), val.IsUint8Slice(), val.IsUint16Slice(), val.IsUint32Slice(), val.IsUint64Slice():

			sliceKey := k
			if key != "" {
				sliceKey = key + "[" + k + "]"
			}

			if useSliceIndex {
				for i, sv := range val.StringSlice() {
					sk := sliceKey + "[" + strconv.FormatInt(int64(i), 10) + "]"
					vals.Set(sk, sv)
				}
			} else {
				sliceKey = sliceKey + urlValuesSliceKeySuffix
				vals[sliceKey] = val.StringSlice()
			}

		default:
			if key == "" {
				vals.Set(k, val.String())
			} else {
				vals.Set(key+"["+k+"]", val.String())
			}
		}
	}
}

// URLQuery gets an encoded URL query representing the given
// Obj. This function requires that the wrapped object be a
// map[string]interface{}
func (m Map) URLQuery() (string, error) {
	return m.URLValues().Encode(), nil
}
