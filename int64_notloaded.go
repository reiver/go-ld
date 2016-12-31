package ld

// Int64Some returns an Int64 which does not have a value;
// and in particular the "not loaded" type of a lack of a value.
func Int64NotLoaded() Int64 {
	return Int64{}
}

type Int64NoneNotLoaded struct{}

func (Int64NoneNotLoaded) Int64Matcher() {
	// Nothing here.
}

func (Int64NoneNotLoaded) String() string {
	return "<Int64NotLoaded>"
}
