// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Rename imports
//
//  1- Import fmt package three times with different names
//
//  2- Print a few messages using those imports
//
// EXPECTED OUTPUT
//  hello
//  hey
//  hi
// ---------------------------------------------------------

// https://scene-si.org/2018/01/25/go-tips-and-tricks-almost-everything-about-imports/

// ?

import "fmt"
import f "fmt"
import f2 "fmt"
// ?
// ?

func main() {
	//prt1 := fmt.Println
	//prt2 := fmt.Println
	prt3 := fmt.Println
	// ?
	f.Println("hello")
	f2.Println("hey")
	prt3("i")
	// ?
	// ?
}
