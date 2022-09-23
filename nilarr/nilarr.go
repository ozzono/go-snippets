package main

import "fmt"

func main() {
	// arr := make([]int, 0)
	arr(nil)
	arr(make([]int, 0, 0))
	arr([]int{})
}

func arr(a []int) {
	fmt.Printf("%#v\n", a)
	fmt.Printf("%d\n", len(a))
}
