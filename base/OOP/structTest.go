package main

import "fmt"

type A struct {
  x,y int
}

func (a A)print() {
  fmt.Println("x:", a.x)
  fmt.Println("y:", a.y)
  a.x += a.y
  fmt.Println(a.x)//ポインタ経由ではないので値は変わらない
}
func (a *A)print_pointa() {
  fmt.Println("x:", a.x)
  fmt.Println("y:", a.y)
  a.x += a.y
  fmt.Println(a.x)//ポインタ経由なので値は変わる
}

func main() {
  var structA A
  structA.x = 3
  structA.y = 5
  fmt.Println("x:", structA.x)
  structA.print()
  structA.print()
  structA.print_pointa()
  structA.print_pointa()
}
