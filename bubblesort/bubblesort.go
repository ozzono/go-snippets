package main

import (
	"fmt"
)

func main() {
	teste := []int{9, 1, 0, 2, 8, 3, 7, 4, 5, 6}
	fmt.Printf("%v,\n", teste)
	fmt.Printf("%v,\n", intSort(teste))
}

func stringSort(pool []string) []string {
	for i := len(pool); i > 0; i-- {
		for j := 1; j < i; j++ {
			if pool[j-1] > pool[j] {
				tmp := pool[j]
				pool[j] = pool[j-1]
				pool[j-1] = tmp
			}
		}
	}
	return pool
}

func intSort(pool []int) []int {
	for i := len(pool); i > 0; i-- {
		for j := 1; j < i; j++ {
			if pool[j-1] < pool[j] {
				tmp := pool[j]
				pool[j] = pool[j-1]
				pool[j-1] = tmp
			}
		}
	}
	return pool
}
