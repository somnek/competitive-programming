package main

import (
	"strconv"
)

func fact(i int) int {
	if i == 0 {return 1}
	return i * fact(i - 1)
}

func Strong(n int) string {
	var ttl int
	ns := strconv.Itoa(n)		          // int -> string
	for _, n := range ns{						  // string [rune, rune, ...]
		c, _ := strconv.Atoi(string(n)) // rune -> int
		ttl += fact(c)
	}
	if ttl == n {return "Strong !!!!"}
	return "Not Strong !!"
}

func main() {
	data := 144
	println(Strong(data))
}

// factorial formula :
// n! = n x (n - 1)!