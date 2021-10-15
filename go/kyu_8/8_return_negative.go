package main

func main(){
	println(MakeNegative(5))
	println(MakeNegative(0))
	println(MakeNegative(-5))
}

func MakeNegative(x int) int {
  if x > 0 {
    return -x
  }
  return x
  }
