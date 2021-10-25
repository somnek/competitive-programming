package main

import (
	"regexp"
	"strings"
)

// Sol1 : regex ------------------------------
func Sol1(s string) (c int) {
	r := regexp.MustCompile("[aeiou]")
	vowels := r.FindAllString(s, -1)
	return len(vowels)
}

// Sol2 : case ------------------------------
func Sol2(s string) (c int) {
	for i := range s {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u' : c ++
		}
	}
	return 
}

// me ------------------------------
func GetCount(s string) (c int) {
	vowels := "aeiou"
	for _, r := range s {
		if strings.ContainsRune(vowels, r) {
			c ++
		}
	}
  return
}

func main() {
	x := "abracadabra"
	println(GetCount(x))
	println(Sol1(x))
	println(Sol2(x))
}
