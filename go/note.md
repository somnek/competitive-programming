## condition
&&: and
||: or

- if else:
    > if cond1{
    > }else{ // else must be after '}' of if
    > }



## variables
var b, c int = 1, 2
var d = true
var e int 
f := "apple"


## map
map[KeyType]ValueType{KEY, VALUE}[input]

## conversion
import ("strconv")
strconv.Itoa(n)

## print
- f string
  fmt.Sprintf("%s -  %d", "apple", 7) => apple - 7
- print array
  fmt.Printf("%v", num_arr)

## bools
!value: return bool of `NOT` value

## arrays // slice ( ?? )
- append to list:
  arr = append(arr, e)

## indexing
trim string:
  "word"[1: len(word) - 1] -> will trim 1st & last character

get indexes char
  string("word"[0])

## string function
split to list
  strings.Split("word", " ")

upper case:
  strings.ToUpper()

`OR` fast hacky way:
  string(word[0] & 95)

## loops
- loop through string (will stop at ' ')
------------------------------
  for text[i] != ' '{
    i ++
  }

- loop & print string
------------------------------
for _, c := range 'word' {
  fmt.Printf("%c", c)
}

## increment (extras)
- multiple by 2: <<
- divide   by 2: >>


## utils
check variable type:
  import "reflect"
  println(reflect.TypeOf(x))


## converstion
- rune to string 
  string(c) // c is type of rune, from range loop: 
    ~ c, _ := strconv.Atoi(string(n)) // rune -> int
- int to string:
    ~ strconv.Itoa(n)





``
## qaveat
- for function with multiple arguments:
  if all the arguments has the same type
  only have to specify once:
  eg:
  >  func SumOfPositive(numbers []int) (sum int) {
  >    for _, n := range numbers {
  >      if n >= 0 {sum += n}
  >    }
  >    return sum
  >  }




# refs
- *Slice Tricks*: https://github.com/golang/go/wiki/SliceTricks