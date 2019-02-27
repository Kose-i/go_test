package main

import(
"fmt"
"sync"
)

func once_1st_test() {
  fmt.Println("Once 1st test")
  count := 0
  increment := func() {count++}

  var once sync.Once

  var increments sync.WaitGroup
  increments.Add(100)
  for i := 0;i < 100;i++ {
    go func() {
      defer increments.Done()
      once.Do(increment)
    }()
  }
  increments.Wait()
  fmt.Println("count is", count)
}

func once_2nd_test() {
  fmt.Println("Once 2nd test")
  count := 0
  increment := func() {count++}
  decrement := func() {count--}

  var once sync.Once
  once.Do(increment)
  once.Do(decrement)

  fmt.Println("count is", count)
}

func main() {
  once_1st_test()
  once_2nd_test()
}
