package main

import (
	"fmt"
	"strings"
)

func Meeting(s string) string {
	var n struct {
		fn string
		ln string
	}
	// var r string
	l := []string{}
	for _, v := range strings.Split(s, ";") {
		n.fn = strings.Split(v, ":")[0]
		n.fn = strings.Split(v, ":")[1]
		l = append(l, fmt.Sprintf("(%s,%s)", n.fn, n.ln))
	}
	fmt.Println(l)
	return ""
}

func main() {
	x := "Fred:Corwill;Wilfred:Corwill;Barney:Tornbull;Betty:Tornbull;Bjon:Tornbull;Raphael:Corwill;Alfred:Corwill"
	Meeting(x)
}

// "(CORWILL, ALFRED)(CORWILL, FRED)(CORWILL, RAPHAEL)(CORWILL, WILFRED)(TORNBULL, BARNEY)(TORNBULL, BETTY)(TORNBULL, BJON)"
