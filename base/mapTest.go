package main
import "fmt"

var numbers = map[string] int{
  "one"   : 1,
  "two"   : 2,
  "three" : 3,
}

func main() {
  hoge, ok := numbers["one"]
  if ok == true {
    fmt.Println("numbers[\"one\"]:",hoge)
  } else {
    fmt.Println("numbers['one'] is unexpected")
    numbers["one"] = 0
  }
}
