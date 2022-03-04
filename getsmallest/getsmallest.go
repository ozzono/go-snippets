package main

import "fmt"

func main() {
	smallest([]int{9, 2, 8, 3, 7, 4, 6, 5, 0, 1})
}

func smallest(pool []int) {
	index := 0
	for i, _ := range pool {
		if pool[i] < pool[index] {
			index = i
		}
	}
	fmt.Println(pool[index])
}
