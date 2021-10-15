package main

import ("fmt"
        "strings")

func multi_table(n int) string{
	ns:= ""
	for i:= 1; i < 11; i++ {
		line := fmt.Sprintf("%d * %d = %d\n", i, n, n*i)
		ns = ns + line
		}
	return ns[:len(ns) - 1]
}

// Solution 1
func sol1(n int) string{
	ns:= ""
	for i:= 1; i < 11; i++ {
		ns += fmt.Sprintf("%d * %d = %d\n", i, n, n*i)
		}
  return strings.TrimRight(ns, "\n")
}


func main() {
	println(multi_table(5))
	println(sol1(5))
}
