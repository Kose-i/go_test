package main

import(
    "fmt"
    "sync"
)

func mutex_1stTest() {
  fmt.Println("mutex 1st Test")
  var count int
  var lock sync.Mutex

  increment:= func() {
    lock.Lock()
    defer lock.Unlock()
    count++
    fmt.Println("Incrementing:", count)
  }
  decrement:= func() {
    lock.Lock()
    defer lock.Unlock()
    count--
    fmt.Println("Decrementing:", count)
  }
  var arithmetic sync.WaitGroup
  for i:= 0;i <= 5;i++ {
    arithmetic.Add(1)
    go func(){
      defer arithmetic.Done()
      increment()
    }()
  }
  for i:= 0;i <= 5;i++ {
    arithmetic.Add(1)
    go func(){
      defer arithmetic.Done()
      decrement()
    }()
  }
  arithmetic.Wait()
  fmt.Println("Arithmetic Complete.")
}

func main() {
  mutex_1stTest()
}
