package ld

type NullComplainer interface {
	error
	NullComplainer()
}

type internalNullComplainer struct{}

func (receiver internalNullComplainer) Error() string {
	return "Null"
}

func (internalNullComplainer) NullComplainer() {
	// Nothing here.
}
