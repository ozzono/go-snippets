package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(repeat(5))
}

func repeat(count int) string {
	output := []string{}
	for i := 0; i < count; i++ {
		output = append(output, "?")
	}
	return fmt.Sprintf("(%s)", strings.Join(output, ", "))
}
