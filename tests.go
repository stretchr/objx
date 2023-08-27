package objx

// IsNil gets whether the data is nil or not.
func (v *Value) IsNil() bool {
	return v == nil || v.data == nil
}
