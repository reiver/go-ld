package ld

// Int64NotLoaded returns an Int64 which does not have a value;
// and in particular the "not loaded" type of a lack of a value.
//
// For example:
//
//	var x ld.Int64 = ld.Int64NotLoaded()
func Int64NotLoaded() Int64 {
	return Int64{}
}
