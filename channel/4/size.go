package main

import "fmt"

func main() {
	c := make(chan bool, 2)
	c <- true
	c <- false
	fmt.Printf("len %d cap %d\n", len(c), cap(c))
	<-c
	fmt.Printf("len %d cap %d\n", len(c), cap(c))
}
