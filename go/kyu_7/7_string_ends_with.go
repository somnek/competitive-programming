package main

import "fmt"

func solution(str, ending string) bool {
	lastTwo := str[len(str)-2:]
	if lastTwo == ending {
		return true
	}
	return false
}

func main() {
	str := "chew"
	end := "wew"
	fmt.Println(solution(str, end))
}
