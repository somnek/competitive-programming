package main

import (
	"fmt"
)

func Permutations(s string) []string {
	var result []string
	PermutationsHelper(s, "", &result)
	return result
}

func PermutationsHelper(s string, prefix string, result *[]string) {
	if len(s) == 0 {
		*result = append(*result, prefix)
	} else {
		for i := 0; i < len(s); i++ {
			PermutationsHelper(s[:i]+s[i+1:], prefix+string(s[i]), result)
		}
	}
}

func main() {
	s := "abcd"
	fmt.Println(Permutations(s))
}
