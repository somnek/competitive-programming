package main

import (
	"fmt"
	"strings"
)

func DNAStrand(dna string) (s string) {
	var complements = map[string]string{
		"A": "T",
		"T": "A",
		"G": "C",
		"C": "G",
	}
	for _, char := range dna {
		s += complements[string(char)]
	}
	return
}

//  ------< solve 1 ------
var dnaReplacer *strings.Replacer = strings.NewReplacer(
	"A", "T",
	"T", "A",
	"C", "G",
	"G", "C",
)

func Sol1A(dna string) string {
	return dnaReplacer.Replace(dna)
}

//  ------ solve 1 >------

//  ------< solve 2 ------
func Sol2(dna string) string {
	conversion := func(r rune) rune {
		switch r {
		case 'A':
			return 'T'
		case 'T':
			return 'A'
		case 'G':
			return 'C'
		case 'C':
			return 'G'
		}
		return 0
	}

	return strings.Map(conversion, dna)
}

//  ------ solve 2 >------

func main() {
	if DNAStrand("AAAA") != "TTTT" {
		fmt.Println("Exp: TTTT, Got: ", DNAStrand("AAAA"))
	} else {
		fmt.Println("🎊 correct!")
	}
}
