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

## f-string
fmt.Sprintf("%s -  %d", "apple", 7) => apple - 7

## bools
!value: return bool of `NOT` value

## indexing
trim string:
  "word"[1: len(word) - 1] -> will trim 1st & last character

get indexes char
  string("word"[0])


## utils
check variable type:
  import "reflect"
  println(reflect.TypeOf(x))

