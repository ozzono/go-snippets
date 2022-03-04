package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
}

func factorial(n int) int {
	out := 1
	for i := 1; i <= n; i++ {
		out *= i
	}
	return out
}
