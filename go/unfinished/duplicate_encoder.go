package main

import (
	"fmt"
	"strings"
)

func DuplicateEncoder(word string) string {
	var res string
	for _, c := range word {
		count := strings.Count(word, string(c))

		// if strings.Count(word, string(c)) > 1 {
		// 	res += ")"
		// } else {
		// 	res += "("
		// }
	}
	return res
}

func main() {
	x := "din" // (((
	fmt.Println(DuplicateEncoder(x))
}
