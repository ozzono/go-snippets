package main

import "fmt"

func main() {
	s2time(3601)
	s2time(601)
	s2time(4601)
	s2time(1234)
	s2time(4321)
	s2time(987)
	s2time(2687)
}

func s2time(input int) {
	duration := input
	totalh, restoh := div(duration, 3600)
	totalm, restom := div(restoh, 60)
	fmt.Printf("%04ds or %02dh %02dm %02ds\n", input, totalh, totalm, restom)
}

func div(input, div int) (int, int) {
	return ((input - (input % div)) / div), (input % div)
}
