package main

import (
	"fmt"
	"strings"
)

func ReverseWords(s string) string {
	r := make([]string, len(strings.Split(s, " ")))
	for i, word := range strings.Split(s, " ") {
		var rv_word string
		for _, c := range word {
			rv_word = string(c) + rv_word
		}
		r[i] = rv_word
	}
	return strings.Join(r, " ")
}

func Sol(str string) string {
	var rev string
	var word string

	for _, i := range str {
		if i == ' ' {
			rev = rev + word + " " // Adds word and space to result
			word = ""              // Empties word variable
		} else {
			word = string(i) + word // Adds letter to temporary word variable
		}
	}

	return rev + word // reverse those words
}

func main() {
	s := "hello world!"
	fmt.Println(Sol(s))
}
