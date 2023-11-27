package main

import "fmt"

func MoveZeros(l []int) []int {
	var res, zes []int

	for _, n := range l {
		if n == 0 {
			zes = append(zes, 0)
		} else {
			res = append(res, n)
		}
	}
	return append(res, zes...)
}

func sol1(l []int) []int {
	// utilizing default value of make()slice which is 0
	// and idx to track next position to the inserted
	res := make([]int, len(l))
	idx := 0
	for i := 0; i < len(l); i++ {
		if l[i] != 0 {
			res[idx] = l[i]
			idx++
		}
	}
	return res
}

func main() {
	l := []int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}
	fmt.Println(MoveZeros(l))
	fmt.Println(sol1(l))

}
