package main

import "fmt"

func main() {
	a := []byte("abc")
	b := []byte("bcd")
	for i := range a {
		fmt.Printf("a[%d]: %v\n", i, a[i])
		fmt.Printf("b[%d]: %v\n", i, b[i])
	}
}
