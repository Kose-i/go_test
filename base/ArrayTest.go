package main

import "fmt"

func printArray(array [100]int) {
  length := len(array)
  fmt.Println("Array-size:", length)
  for i := 0;i < length;i+=1 {
    fmt.Printf("%d ", array[i])
  }
  fmt.Println("")
}

func set_three(array []int) {
  for i,_ := range(array) {
    array[i] = 3
  }
}

func main() {
  var test_array [100]int
  test_array_cut := test_array[0:50]
  fmt.Println("set_three:test_array[0:50]")
  set_three(test_array_cut)
  printArray(test_array)
}
