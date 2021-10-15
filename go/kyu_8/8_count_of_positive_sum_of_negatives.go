package main

import "fmt"

// Sol 1
func Sol1(x []int) []int {
	res := []int{0, 0}
	for _, v := range(x) {
		if v < 0 {
			res[1] += v
		} else {
			res[0] ++
		}
	}
	return res
}


func CountPositivesSumNegatives(x []int) (res []int) {
	var p int
	var n int
	for _, e := range x {
		if e > 0 {
			p += 1
		}else if e < 0 {
			n += e
		}
	}
	res = append(res, p)
	res = append(res, n)
	return
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	fmt.Printf("%v", CountPositivesSumNegatives(input))
	fmt.Printf("%v", Sol1(input))
}