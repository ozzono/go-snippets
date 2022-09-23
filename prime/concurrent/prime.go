/*
Verify if a number is prime
Conditions:
- number divisible by 1 or by itself;
- must be above 0
- must be integer
*/

package main

import (
	"fmt"
	"log"
	"time"
)

const (
	defaultSize = 20000
	rowSize     = 20
)

func main() {
	start := time.Now()
	log.Println("open the gates")
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
	pad := len(fmt.Sprint(defaultSize / 10))
	if pad < 2 {
		pad = 2
	}
	leftPad := "%0" + fmt.Sprint(pad) + "d "
	for i := 0; i < len(n); i++ {
		if c := <-control; c != 0 {
			count++
			fmt.Printf(leftPad, c)
			if count%rowSize == 0 { // rows with rowSize columns
				fmt.Println("-", count/rowSize)
			}
		}
	}
	if count%rowSize != 0 {
		fmt.Println("-", count/rowSize+1)
	}
	fmt.Println(count, "numbers found")
	fmt.Printf("%dms\n", time.Now().Sub(start).Milliseconds())
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
		if j > 2 {
			j++
		}
	}

	for j := 0; j < (i-2)/2; j++ { // subtract one because starting from 2; subract one because skipping i==j
		if <-control {
			return false
		}
	}
	return true
}
