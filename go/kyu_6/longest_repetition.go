package main

import (
	"fmt"
)

type Result struct {
	a rune
	i int
}

func LongestRepetition(s string) Result {
	var mx int
	var a rune
	c := 1

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			c++
		} else {
			c = 1
		}

		if c > mx {
			mx = c
			a = rune(s[i])
		}
	}
	return Result{a, mx}
}

func Sol(s string) (res Result) {
	var a rune
	var c int
	for _, x := range s {
		if a != x {
			a, c = x, 0
		}
		c++
		if c > res.i {
			res.a = x
			res.i++
		}
	}
	return res
}

func main() {
	tests := []string{
		"aaaabb",      // ('a', 4)
		"bbbaaabaaaa", // ('a', 4)
		"cbdeuuu900",  // ('u', 3)
		"abbbbb",      // ('b', 5)
		"aabb",        // ('a', 2)
		"ba",          // ('b', 1)
		"",            // ('', 0)
	}

	for _, t := range tests {
		fmt.Println(LongestRepetition(t))
		fmt.Println(Sol(t))
	}
}
