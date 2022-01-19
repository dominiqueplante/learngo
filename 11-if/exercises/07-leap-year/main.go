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
// EXERCISE: Leap Year
//
//  Find out whether a given year is a leap year or not.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a year number
//
//  go run main.go eighties
//    "eighties" is not a valid year.
//
//  go run main.go 2018
//    2018 is not a leap year.
//
//  go run main.go 2100
//    2100 is not a leap year.
//
//  go run main.go 2019
//    2019 is not a leap year.
//
//  go run main.go 2020
//    2020 is a leap year.
//
//  go run main.go
//    2024 is a leap year.
// ---------------------------------------------------------

/*
To determine whether a year is a leap year, follow these steps:

1 If the year is evenly divisible by 4, go to step 2. Otherwise, go to step 5.
2 If the year is evenly divisible by 100, go to step 3. Otherwise, go to step 4.
3 If the year is evenly divisible by 400, go to step 4. Otherwise, go to step 5.
4 The year is a leap year (it has 366 days).
5 The year is not a leap year (it has 365 days).
*/

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
		} else {
			return true
		}

	}
	return false
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Give me a year year")
		return
	}
	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a valid number", os.Args[1])
		return
	}

	if isLeapYear(year) {
		fmt.Printf("%d is  a leap year", year)
	} else {
		fmt.Printf("%d is not a leap year", year)
	}
}

// https://docs.microsoft.com/en-us/office/troubleshoot/excel/determine-a-leap-year
// 1600 Yes
// 1700 No
// 1800 No
// 1900 No
// 2000 Yes
// 2018 No
// 2020 Yes
// 2024 Yes
// 2100 No
// 2200 No
// 2300 No
// 2400 Yes
// 2500 No
// 2600 No
