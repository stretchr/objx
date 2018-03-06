package objx

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/url"
)

// SignatureSeparator is the character that is used to
// separate the Base64 string from the security signature.
const SignatureSeparator = "_"

var QuerySliceKeySuffix = "[]"

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

	m.parseUrlQuery(m, vals, "")

	return vals
}

func (m Map) parseUrlQuery(queryMap Map, vals url.Values, key string) {
	for k, v := range queryMap {
		val := &Value{data: v}
		switch {
        case val.IsObjxMap():
        	if key == "" {
	        	m.parseUrlQuery(v.(Map), vals, k)
	        } else {
	        	m.parseUrlQuery(v.(Map), vals, key + "[" + k + "]")
        	}
        case val.IsObjxMapSlice():
            sliceKey := k + QuerySliceKeySuffix
        	if key != "" {
        		sliceKey = key + "[" + k + "]" + QuerySliceKeySuffix
        	}

			for _, sv := range val.MustObjxMapSlice() {
	        	m.parseUrlQuery(sv, vals, sliceKey)
			}
        case val.IsMSI():
        	if key == "" {
	        	m.parseUrlQuery(New(v), vals, k)
	        } else {
	        	m.parseUrlQuery(New(v), vals, key + "[" + k + "]")
        	}
		case val.IsMSISlice():
            sliceKey := k + QuerySliceKeySuffix
        	if key != "" {
        		sliceKey = key + "[" + k + "]" + QuerySliceKeySuffix
        	}

			for _, sv := range val.MustMSISlice() {
	        	m.parseUrlQuery(New(sv), vals, sliceKey)
			}
        case val.IsStrSlice(), val.IsBoolSlice(),
            val.IsFloat32Slice(), val.IsFloat64Slice(),
            val.IsIntSlice(), val.IsInt8Slice(), val.IsInt16Slice(), val.IsInt32Slice(), val.IsInt64Slice(),
            val.IsUintSlice(), val.IsUint8Slice(), val.IsUint16Slice(), val.IsUint32Slice(), val.IsUint64Slice():

            sliceKey := k + QuerySliceKeySuffix
        	if key != "" {
        		sliceKey = key + "[" + k + "]" + QuerySliceKeySuffix
        	}

			vals[sliceKey] = val.StringSlice()
        default:
        	if key == "" {
				vals.Set(k, val.String())
	        } else {
				vals.Set(key + "[" + k + "]", val.String())
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
