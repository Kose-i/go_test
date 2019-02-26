package main

import "fmt"

func recoverly() {
    if x:= recover();x!= nil {
      fmt.Println("error:", x," happend.")
      fmt.Println("panic is recovered")
    }
}
func main() {
  defer recoverly()

  panic("panic_func1()")
}
