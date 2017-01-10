package ld

import (
	"testing"
)

func TestNullableInt64Int64(t *testing.T) {

	{
		var x NullableInt64

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
		var x NullableInt64 = NullableInt64Null()

		i64, err := x.Int64()
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one: (%T) %v; got value: %d", err, err, i64)
			return
		}
		if _, ok := err.(NullComplainer); !ok {
			t.Errorf("Expected error to fit \"ld.NullComplainer\", but actually didn't.")
			return
		}
	}

	{
		expected := int64(5)

		var x NullableInt64 = NullableInt64Some(expected)

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
