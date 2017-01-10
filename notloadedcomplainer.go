package ld

type NotLoadedComplainer interface {
	error
	NotLoadedComplainer()
}

type internalNotLoadedComplainer struct{}

func (receiver internalNotLoadedComplainer) Error() string {
	return "Not Loaded"
}

func (internalNotLoadedComplainer) NotLoadedComplainer() {
	// Nothing here.
}
