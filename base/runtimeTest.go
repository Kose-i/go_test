package main
import(
    "fmt"
    "runtime"
)

func runtime_1st_test() {
  fmt.Println("Num CPU is ", runtime.NumCPU())
}

func main() {
  runtime_1st_test()
}
