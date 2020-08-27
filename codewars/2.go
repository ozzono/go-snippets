// Write a function that accepts an array of 10 integers (between 0 and 9), that returns a string of those numbers in the form of a phone number.
// Example:
// CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})  // returns "(123) 456-7890"
// The returned format must be correct in order to complete this challenge.
// Don't forget the space after the closing parentheses!

package main

import (
	"fmt"
)

func main() {
	numbers := [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// fmt.Println(numbers[:3])
	// fmt.Println(numbers[3:])
	fmt.Println(CreatePhoneNumber(numbers))
}

func CreatePhoneNumber(numbers [10]uint) string {
	output := ""
	for i := range numbers {
		output += fmt.Sprintf("%d", numbers[i])
	}
	return fmt.Sprintf("(%s) %s-%s", output[:3], output[3:6], output[6:])
}
