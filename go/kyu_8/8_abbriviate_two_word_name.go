package main
import ("strings"
				"fmt")

func AbbrevName(name string ) (r string){
	s := strings.Split(strings.ToUpper(name), " ")
	return fmt.Sprintf("%s.%s", string(s[0][0]), string(s[1][0]))
}

// Solition 1
func Sol1(name string) string {
  i := 0
  
  for name[i] != ' ' {
    i++
  }
  
  return string(name[0] & 95) + "." + string(name[i + 1] & 95)
}

// Solution 2
func Sol2(name string) string{
  var parts []string
  for _, part := range strings.Split(name, " ") {
    parts = append(parts, strings.ToUpper(part[:1]))
  }
  return strings.Join(parts, ".")
}



func main(){
	fmt.Println(AbbrevName("Som Nek"))
}