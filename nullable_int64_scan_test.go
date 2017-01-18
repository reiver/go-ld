package ld

import (
	"math"

	"testing"
)

func TestNullableInt64ScanNilReceiver(t *testing.T) {

	value := int64(5)

	receiver := (*NullableInt64)(nil)

	err := receiver.Scan(value)
	if nil == err {
		t.Errorf("Expected an error, but did not actually get one: (%T) %v", err, err)
		return
	}

	if expected, actual := errNilReceiver, err; expected != actual {
		t.Errorf("Expected (%T) %v, but actually got (%T) %v.", expected, expected, actual, actual)
		return
	}
}

func TestNullableInt64ScanFail(t *testing.T) {

	tests := []struct{
		Value interface{}
	}{
		{
			Value: false,
		},
		{
			Value: true,
		},



		{
			Value: 7.2,
		},



		{
			Value: "apple banana cherry",
		},
		{
			Value: "5",
		},



		{
			Value: struct{ Value int64 }{Value: 5},
		},
	}


	for testNumber, test := range tests {

		var x NullableInt64

		if err := x.Scan(test.Value); nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually get one: (%T) %v", testNumber, err, err)
			continue
		}

		if _, ok := x.Match().(Int64NoneNotLoaded); !ok {
			t.Errorf("For test #%d, expected it to continue to not be loaded, but wasn't.", testNumber)
			continue
		}
	}
}

func TestNullableInt64Scan(t *testing.T) {

	tests := []struct{
		Value    interface{}
		Expected int64
	}{
		{
			Value:    int64(math.MinInt64),
			Expected: int64(math.MinInt64),
		},
		{
			Value:    int64(-1),
			Expected: int64(-1),
		},
		{
			Value:    int64(0),
			Expected: int64(0),
		},
		{
			Value:    int64(1),
			Expected: int64(1),
		},
		{
			Value:    int64(math.MaxInt64),
			Expected: int64(math.MaxInt64),
		},
	}

	{
		const numRand = 20
		for i:=0; i<numRand; i++ {
			{
				x := uint8(randomness.Int63n(math.MaxUint8))

				test := struct{
					Value    interface{}
					Expected int64
				}{
					Value:          x,
					Expected: int64(x),
				}
				tests = append(tests, test)
			}



			{
				x := uint16(randomness.Int63n(math.MaxUint16))

				test := struct{
					Value    interface{}
					Expected int64
				}{
					Value:          x,
					Expected: int64(x),
				}
				tests = append(tests, test)
			}



			{
				x := uint32(randomness.Int63n(math.MaxUint32))

				test := struct{
					Value    interface{}
					Expected int64
				}{
					Value:          x,
					Expected: int64(x),
				}
				tests = append(tests, test)
			}



			{
				x := int8(randomness.Int63n(math.MaxInt8))

				test := struct{
					Value    interface{}
					Expected int64
				}{
					Value:          x,
					Expected: int64(x),
				}
				tests = append(tests, test)
			}

			{
				x := -1 * int8(randomness.Int63n(math.MaxInt8))

				test := struct{
					Value    interface{}
					Expected int64
				}{
					Value:          x,
					Expected: int64(x),
				}
				tests = append(tests, test)
			}



			{
				x := int16(randomness.Int63n(math.MaxInt16))

				test := struct{
					Value    interface{}
					Expected int64
				}{
					Value:          x,
					Expected: int64(x),
				}
				tests = append(tests, test)
			}

			{
				x := -1 * int16(randomness.Int63n(math.MaxInt16))

				test := struct{
					Value    interface{}
					Expected int64
				}{
					Value:          x,
					Expected: int64(x),
				}
				tests = append(tests, test)
			}



			{
				x := int32(randomness.Int63n(math.MaxInt32))

				test := struct{
					Value    interface{}
					Expected int64
				}{
					Value:          x,
					Expected: int64(x),
				}
				tests = append(tests, test)
			}

			{
				x := -1 * int32(randomness.Int63n(math.MaxInt32))

				test := struct{
					Value    interface{}
					Expected int64
				}{
					Value:          x,
					Expected: int64(x),
				}
				tests = append(tests, test)
			}



			{
				x := randomness.Int63()

				test := struct{
					Value    interface{}
					Expected int64
				}{
					Value:    x,
					Expected: x,
				}
				tests = append(tests, test)
			}

			{
				x := -1 * randomness.Int63()

				test := struct{
					Value    interface{}
					Expected int64
				}{
					Value:    x,
					Expected: x,
				}
				tests = append(tests, test)
			}
		}
	}


	for testNumber, test := range tests {

		var x NullableInt64

		if err := x.Scan(test.Value); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		value, ok := x.Match().(Int64Value)
		if !ok {
			t.Errorf("For test #%d, expected it have a value, but doesn't.", testNumber)
			continue
		}

		datum := int64(value)

		if expected, actual := test.Expected, datum; expected != actual {
			t.Errorf("For test #%d, expected (%T) %v, but actually got (%T) %v.", testNumber, expected, expected, actual, actual)
			continue
		}
	}
}

func TestNullableInt64ScanNil(t *testing.T) {

	var x NullableInt64

	if err := x.Scan(nil); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	_, ok := x.Match().(Int64NoneNull)
	if !ok {
		t.Errorf("Expected it to be null, but wasn't.")
		return
	}
}
