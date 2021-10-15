package main

func even_or_odd(n int) string {
	if n%2 == 0 {
		return "Even"
	}
	return "Odd"
}

// sol 1
func sol1(n int) string {
	// using bitwise & operator
	return []string{"Even", "Odd"}[n&1]
}

// input
func main() {
	println(even_or_odd(9))
	println(sol1(9))
}
