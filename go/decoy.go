package main

//import ("strings")

func SumOfPositive(numbers []int) (sum int) {
	for _, n := range numbers {
		if n >= 0 {sum += n}
	}
	return sum
}

func main() {
	println()
}
