package main

import(
  "errors"
  "fmt"
  //"strings"
)

type Person struct{
  name string
}

func (a Person) SetName(name string) error {
  if len(name) > 5 {
    return errors.New("To long xx");
  }
  return nil;
}


//nil is the zero value of pointers and interfaces, uninitialized pointers and interfaces will be nil.
func main() {
  choi := Person{};
  err := choi.SetName("choijunyoung")
  if err != nil { 
    fmt.Println(err)
  }
}