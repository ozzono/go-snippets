package main

import "fmt"

func main() {
	a := []string{"a", "b", "c"}
	fmt.Println(pick(a, -1))
	fmt.Println(pick(a, 0))
	fmt.Println(pick(a, 1))
	fmt.Println(pick(a, 2))
	fmt.Println(pick(a, 3))
}

func pick(input []string, cherry int) ([]string, error) {
	if cherry < 0 {
		return nil, fmt.Errorf("invalid index; must be >= 0")
	}
	if cherry+1 > len(input) {
		return nil, fmt.Errorf("index removal out of range")
	}
	return append(input[:cherry], input[cherry+1:]...), nil
}
