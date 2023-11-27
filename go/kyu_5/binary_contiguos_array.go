package main

import (
	"fmt"
)

func BinArray(l []int) int {
	var maxL, ttl int
	hash := map[int]int{0: -1} // count: idx

	for i := 0; i < len(l); i++ {
		if l[i] == 0 {
			ttl--
		} else {
			ttl++
		}
		_, ok := hash[ttl]
		if !ok {
			hash[ttl] = i
		} else {
			if i-hash[ttl] > maxL {
				maxL = i - hash[ttl]
			}
		}
	}
	return maxL
}

func main() {
	a := []int{1, 1, 0, 1, 1, 0, 1, 1}
	b := []int{1, 0}
	fmt.Println(BinArray(a)) // 4
	fmt.Println(BinArray(b)) // 2
}
