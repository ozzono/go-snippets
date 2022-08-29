package main

import "fmt"

const defaultSize = 19

func main() {
	n := []int{}
	for i := 1; i <= defaultSize; i++ {
		n = append(n, i)
	}
	control := make(chan int, len(n))
	for i := range n {
		go func(i int) {
			if isPrime(i) {
				control <- i
				return
			}
			control <- 0
		}(n[i])
	}
	count := 0
	for i := 0; i < len(n); i++ {
		if c := <-control; c != 0 {
			count++
			fmt.Printf("%05d ", c)
			if count%10 == 0 {
				fmt.Println()
			}
		}
	}
	fmt.Println()
}

func isPrimeString(i int) string {
	if isPrime(i) {
		return "is"
	}
	return "is not"
}

func isPrime(i int) bool {
	if i == 1 {
		return true
	}
	if i <= 0 {
		return false
		/*
			- negative numbers can't be prime by deffinition
			    https://primes.utm.edu/notes/faq/negative_primes.html
			- zero is not prime
			    https://brilliant.org/wiki/is-0-prime/
		*/

	}
	control := make(chan bool, i-1)
	for j := 2; j <= i; j++ { // start from 2 because all numbers are divisible by 1
		go func(j int) {
			if i == j {
				// skip comparing the number with itself
				return
			}
			if i%j == 0 {
				control <- true
				return
			}
			control <- false
		}(j)
	}
	for j := 0; j < i-2; j++ {
		if <-control {
			return false
		}
	}
	return true
}
