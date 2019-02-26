package main

import "fmt"

func main() {
  switchtest_value := 1

  switch switchtest_value {
    case 0,1:
      fmt.Println("switchtest_value:0 or 1")
      fallthrough
    case 2:
      fmt.Println("switchtest_value:2")
    case 3:
      fmt.Println("switchtest_value:3")
    default:
      fmt.Println("else 0 1 2 3 switchtest_value")
  }
}
