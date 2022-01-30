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
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

const (
	maxTurns = 5
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, harder it gets.

Wanna play?

(Provide -v flag as first argument to see the picked numbers.)

`
)

var (
	verbose = false
	err     error
	guess   = 0
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) == 2 {
		if args[0] != "-v" {
			fmt.Printf(usage, maxTurns)
			return
		} else {
			verbose = true
			guess, err = strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Not a number.")
				return
			}
		}
	} else if len(args) == 1 {
		guess, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Not a number.")
			return
		}

	}

	if guess <= 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess) + 1

		if verbose {
			fmt.Printf("%d ", n)
		}

		if n != guess {
			continue
		}

		fmt.Printf("🎉  YOU WIN!")

		return
	}
	fmt.Println("You lose")
}
