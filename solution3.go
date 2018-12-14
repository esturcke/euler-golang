package euler

import (
	"strconv"
)

/*
# Largest prime factor

[Problem 3](http://projecteuler.net/problem=3)

   The prime factors of 13195 are 5, 7, 13 and 29.

   What is the largest prime factor of the number 600851475143 ?

*/
func solve3() string {
	factors := make(chan int)
	go primeFactors(600851475143, factors)
	max := 0
	for factor := range factors {
		max = factor
	}
	return strconv.Itoa(max)
}

func primeFactors(n int, out chan<- int) {
	defer close(out)
	for i := 2; n > 1; i++ {
		for n%i == 0 {
			out <- i
			n /= i
		}
	}
}
