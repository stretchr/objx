package objx

// IsNil gets whether the data is nil or not.
func (o *O) IsNil() bool {
	return o.Obj() == nil
}
