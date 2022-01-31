// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Crunch the primes
//
//  1. Get numbers from the command-line.
//  2. `for range` over them.
//  4. Convert each one to an int.
//  5. If one of the numbers isn't an int, just skip it.
//  6. Print the ones that are only the prime numbers.
//
// RESTRICTION
//  The user can run the program with any number of
//  arguments.
//
// HINT
//  Find whether a number is prime using this algorithm:
//  https://stackoverflow.com/a/1801446/115363
//
// EXPECTED OUTPUT
//  go run main.go 2 4 6 7 a 9 c 11 x 12 13
//    2 7 11 13
//
//  go run main.go 1 2 3 5 7 A B C
//    2 3 5 7
// ---------------------------------------------------------

func isPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	if n%3 == 0 {
		return false
	}

	i := 5
	w := 2
	for {
		if i*i > n {
			break
		}

		if n%i == 0 {
			return false
		}

		i += w
		w = 6 - w
	}

	return true
}

func main() {
	nums := os.Args[1:]
	for _, current := range nums {
		currentInt, err := strconv.Atoi(current)
		if err != nil {
			//fmt.Printf("%d is not a valid number.\n", current)
			continue
		}
		if isPrime(currentInt) {
			fmt.Printf("%d ", currentInt)
		}
	}
}
