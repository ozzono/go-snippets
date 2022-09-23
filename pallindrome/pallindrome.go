/*
array[i] == array[len(array)-i-1]
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPallindromeInt(1235321))
	fmt.Println(isPallindromeInt(12355321))
	fmt.Println(isPallindromeInt(123532))
}

func isPallindromeInt(i int) bool {
	return isPallindromeTxt(fmt.Sprint(i))
}

func isPallindromeTxt(input string) bool {
	txtSlice := make([]string, len(input))
	for _, t := range fmt.Sprint(input) {
		txtSlice = append(txtSlice, string(t))
	}
	for i := 0; i < len(txtSlice)/2; i++ {
		if txtSlice[i] != txtSlice[len(txtSlice)-i-1] {
			return false
		}
	}
	s := strings.Contains("texto1", "1")
	return true
}
