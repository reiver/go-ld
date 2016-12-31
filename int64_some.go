package ld

func Int64Some(v int64) Int64 {
	return Int64 {
		value: v,
		loaded: true,
	}
}
