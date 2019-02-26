package main

import(
    "fmt"
    "time"
)

func select_test1(x chan int, end chan int) {
  var t [2]int
  t[0] = 1
  t[1] = 1
  for i := 0;i < 50;i++ {
    tmp := t[0] + t[1]
    x <- tmp
    t[i%2] = tmp
    time.Sleep(time.Millisecond*30)
  }
  end <-1
}

func select_test2(x chan int) {
  for {
    var hoge int
    hoge =<-x
    fmt.Println("x:",hoge)
  }
}

func main() {
  x := make(chan int)
  end := make(chan int)
  go select_test1(x, end)
  go select_test2(x)
  <-end
}
