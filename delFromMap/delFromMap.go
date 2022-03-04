package main

import "fmt"

func main() {
	m := map[int]int{1: 1}
	delete(m, 1)
	fmt.Println(m)
}
