package main

import "fmt"

func ValidBraces(str string) bool {
	// l := len(str)
	for i, c := range str {
		if str[c] {

		}
	}
	return true
}

func main() {
	// x := "([{}])"
	y := "(){}[]"
	fmt.Println(ValidBraces(y))
}

// "(){}[]"   =>  True
// "([{}])"   =>  True
// "(}"       =>  False
// "[(])"     =>  False
// "[({})](]" =>  False
