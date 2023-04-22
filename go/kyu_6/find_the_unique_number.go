package main

import "fmt"

func solve(arr []float32) float32 {
	m := make(map[float32]int)
	for _, v := range arr {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

func main() {
	data := []float32{1, 1, 1, 2, 1, 1}
	fmt.Println(solve(data))
}
