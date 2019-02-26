package main

import "fmt"

func main() {
  var x int

  go func () {
    x++
  }()

  if x == 0 {
    fmt.Println("x is 0.")
  } else {
    fmt.Println("x is 3")
  }
}
