package main

import (
	"fmt"
	"os"
)

func main() {
	const (
		usage           = "Usage: [username] [password]"
		u1, u2          = "jack", "inanc"
		p1, p2          = "1888", "1879"
		userUnknown     = "Unknown user %q"
		invalidPassword = "Invalid password for %q"
		ok              = "Access granted for for %q"
	)
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}
	u, p := args[1], args[2]
	if u != u1 && u != u2 {
		fmt.Printf(userUnknown, u)
		return
	}

	if u == u1 && p != p1 || u == u2 && p != p2 {
		fmt.Printf(invalidPassword, u)
		return
	}

	if u == u1 && p == p1 || u == u2 && p == p2 {
		fmt.Printf(ok, u)
		return
	}
}
