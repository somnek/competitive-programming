package main


func SquareSum(numbers []int) (total int) {
  for i:=0; i<len(numbers); i++ {
    total += numbers[i] * numbers[i]
  }
  return total
}

// Solution 1
func Sol1(numbers []int) (total int) {
  for _, n := range numbers {
     total = total + n*n
  }
  return
}

func main() {
  println(SquareSum([]int{1, 2}))
  println(Sol1([]int{1, 2}))
}
