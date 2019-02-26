package main

import "fmt"
import "time"

func parallel_test1(x chan int, end chan int) {
  var tmp int
  for i:= 0;i < 3;i++ {
    tmp++
    x<-tmp
    time.Sleep(time.Second)
  }
  fmt.Println("3 Seconds has passed")
  end <-1
  return
}

func parallel_test2(x chan int) {
  var y int
  for {
    y=<-x
    fmt.Println("x:", y)
  }
}

func main() {
  x := make(chan int)
  end := make(chan int)
  go parallel_test1(x, end)
  go parallel_test2(x)
  <-end
  fmt.Println("finish main")
}
