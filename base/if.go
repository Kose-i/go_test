package main

import "fmt"

func main() {
  a := 3
  if a == 3 {
    fmt.Printf("a == 3\n")
  } else if a == 2 {
    fmt.Printf("a == 2\n")
  } else {
    fmt.Printf("else\n")
  }
  fmt.Printf("Hello\n")
}
