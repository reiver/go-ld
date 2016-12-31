package ld_test

import (
	"fmt"

	"github.com/reiver/go-ld"
)

func ExampleInt64_notEqual() {

	var a ld.Int64 = ld.Int64Some(17)
	fmt.Printf("a = %v\n", a)

	var b ld.Int64 = ld.Int64NotLoaded()
	fmt.Printf("b = %v\n", b)

	if a == b {
		fmt.Println("They are equal.\n")
	}
	if a != b {
		fmt.Println("They are not equal.\n")
	}

	// Output:
	// a = Int64Some(17)
	// b = Int64NotLoaded()
	// They are not equal.
}

func ExampleInt64_equal() {

	var a ld.Int64 = ld.Int64Some(12)
	fmt.Printf("a = %v\n", a)

	var b ld.Int64 = ld.Int64Some(12)
	fmt.Printf("b = %v\n", b)

	if a == b {
		fmt.Println("They are equal.\n")
	}
	if a != b {
		fmt.Println("They are not equal.\n")
	}

	// Output:
	// a = Int64Some(12)
	// b = Int64Some(12)
	// They are equal.
}

func ExampleInt64_Match_noneNotLoaded() {

	switch x := ld.Int64NotLoaded().Match().(type) {
	case ld.Int64Value:
		fmt.Printf("MATCHED Int64Value: (%T) %v\n", x, x)
	case ld.Int64NoneNotLoaded:
		fmt.Printf("MATCHED Int64NoneNotLoaded: (%T) %v\n", x, x)
	default:
		fmt.Println("Oops.")
	}


	var n ld.Int64 = ld.Int64NotLoaded()

	switch x := n.Match().(type) {
	case ld.Int64Value:
		fmt.Printf("MATCHED Int64Value: (%T) %v\n", x, x)
	case ld.Int64NoneNotLoaded:
		fmt.Printf("MATCHED Int64NoneNotLoaded: (%T) %v\n", x, x)
	default:
		fmt.Println("Oops.")
	}

	// Output:
	// MATCHED Int64NoneNotLoaded: (ld.Int64NoneNotLoaded) <Int64NotLoaded>
	// MATCHED Int64NoneNotLoaded: (ld.Int64NoneNotLoaded) <Int64NotLoaded>
}

func ExampleInt64_Match_some() {

	switch x := ld.Int64Some(42).Match().(type) {
	case ld.Int64Value:
		fmt.Printf("MATCHED Int64Value: (%T) %v\n", x, x)
	case ld.Int64NoneNotLoaded:
		fmt.Printf("MATCHED Int64NoneNotLoaded: (%T) %v\n", x, x)
	default:
		fmt.Println("Oops.")
	}


	var n ld.Int64 = ld.Int64Some(42)

	switch x := n.Match().(type) {
	case ld.Int64Value:
		fmt.Printf("MATCHED Int64Value: (%T) %v\n", x, x)
	case ld.Int64NoneNotLoaded:
		fmt.Printf("MATCHED Int64NoneNotLoaded: (%T) %v\n", x, x)
	default:
		fmt.Println("Oops.")
	}

	// Output:
	// MATCHED Int64Value: (ld.Int64Value) 42
	// MATCHED Int64Value: (ld.Int64Value) 42
}

func ExampleInt64_Match_someMath() {

	var a ld.Int64 = ld.Int64Some(2)

	switch x := a.Match().(type) {
	case ld.Int64Value:
		b := x + 1

		fmt.Printf("RESULT: (%T) %v\n", b, b)
	case ld.Int64NoneNotLoaded:
		fmt.Printf("MATCHED Int64NoneNotLoaded: %v\n", x)
	default:
		fmt.Println("Oops.")
	}

	// Output:
	// RESULT: (ld.Int64Value) 3
}
