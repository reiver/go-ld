package ld

// NullableInt64NotLoaded returns a NullableInt64 which does not have a value;
// and in particular the "not loaded" type of a lack of a value.
//
// For example:
//
//	var x ld.NullableInt64 = ld.NullableInt64NotLoaded()
func NullableInt64NotLoaded() NullableInt64 {
	return NullableInt64{}
}
