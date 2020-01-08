package main

import "fmt"

func main() {
	AToBin("teste")
}

func AToBin(s string) {
	var binString string
	for _, c := range s {
		binString = fmt.Sprintf("%s%.8b", binString, c)
	}
	fmt.Println(binString)
}
