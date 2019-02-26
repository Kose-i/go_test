package main

import "fmt"

type A struct{
  x,y int
  str string
}

func (a A)Add()int {
  fmt.Println("A Add", a.x, "+",a.y)
  return a.x+a.y
}

type B struct{
  x,y int
}
func (b B)Add()int {
  fmt.Println("B Add", b.x, "+",b.y)
  return b.x+b.y
}

type NumAdd interface{
  Add() int
}

func test_Add(a NumAdd, b NumAdd) int {
  fmt.Println("test_Add()" ,a.Add(), "+", b.Add())
  return a.Add() + b.Add()
}

func main() {
  var a A
  var b B
  a.x = 2
  a.y = 3
  b.x = 4
  b.y = 2
  fmt.Println(test_Add(a,b))
}
