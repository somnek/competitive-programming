package main

import (
	"fmt"
)

func ReverseWords(s string) string {
	var r string
	for _, c := range s {
		r = string(c) + r
	}
	return r

}

func main() {
	s := "hello world!"
	fmt.Println(ReverseWords(s))
}
