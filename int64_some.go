package ld

// Int64Some returns an Int64 with the value.
func Int64Some(v int64) Int64 {
	return Int64 {
		value: v,
		loaded: true,
	}
}
