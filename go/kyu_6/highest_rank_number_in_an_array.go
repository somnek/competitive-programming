package main

import (
	"fmt"
)

func HighestRank(l []int) int {
	m, mk, mv := map[int]int{}, 0, 0 // hash, max key, max count
	for _, x := range l {
		m[x]++

		if m[x] > mv {
			mv = m[x]
			mk = x
		} else if m[x] == mv && x > mk {
			mk = x
		}
	}
	return mk
}

func Sol(l []int) int {
	mii, maxK, maxV := map[int]int{}, 0, 0
	for _, v := range l {
		mii[v]++
		if mii[v] > maxV || (mii[v] == maxV && v > maxK) {
			maxK = v
			maxV = mii[v]
		}
	}
}

func main() {
	l1 := []int{12, 10, 8, 12, 7, 6, 4, 10, 10}     // 10
	l2 := []int{12, 10, 8, 12, 7, 6, 4, 10, 12, 10} // 12
	l3 := []int{2}                                  // 2
	fmt.Println(HighestRank(l1))
	fmt.Println(HighestRank(l2))
	fmt.Println(HighestRank(l3))
}

// var pl = []int{}
// for k, v := range m {
// 	if v == maxCount {
// 		pl = append(pl, k)
// 	}
// }

// finalMax := 0
// for _, n := range pl {
// 	if n > finalMax {
// 		finalMax = n
// 	}
// }
// return finalMax
