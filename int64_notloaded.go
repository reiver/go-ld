package ld

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
