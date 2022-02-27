package main

import (
	"fmt"
	"strings"
)

func Accum(s string) string {
	var mumble string
	for i, c := range s {
		mumble += strings.ToUpper(string(c)) + strings.Repeat(string(c), i) + "-"
	}
	return mumble[:len(mumble)-1]
}

func Sol(s string) string {
	parts := make([]string, len(s))
	for i, c := range s {
		parts[i] = strings.Title(strings.Repeat(string(c), i))
	}
	return strings.Join(parts, "-")
}

func main() {
	fmt.Println(Accum("abc"))
	fmt.Println(Accum("abc"))
}
