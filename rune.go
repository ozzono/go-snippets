package main

import "fmt"

func main() {
	rune("teste", 3)
	rune("Aav5HcyPVzj3", 11)
}

func rune(s string, max int) {
	if len(s) >= max {
		fmt.Printf("rune: %s@\n", s[:max])
	}
}
