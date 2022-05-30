package main

import (
	"fmt"
)

// 0^1 = 1
// 1^0 = 1
// 0^0 = 0
// 1^1 = 0

func sol1(seq []int) int {
	// solve by cancelling out pairs
	res := 0
	for _, x := range seq {
		res ^= x
	}
	return res
}

func findOdd(seq []int) int {
	m := make(map[int]int)
	for _, v := range seq {
		m[v]++
	}
	for k, v := range m {
		if v%2 != 0 {
			return k
		}
	}
	return 1
}

func main() {
	x := []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}
	fmt.Println(findOdd(x))
	fmt.Println(sol1(x))
}
