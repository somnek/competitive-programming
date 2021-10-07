package main

//import ""

func main() {
	input := 9
	println(Summation(input))
	println(Sol1(input))
	println(Sol2(input))
}

// Sol1 (number theory)
func Sol1(n int) int {
	return n * (n + 1) / 2
}

// Sol2 (recursion)
func Sol2(n int) int {
	if n == 1 {
		return n
	}
	return n + Summation(n - 1)
}

// me
func Summation(n int) (sum int) {
	for i := 1; i < n + 1; i ++ {
		sum += i
	}
	return
}