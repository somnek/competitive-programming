package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Sol 1 (similar but using '0'+n) MAGIC
func Sol1(numbers []uint) string {
	var s string
	for _, c := range numbers {
		s += string('0' + c)
	}
	return fmt.Sprintf("(%s) %s-%s", s[0:3], s[3:6], s[6: 10])
}

// Sol 2 (using interface in dict)
func Sol2(numbers []uint) string {
	sNum := make([]interface{}, len(numbers))
	for i, n := range numbers {
		sNum[i] = n
	}
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", sNum...)
}

// Sol 3 (messy using Trim, Replace)
func Sol3(numbers []uint) string {
	s := strings.Trim(strings.Replace(fmt.Sprint(numbers), " ", "", -1), "[]")
	return fmt.Sprintf("(%s) %s-%s", s[0:3], s[3:6], s[6:10])
}

// me
func CreatePhoneNumber(numbers []uint) string {
	var s string
	for _, c := range numbers {
		s += strconv.FormatUint(uint64(c), 10)
	}
	return fmt.Sprintf("(%s) %s-%s", s[0:3], s[3:6], s[6: 10])
}

func main() {
	x := []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	println(CreatePhoneNumber(x))
	println(Sol1(x))
	println(Sol2(x))
	println(Sol3(x))
}