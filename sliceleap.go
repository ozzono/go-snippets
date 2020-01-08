package main

import "fmt"

var padsize string

const (
	size    = 123
	rowsize = 10
)

func main() {

	padsize = fmt.Sprintf("%d", len(fmt.Sprintf("%d", size)))
	slice := fillslice(size)
	max := 0
	for i := 0; i < len(slice); i += rowsize {
		max = i + rowsize
		if max > len(slice)-1 {
			max = len(slice)
		}
		fmt.Printf("[%0"+padsize+"d:%0"+padsize+"d] slice: %v\n", i, max, slice[i:max])
	}
}

func fillslice(size int) []string {
	output := []string{}
	for i := 0; i < size; i++ {
		output = append(output, fmt.Sprintf("%0"+padsize+"d", i))
	}
	return output
}
