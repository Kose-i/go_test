package foo

//go:generate echo Hello, go generate!

func Foo() {
  return nil
}
/*
$go generate filename.go
*/
