package main

import "fmt"

const (
	max = int(10)
)

func main() {
	arr := []string{}
	arrsize := 112
	padPattern := pattern(len(fmt.Sprintf("%d", arrsize)))
	for i := 0; i < arrsize; i++ {
		arr = append(arr, fmt.Sprintf(padPattern, i))
	}
	maxSplit(arr)
}

func maxSplit(input []string) {
	for i := 0; i < len(input); {
		last := max
		if len(input[i:]) < max {
			last = len(input[i:])
		}
		fmt.Printf("%v\n", input[i:i+last])
		i += max
	}
}

func pattern(padsize int) string {
	return fmt.Sprintf("%s%s%s", "%0", fmt.Sprintf("%d", padsize), "d")
}
