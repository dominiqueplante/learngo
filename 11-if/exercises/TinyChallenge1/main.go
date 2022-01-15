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
)

// ---------------------------------------------------------
// Tiny Challenge - Validate a single user
// EXAMPLE USERS
//  username: jack
//  password: 1888
//

//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 1888
//    Access granted to "jack".
//
//  go run main.go jack 1879
//    Invalid password for "jack".
//
// ---------------------------------------------------------

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for %q.\n"
	errPwd   = "Invalid password for %q.\n"
	accessOK = "Access granted to %q.\n"
	user     = "jack"
	pass     = "1888"
)

func main() {
	args := os.Args // create an args slice

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2] // short declaration, u = passed username, p = passed password

	if u != user {
		fmt.Printf(errUser, u)
	} else if u == user && p == pass {
		fmt.Printf(accessOK, u)
	} else {
		fmt.Printf(errPwd, u)
	}
}
