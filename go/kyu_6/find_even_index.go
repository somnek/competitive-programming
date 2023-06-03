package main

import "fmt"

// ------ sol me ------
func sum(l []int) (ttl int) {
	for _, d := range l {
		ttl += d
	}
	return ttl
}

func FindEvenIndex(l []int) int {
	for i := 0; i < len(l); i++ {
		if sum(l[:i]) == sum(l[i+1:]) {
			return i
		}
	}
	return -1
}

// ------ sol 1 ------
func Sol1(arr []int) int {
	var sum, b int
	for _, n := range arr {
		sum += n
	}
	for i, n := range arr {
		sum -= n
		if sum == b {
			return i
		}
		b += n
	}
	return -1
}

func main() {
	fmt.Println(FindEvenIndex([]int{1, 2, 3, 4, 3, 2, 1}))
}
