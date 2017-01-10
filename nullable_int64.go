package ld

import (
	"fmt"
)

type NullableInt64 struct {
	value int64
	loaded bool
	null bool
}

func (receiver NullableInt64) Match() NullableInt64Matcher {
	if !receiver.loaded {
		return Int64NoneNotLoaded{}
	}

	if receiver.null {
		return Int64NoneNull{}
	}

	return Int64Value(receiver.value)
}

func (receiver NullableInt64) String() string {
	if !receiver.loaded {
		return "NullableInt64NotLoaded()"
	}

	if receiver.null {
		return "NullableInt64Null()"
	}

	return fmt.Sprintf("NullableInt64Some(%d)", receiver.value)
}

func (receiver NullableInt64) Int64() (int64, error) {
	if !receiver.loaded {
		return 0, errNotLoaded
	}

	if receiver.null {
		return 0, errNull
	}

	return receiver.value, nil
}
