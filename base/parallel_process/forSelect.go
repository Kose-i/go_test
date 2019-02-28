package main

import(
    "fmt"
    "time"
)


func for_select_test1() {
  fmt.Println("for_select_test1")

  done := make(chan int)

  go func(){
    time.Sleep(5*time.Second)
    done <-1
    //close(done)
  }()
  workCount := 0
  loop:
  for {
    select{
      case <-done:
        fmt.Println("case <-done")
        break loop
      default:
    }

    workCount++
    time.Sleep(1*time.Second)
  }
  fmt.Printf("Archieve %v cycles of work\n", workCount)
}

func main() {
  for_select_test1()
}
