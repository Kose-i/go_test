package main

import(
"fmt"
"sync"
)

func Waitgroup_1stTest() {
  fmt.Println("sync.Waitgroup 1st Test")
  var wg sync.WaitGroup
  sayHello := func() {
    defer wg.Done()
    fmt.Println("Hello")
  }
  wg.Add(1)
  go sayHello()
  wg.Wait()
}

func Waitgropu_2ndTest() {
  fmt.Println("sync.Waitgroup 2nd Test")
  var wg sync.WaitGroup
  for _, salutation := range []string{"hello","greetings","good bye"} {
    wg.Add(1)
    go func() {
      defer wg.Done()
      fmt.Println(salutation)
    }()
  }
  for _, salutation := range []string{"hello","greetings","good bye"} {
    wg.Add(1)
    go func(str string) {
      defer wg.Done()
      fmt.Println(str)
    }(salutation)
  }
  wg.Wait()
}

func Waitgropu_3rdTest() {
  fmt.Println("sync.Waitgroup 3rd Test")
  var wg sync.WaitGroup
  sayHello := func(wg *sync.WaitGroup, id int) {
    defer wg.Done()
    fmt.Println("Hello", id)
  }

  const numGreeters = 5
  wg.Add(numGreeters)
  for i:= 0;i < numGreeters;i++ {
    go sayHello(&wg, i+1)
  }
  wg.Wait()
}

func main() {
  Waitgroup_1stTest()
  Waitgropu_2ndTest()
  Waitgropu_3rdTest()
}
