package main

func SumOfPositive(numbers []int) (sum int) {
	for _, n := range numbers {
		if n >= 0 {sum += n}
	}
	return sum
}

func main(){
	input := []int{1, -6, 4, 5, -9}
	println(SumOfPositive(input))
}