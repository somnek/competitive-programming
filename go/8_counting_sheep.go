package main

func main() {
	input := []bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true}
	println(CountSheep(input))
	println(Sol1(input))
}

// Sol 1
func Sol1(numbers []bool) (count int) {
	for _, c := range numbers {
		if c {count++}
	}
	return
}

// me
func CountSheep(numbers []bool) int {
	var count int
	for _, c := range numbers {
		if c == true {count += 1}
	}
	return count

}
