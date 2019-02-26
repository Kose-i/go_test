package main

import "fmt"

func test1(x int) (y int, ok bool) {
  if x < 0{
    return -x, false
  } else {
    return x, true
  }
}

func test2(x int) (y int, z string, err bool) {
  z = "test2"
  if x > 0 {
    y = x
  } else {
    y = -x
  }
  return y, z,true
}

func function_12() {
  test1_int, test1_ok := test1(2)
  test2_int, test2_str, test2_err := test2(3)
  if test1_ok == true {
    fmt.Println(test1_int)
  } else {
    fmt.Println("test1_ok == false")
  }
  if test2_err == false {
    fmt.Println("test2_int:", test2_int, " test2_str:", test2_str)
  } else {
    fmt.Println("test2_err == true")
  }
}
func closure_test() {
  var func_closure = func(str string)int{fmt.Println(str);return 3}
  fmt.Println(func_closure("closure_test"))
}
func main() {
  function_12()
  closure_test()
}
