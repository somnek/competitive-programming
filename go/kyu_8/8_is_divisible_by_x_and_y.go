package main

func IsDivisible(n, x, y int) bool {
	return n % x == 0 && n % y == 0
	// sol: return n%x + n%y == 0
}
func main() {
	println(IsDivisible(8, 2, 1))
}