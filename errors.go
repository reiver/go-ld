package ld

import (
	"errors"
)

var (
	errNilReceiver = errors.New("Nil Receiver")
	errNotLoaded   = internalNotLoadedComplainer{}
	errNull        = internalNullComplainer{}
)
