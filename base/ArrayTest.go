package main

import "fmt"

func main() {
  var test_array [100]int
  fmt.Println(test_array[1])
  for i := 0;i < 100;i+=1 {
    fmt.Printf("%d ", i)
  }
  fmt.Println("")
}
