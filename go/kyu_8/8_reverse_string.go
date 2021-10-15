package main

func main() {
	println(rs("something"))
}

func rs(word string) string {
	for _, c := range word {
		ns = string(c) + ns
	}
	return ns
}