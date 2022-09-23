package main

import "fmt"

func main() {
	m := map[int]string{
		30: "tamara",
		34: "hugo",
	}
	for key, value := range m {
		fmt.Printf("key: %v - value %v\n", key, value)
	}
}
