package main

import "fmt"

func main() {
	for key := range map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	} {
		fmt.Println(key)
	}
}
