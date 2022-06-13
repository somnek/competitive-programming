package main

import "fmt"

func numberOfPairs(data []string) int {
	m := make(map[string]int)
	s := 0
	for _, g := range data {
		m[g]++
	}
	for _, v := range m {
		s += v / 2
	}
	fmt.Println(s)
	return s
}

func main() {
	var data = [][]string{
		{"red", "green", "blue"},
		{"gray", "black", "purple", "purple", "gray", "black"},
		{},
		{"red", "green", "blue", "blue", "red", "green", "red", "red", "red"},
	}
	for _, d := range data {
		numberOfPairs(d)
	}
}
