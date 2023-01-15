//
// Problem:
// Functions may have multiple arguments and multiple return values.
// To create a function with multiple return values surround the
// return types in parentheses.
//
//  func fruits() (string, string) {
//  	return "banana", "tangerine"
//  }
//
// If two or more consecutive arguments have the same type one can omit
// the type name until the last argument.
//
//  func divmod(a, b int) (div, mod int) {
//  	div = a/b
//  	mod = a%b
//  	return div, mod
//  }
//

package main

import "fmt"

func main() {
	var a, b int = 1, 2
	var c, d float32 = 1, 2
	fmt.Println(return1And2(a, b, c, d))
}

// This function should return two values, 1 and 2.
// Fix the function signature (the line with the `func` keyword)
// by adding the two returned parameters with names like
// the last example in the problem description.
// TODO: Suggest doing this exercise for getting familiar with type omitting feature
func return1And2(a, b int, c, d float32) (int, int, float32, float32) {
	return a, b, c, d
}
