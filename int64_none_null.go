package ld

type Int64NoneNull struct{}

func (Int64NoneNull) String() string {
	return "<Int64Null>"
}

// NullableInt64Matcher makes Int64NoneNull fit the NullableInt64Matcher interface.
func (Int64NoneNull) NullableInt64Matcher() {
	// Nothing here.
}
