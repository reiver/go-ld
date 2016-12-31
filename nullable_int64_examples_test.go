package ld_test

import (
	"fmt"

	"github.com/reiver/go-ld"
)

func ExampleNullableInt64_notEqual() {

	var a ld.NullableInt64 = ld.NullableInt64Some(17)
	fmt.Printf("a = %v\n", a)

	var b ld.NullableInt64 = ld.NullableInt64NotLoaded()
	fmt.Printf("b = %v\n", b)

	if a == b {
		fmt.Println("They are equal.\n")
	}
	if a != b {
		fmt.Println("They are not equal.\n")
	}

	// Output:
	// a = NullableInt64Some(17)
	// b = NullableInt64NotLoaded()
	// They are not equal.
}

func ExampleNullableInt64_equal() {

	var a ld.NullableInt64 = ld.NullableInt64Some(12)
	fmt.Printf("a = %v\n", a)

	var b ld.NullableInt64 = ld.NullableInt64Some(12)
	fmt.Printf("b = %v\n", b)

	if a == b {
		fmt.Println("They are equal.\n")
	}
	if a != b {
		fmt.Println("They are not equal.\n")
	}

	// Output:
	// a = NullableInt64Some(12)
	// b = NullableInt64Some(12)
	// They are equal.
}

func ExampleNullableInt64_Match_noneNull() {

	switch x := ld.NullableInt64Null().Match().(type) {
	case ld.Int64Value:
		fmt.Printf("MATCHED Int64Value: (%T) %v\n", x, x)
	case ld.Int64NoneNotLoaded:
		fmt.Printf("MATCHED Int64NoneNotLoaded: (%T) %v\n", x, x)
	case ld.Int64NoneNull:
		fmt.Printf("MATCHED Int64NoneNull: (%T) %v\n", x, x)
	default:
		fmt.Println("Oops.")
	}


	var n ld.NullableInt64 = ld.NullableInt64Null()

	switch x := n.Match().(type) {
	case ld.Int64Value:
		fmt.Printf("MATCHED Int64Value: (%T) %v\n", x, x)
	case ld.Int64NoneNotLoaded:
		fmt.Printf("MATCHED Int64NoneNotLoaded: (%T) %v\n", x, x)
	case ld.Int64NoneNull:
		fmt.Printf("MATCHED Int64NoneNull: (%T) %v\n", x, x)
	default:
		fmt.Println("Oops.")
	}

	// Output:
	// MATCHED Int64NoneNull: (ld.Int64NoneNull) <Int64Null>
	// MATCHED Int64NoneNull: (ld.Int64NoneNull) <Int64Null>
}

func ExampleNullableInt64_Match_noneNotLoaded() {

	switch x := ld.NullableInt64NotLoaded().Match().(type) {
	case ld.Int64Value:
		fmt.Printf("MATCHED Int64Value: (%T) %v\n", x, x)
	case ld.Int64NoneNotLoaded:
		fmt.Printf("MATCHED Int64NoneNotLoaded: (%T) %v\n", x, x)
	case ld.Int64NoneNull:
		fmt.Printf("MATCHED Int64NoneNull: (%T) %v\n", x, x)
	default:
		fmt.Println("Oops.")
	}


	var n ld.NullableInt64 = ld.NullableInt64NotLoaded()

	switch x := n.Match().(type) {
	case ld.Int64Value:
		fmt.Printf("MATCHED Int64Value: (%T) %v\n", x, x)
	case ld.Int64NoneNotLoaded:
		fmt.Printf("MATCHED Int64NoneNotLoaded: (%T) %v\n", x, x)
	case ld.Int64NoneNull:
		fmt.Printf("MATCHED Int64NoneNull: (%T) %v\n", x, x)
	default:
		fmt.Println("Oops.")
	}

	// Output:
	// MATCHED Int64NoneNotLoaded: (ld.Int64NoneNotLoaded) <Int64NotLoaded>
	// MATCHED Int64NoneNotLoaded: (ld.Int64NoneNotLoaded) <Int64NotLoaded>
}

func ExampleNullableInt64_Match_some() {

	switch x := ld.NullableInt64Some(42).Match().(type) {
	case ld.Int64Value:
		fmt.Printf("MATCHED Int64Value: (%T) %v\n", x, x)
	case ld.Int64NoneNotLoaded:
		fmt.Printf("MATCHED Int64NoneNotLoaded: (%T) %v\n", x, x)
	case ld.Int64NoneNull:
		fmt.Printf("MATCHED Int64NoneNull: (%T) %v\n", x, x)
	default:
		fmt.Println("Oops.")
	}


	var n ld.NullableInt64 = ld.NullableInt64Some(42)

	switch x := n.Match().(type) {
	case ld.Int64Value:
		fmt.Printf("MATCHED Int64Value: (%T) %v\n", x, x)
	case ld.Int64NoneNotLoaded:
		fmt.Printf("MATCHED Int64NoneNotLoaded: (%T) %v\n", x, x)
	case ld.Int64NoneNull:
		fmt.Printf("MATCHED Int64NoneNull: (%T) %v\n", x, x)
	default:
		fmt.Println("Oops.")
	}

	// Output:
	// MATCHED Int64Value: (ld.Int64Value) 42
	// MATCHED Int64Value: (ld.Int64Value) 42
}

func ExampleNullableInt64_Match_someMath() {

	var a ld.NullableInt64 = ld.NullableInt64Some(2)

	switch x := a.Match().(type) {
	case ld.Int64Value:
		b := x + 1

		fmt.Printf("RESULT: (%T) %v\n", b, b)
	case ld.Int64NoneNotLoaded:
		fmt.Printf("MATCHED Int64NoneNotLoaded: %v\n", x)
	case ld.Int64NoneNull:
		fmt.Printf("MATCHED Int64NoneNull: (%T) %v\n", x, x)
	default:
		fmt.Println("Oops.")
	}

	// Output:
	// RESULT: (ld.Int64Value) 3
}
