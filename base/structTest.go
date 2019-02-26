package main

import "fmt"

type A struct {
  x,y int
}
/*func (A struct)print() {
  fmt.Println("x:", A.x)
  fmt.Println("y:", A.y)
}*/

func main() {
  var structA A
  structA.x = 3
  structA.y = 5
  fmt.Println("x:", structA.x)
  //structA.print()
}
