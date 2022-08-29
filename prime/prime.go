package main

import "fmt"

func main() {
	for _, i := range []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
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
	count := 0
	for j := 2; j <= i; j++ { // start from 2 because all numbers are divisible by 1
		count++
		if i == j {
			continue // skip comparing the number with itself
		}
		if i%j == 0 {
			return false
		}
	}
	fmt.Println(i, "count", count)
	return true
}
