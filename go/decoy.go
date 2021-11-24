package main

import (
	"fmt"
)

// write a function for bubble sort
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{5, 3, 4, 1, 2}
	fmt.Println(bubbleSort(arr))
}

//func addTo(base int, vals ...int) []int {
//	out := make([]int, 0, len(vals))
//	for _, v := range vals {
//		out = append(out, base+v)
//	}
//	return out
//}
//
//func main() {
//	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))
//}
