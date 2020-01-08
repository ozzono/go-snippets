package main

import "fmt"

func main() {
	fmt.Println(sort([]int{0, 9, 1, 8, 2, 7, 3, 6, 4, 5}))
}

func sort(pool []int) []int {
	for i := len(pool); i > 0; i-- {
		for j := 1; j < i; j++ {
			if pool[j-1] > pool[j] {
				fmt.Println("troca")
				tmp := pool[j]
				pool[j] = pool[j-1]
				pool[j-1] = tmp
			}
		}
	}
	return pool
}
