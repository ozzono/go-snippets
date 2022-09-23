package main

import "fmt"

func main() {
	n := []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range n {
		fmt.Printf("%02d %s prime\n", i, isPrimeString(i))
	}

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

	if i%2 == 0 {
		if i == 2 {
			return true
		}
		return false
	}

	for j := 3; j <= i; j++ { // start from 2 because all numbers are divisible by 1
		if i == j {
			continue // skip comparing the number with itself
		}
		if i%j == 0 {
			return false
		}
		j++
	}
	return true
}
