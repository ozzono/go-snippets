package main

import "fmt"

func main() {

	type test struct {
		a string
		b int
	}

	var t test
	fmt.Println(t == nil)
}
