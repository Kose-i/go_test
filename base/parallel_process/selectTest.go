package main

import(
    "fmt"
    "time"
)

func select_test1(x chan int) {
  var hoge int
  for i := 0;i < 1;i++ {
    time.Sleep(time.Second)
  }
  x <- hoge
  fmt.Println("select_1-x:",hoge)
}
func select_test2(x chan int) {
  var hoge int
  for i := 0;i < 100;i++ {
    time.Sleep(time.Second)
  }
  x <- hoge
  fmt.Println("select_2-x:", hoge)
}
func main() {
  fmt.Println("select start")
  c1 := make(chan int)
  c2 := make(chan int)
  go select_test1(c1)
  go select_test2(c2)
  time.Sleep(time.Second*2)
  select {
    case _ =<- c1:
      fmt.Println("select_test1")
    case _ =<- c2:
      fmt.Println("select_test2")
    default:
      fmt.Println("default")
  }
}
