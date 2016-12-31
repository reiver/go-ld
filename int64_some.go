package ld

// Int64Some returns an Int64 with a value.
//
// For example:
//
//	var x ld.Int64 = ld.Int64Some(41)
func Int64Some(v int64) Int64 {
	return Int64 {
		value: v,
		loaded: true,
	}
}
