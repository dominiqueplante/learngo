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
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------

var (
	winString string
	loseSring string
)

const (
	maxTurns = 5
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, harder it gets.

Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess <= 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess) + 1

		if n != guess {
			continue
		}

		value := rand.Intn(2)
		switch value {
		case 0:
			winString = "You won"
			//fmt.Printf("\r/ Please Wait. Processing....")
		case 1:
			winString = "You're awesome"
			//fmt.Printf("\r| Please Wait. Processing....")
		}
		fmt.Printf("%s", winString)
		return
	}
	value := rand.Intn(2)
	switch value {
	case 0:
		loseSring = "You lost"
		//fmt.Printf("\r/ Please Wait. Processing....")
	case 1:
		loseSring = "You suck"
		//fmt.Printf("\r| Please Wait. Processing....")
	}
	fmt.Printf("%s", loseSring)
}
