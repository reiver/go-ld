package ld

type Int64Value int64

// Int64Matcher makes Int64Value fit the Int64Matcher interface.
func (Int64Value) Int64Matcher() {
	// Nothing here.
}

// NullableInt64Matcher makes Int64Value fit the NullableInt64Matcher interface.
func (Int64Value) NullableInt64Matcher() {
	// Nothing here.
}

