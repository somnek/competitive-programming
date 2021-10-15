package main

import (
	"fmt"
)

func Map(x []int) (doubled []int) {
	for _, n := range x {
		doubled = append(doubled, n * 2)
	}
	return
}

func main() {
	input := []int{4, 2, 3}
	fmt.Printf("%v", Map(input))
}