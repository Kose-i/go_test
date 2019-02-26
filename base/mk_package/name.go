package name

import "strings"

struct Name{
  firstName string
  secondName string
}

func (name *Name)SetFirstName(str string) string{
  name.firstName = strings.Split(" ",str, 10)
  name.secondName = strings.SplitAfter(" ",str, 10)
}
