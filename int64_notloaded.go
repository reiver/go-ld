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

type Int64NoneNotLoaded struct{}

func (Int64NoneNotLoaded) String() string {
	return "<Int64NotLoaded>"
}

// Int64Matcher makes Int64NoneNotLoaded fit the Int64Matcher interface.
func (Int64NoneNotLoaded) Int64Matcher() {
	// Nothing here.
}

// NullableInt64Matcher makes Int64NoneNotLoaded fit the NullableInt64Matcher interface.
func (Int64NoneNotLoaded) NullableInt64Matcher() {
	// Nothing here.
}
