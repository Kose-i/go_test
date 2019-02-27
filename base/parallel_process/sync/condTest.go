package main

import(
    "fmt"
    "time"
    "sync"
)

func cond_1stTest() {
  fmt.Println("cond 1st Test")
  c := sync.NewCond(&sync.Mutex{})
  queue := make([]interface{}, 0, 10)

  removeFromQueue := func(delay time.Duration) {
    time.Sleep(delay)
    c.L.Lock()
    queue = queue[1:]
    fmt.Println("Removed from queue")
    c.L.Unlock()
    c.Signal()
  }

  for i := 0;i < 10;i++ {
    c.L.Lock()
    for len(queue) == 2 {
      c.Wait()
    }
    fmt.Println("Adding to queue")
    queue = append(queue, struct{}{})
    go removeFromQueue(1*time.Second)
    c.L.Unlock()
  }
}

func cond_2ndTest() {
  fmt.Println("cond 2nd Test")
  type Button struct{
    Clicked *sync.Cond
  }
  button := Button {Clicked: sync.NewCond(&sync.Mutex{})}

  subscribe := func(c *sync.Cond, fn func()) {
    var goroutineRunning sync.WaitGroup
    goroutineRunning.Add(1)
    go func() {
      goroutineRunning.Done()
      c.L.Lock()
      defer c.L.Unlock()
      c.Wait()
      fn()
    }()
    goroutineRunning.Wait()
  }

  var clickRegistered sync.WaitGroup
  clickRegistered.Add(3)
  subscribe(button.Clicked, func() {
    fmt.Println("Mouse Clicked.")
    clickRegistered.Done()
  })
  button.Clicked.Broadcast()
  clickRegistered.Wait()
}

func main() {
//  cond_1stTest()
  cond_2ndTest()
}
