package main

import "fmt"

type Object1 struct {
  x,y int
}

func (t *Object1) Add() int {
  t.x += t.y
  return t.x
}

func (t *Object1) Sub() int {
  t.y -= t.x
  return t.y
}

type Object2 struct {
  *Object1
  z int
}

func (t *Object2) print() {
  fmt.Println("Object1.x:", t.x)
  fmt.Println("Object1.y:", t.y)
  fmt.Println("Object1.z:", t.z)
}

func main() {
  obj := &Object2{new(Object1), 0}
  obj.x = 1
  obj.y = 3
  obj.z = 100
  fmt.Println("obj.Add()", obj.Add())
  fmt.Println("obj.Sub()", obj.Sub())
  obj.print()
  return
}
