package ld

// NullableInt64Null returns a NullableInt64 which does not have a value;
// and in particular the "null" type of a lack of a value.
//
// For example:
//
//	var x ld.NullableInt64 = ld.NullableInt64Null()
func NullableInt64Null() NullableInt64 {
	return NullableInt64{
		loaded: true,
		null: true,
	}
}
