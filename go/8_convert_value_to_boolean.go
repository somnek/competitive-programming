// return "No" or "Yes" for bool
package main

func BoolToWord(word bool) string {
	r := "Yes"
	if word == false {
		r = "No"
	}
	return r
}

// Solution 1
func sol1(word bool) string {
	return map[bool]string{false: "No", true: "Yes"}[word]
}

// Solution 2
func sol2(word bool) string {
	if !word {
		return "No"
	}
	return "Yes"
}

func main() {
	println(BoolToWord(true))
	println(sol1(true))
	println(sol2(true))
}
