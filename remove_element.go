package main

import "fmt"

func main() {
	slice := []string{"a", "b", "c", "d", "e"}
	fmt.Println(slice)
	removeElement(slice, 3)
	removeElement(slice, 1)
}

func removeElement(slice []string, index int) []string {
	fmt.Printf("input:  %v\n", slice)
	output := append(slice[:index], slice[index+1:]...)
	fmt.Printf("output: %v\n", output)
	return output
}
