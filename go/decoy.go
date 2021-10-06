package main

import ("strings")


func sol1(name string) string {
	var parts [] string
	for _, part := range strings.Split(name, " ") {
		println(part[:1])
		parts = append(parts, strings.ToUpper(part[:1]))
	}
	return strings.Join(parts, ".")
}


func main() {
	println(sol1("hello you"))
}


