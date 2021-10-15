package main

import "sort"

func main() {
	input := []int{5, 11, 2, 4, 5, 6}
	println(Smallest(input))
	println(Sol1(input))
}

// sol1: sort
func Sol1(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}

// me
func Smallest(numbers []int) int {
	var sml int
		for i, n := range(numbers){
			if i == 0 || n < sml {sml = n}
		}
	return sml
}