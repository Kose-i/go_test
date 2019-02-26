package main

import "fmt"


const(
  one = iota
  two
  three
)
func main() {
  fmt.Println("one:", one)
  fmt.Println("two:", two)
  fmt.Println("three:", three)
}
