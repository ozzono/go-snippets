package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Printf(pattern(3, pot(10, i)), pot(10, i))
	}
}

func pattern(padsize, element int) string {
	pattern := "%d"
	if padsize > 0 {
		pattern = "%0" + fmt.Sprintf("%d", padsize) + "d"
		fmt.Println(pattern)
	}
	return pattern
}

func pot(base, exp int) int {
	if exp == 0 {
		return 1
	}
	for i := 0; i < exp-1; i++ {
		base = base * base
	}
	return base
}
