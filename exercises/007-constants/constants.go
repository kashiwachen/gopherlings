//
// Problem:
// In Go constants are just what the name suggests, constant- unchangeable,
// immutable.
// Go supports constant string, numeric, boolean values.
//
// You can't modify or reassign to a constant in Go.

package main

import (
	"fmt"
	"math"
)

func main() {
	// In Go variable names must start with a letter.
	const num_20 = 20
	fmt.Println(math.Sin(num_20))

	const abc = "abc"
	fmt.Println(abc)

	sum := abc + "tyk"
	fmt.Println(sum)

	// Something's wrong here- an attempt to
	// assign to a constant will cause a compiler error.
	// Instead of assigning to an existing constant, declare a new one!
	sum = sum + "cheeky addition"
	fmt.Println(sum)
}
