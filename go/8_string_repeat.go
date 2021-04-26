package main

import "strings"

func RepeatStr(repetitions int, value string) string {
	var ns string

	for i := 0; i < repetitions; i++ {
		ns += value
	}
	return ns
}

// Solution 1
func sol1(r int, v string) string {
	return strings.Repeat(v, r)
}

// Solution 2
func sol2(r int, v string) (result string) {
	for ; r > 0; r-- {
		result += v
	}
	return result
}

func main() {
	println(RepeatStr(2, "|"))
	println(sol1(2, "|"))
	println(sol1(2, "|"))
}
