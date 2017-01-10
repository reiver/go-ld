package ld

import (
	"github.com/reiver/go-cast"

	"fmt"
)

// Int64 is this package's version of Go's built-in int64 type.
//
// A difference being that this package's Int64 has the ability to express (a certain kind of)
// a lack of a value.
//
// More specifically, that the value has not been loaded.
//
// (For example, that the value for this field was not loaded from the database.)
//
// To assign a value to something of type Int64 you would do something like the following:
//
//	var i64 ld.Int64 = ld.Int64Some(17)
//
// To assign the "not loaded" version of a "lack of a value" to a something of type Int64
// you would do something like the following:
//
//	var i64 ld.Int64 = Int64NotLoaded()
//
// To see if one thing of type Int64 is equal to something else of type Int64, you can do something
// like the following:
//
//	var a ld.Int64
//	// ...
//	
//	var b ld.Int64
//	// ...
//	
//	if a == b {
//		// ...
//	}
//
// (I.e., you can use the normal Go "==" operator for checking for equality.)
//
// To see if one thing of type Int64 is not equal to something else of type Int64, you can do something
// like the following:
//
//	var a ld.Int64
//	// ...
//	
//	var b ld.Int64
//	// ...
//	
//	if a != b {
//		// ...
//	}
//
// (I.e., you can you the normal Go "!=" operator for checking for inequality.)
//
// So that means you can do things such as:
//
//	if i64 == ld.Int64Some(5) {
//		// ...
//	}
//
// And:
//
//	if i64 == ld.Int64NotLoaded() {
//		// ...
//	}
//
// To unwrap an Int64, and (potentially) get at the number (the int64) inside of it,
// if indeed the Int64 has a value, you could do something like the following:
//
//	var i64 ld.Int64
//
//	switch n := i64.Match().(type) {
//	case ld.Int64Value:
//		// ...
//
//	case ld.Int64NoneNotLoaded:
//		// ...
//
//	default:
//		// ...
//	}
type Int64 struct {
	value int64
	loaded bool
}

func (receiver Int64) Match() Int64Matcher {
	if !receiver.loaded {
		return Int64NoneNotLoaded{}
	}

	return Int64Value(receiver.value)
}

func (receiver Int64) String() string {
	if !receiver.loaded {
		return "Int64NotLoaded()"
	}

	return fmt.Sprintf("Int64Some(%d)", receiver.value)
}

func (receiver Int64) Int64() (int64, error) {
	if !receiver.loaded {
		return 0, errNotLoaded
	}

	return receiver.value, nil
}

func (receiver *Int64) Scan(v interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	i64, err := cast.Int64(v)
	if nil != err {
		return err
	}

	receiver.value = i64
	receiver.loaded = true

	return nil
}
