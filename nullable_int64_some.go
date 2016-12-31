package ld

// NullableInt64Some returns a NullableInt64 with a value.
//
// For example:
//
//	var x ld.NullableInt64 = ld.NullableInt64Some(41)
func NullableInt64Some(v int64) NullableInt64 {
	return NullableInt64 {
		value: v,
		loaded: true,
		null: false,
	}
}
