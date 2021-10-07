package main

import (
	"strings"
	"unicode"
)

func main() {
	println(NoSpace("some words here.. ."))
	println(Sol1("some words here.. ."))
}


// Sol 1
func Sol1(word string) string {
	return strings.ReplaceAll(word, " ", "")
}

// me
func NoSpace(word string) (res string) {
		for _, c:= range(word){
			if !(unicode.IsSpace(c)) {
				res += string(c)
			}
		}
		return
	}