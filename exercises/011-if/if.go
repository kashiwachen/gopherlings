// Problem:
// `if` statements specify a block of code to execute given a condition is true.
// These come in a variety of flavours.
//
// Simple if statements:
//
//  if x > max {
//  	max = x
//  }
//

package main

import "fmt"

func main() {
	const tuesdayDiscount = 0.50 // 50% discount! Wow!
	isTuesday := true

	price := 32.00
	// fix this condition to accept numbers:
	if isTuesday {
		price = price * (1 - tuesdayDiscount)
	}
	fmt.Println("The price is", price)
}
