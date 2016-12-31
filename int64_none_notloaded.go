package ld

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
