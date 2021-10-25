package main

import (
	"strconv"
	"strings"
)

func FreqSeq(str string, sep string) string {
	var r string
  for _, c := range str {
		r += strconv.Itoa(strings.Count(str, string(c))) + sep
  }
  return r[0:len(r) - 1]
}

func main() {
	x1 := "hello world"
	x2 := "-"
	println(FreqSeq(x1, x2))
}