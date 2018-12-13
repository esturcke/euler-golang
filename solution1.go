package euler

import "strconv"

/*
# Multiples of 3 and 5

[Problem 1](http://projecteuler.net/problem=1)

If we list all the natural numbers below 10 that are multiples of 3 or 5,
we get 3, 5, 6 and 9. The sum of these multiples is 23. Find the sum of all
the multiples of 3 or 5 below 1000.
*/
func solve1() string {
	sum := 0
	for n := 1; n < 1000; n++ {
		if n%3 == 0 || n%5 == 0 {
			sum += n
		}
	}
	return strconv.Itoa(sum)
}
