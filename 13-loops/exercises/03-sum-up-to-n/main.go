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
// EXERCISE: Sum up to N
//
//  1. Get two numbers from the command-line: min and max
//  2. Convert them to integers (using Atoi)
//  3. By using a loop, sum the numbers between min and max
//
// RESTRICTIONS
//  1. Be sure to handle the errors. So, if a user doesn't
//     pass enough arguments or she passes non-numerics,
//     then warn the user and exit from the program.
//
//  2. Also, check that, min < max.
//
// HINT
//  For converting the numbers, you can use `strconv.Atoi`.
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage program min max")
		return
	}
	min, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf(`Wrong min: %q`+"\n", os.Args[1])
		return
	}

	max, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Printf(`Wrong max: %q`+"\n", os.Args[1])
		return
	}

	if min > max {
		fmt.Println("Min should be less than max")
	}

	var sum = 0
	for i := min; i <= max; i++ {
		fmt.Printf("%d", i)
		if i < max {
			fmt.Print(" + ")
		}

		sum += i
	}
	fmt.Print(" = ")
	fmt.Print(sum)
}
