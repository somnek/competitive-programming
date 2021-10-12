package main

func Grow(arr []int) int {
	sum := 1
	for _, n := range arr {
		sum *= n
	}
	return sum
}

func main() {
	input := []int{1, 2, 4, 5}
	println(Grow(input))
}