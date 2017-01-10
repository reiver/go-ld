package ld

import (
	"testing"
)

func TestInt64Int64(t *testing.T) {

	{
		var x Int64

		i64, err := x.Int64()
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one: (%T) %v; got value: %d", err, err, i64)
			return
		}
		if _, ok := err.(NotLoadedComplainer); !ok {
			t.Errorf("Expected error to fit \"ld.NotLoadedComplainer\", but actually didn't.")
			return
		}
	}

	{
		expected := int64(5)

		var x Int64 = Int64Some(expected)

		actual, err := x.Int64()
		if nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
			return
		}

		if expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}
	}
}
