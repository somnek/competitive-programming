package main

import (
	"fmt"
	"regexp"
)

func alphanumeric(s string) bool {
	var valid = regexp.MustCompile("^[0-9a-zA-Z]+$")
	return valid.MatchString(s)
}

func main() {
	x := "Helloworld"
	y := "gigachad_"
	z := "  "
	fmt.Println(alphanumeric(x))
	fmt.Println(alphanumeric(y))
	fmt.Println(alphanumeric(z))
}
/*
func main() {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

	fmt.Println(validID.MatchString("adam[23]"))
	fmt.Println(validID.MatchString("eve[7]"))
	fmt.Println(validID.MatchString("Job[48]"))
	fmt.Println(validID.MatchString("snakey"))
}
*/