package main

import "fmt"

func cakes(need, have map[string]int) int {
	var total int
	bake(need, have, &total)
	return total
}

func bake(need, have map[string]int, total *int) {
	for k, v := range need {
		if have[k] < v {
			return
		} else {
			have[k] -= v
		}
	}
	*total += 1
	bake(need, have, total)

}

// sol
func sol_2(recipe, available map[string]int) int {
	for k, v := range recipe {
		available[k] -= v
		if 0 > available[k] {
			return 0
		}
	}
	return 1 + sol_2(recipe, available)
}

func main() {
	need := map[string]int{"cream": 200, "flour": 300, "milk": 100, "oil": 100, "sugar": 150}
	have := map[string]int{"cream": 5000, "flour": 20000, "milk": 20000, "oil": 30000, "sugar": 1700}
	fmt.Println(cakes(need, have))
}
