package main

import(
  "fmt"
  //"strings"
)

type Person struct{
  name string
}

func (a Person) String() string{
  return "이게 호출된다!!"
}


//nil is the zero value of pointers and interfaces, uninitialized pointers and interfaces will be nil.
func main() {
  choi := Person{};
  fmt.Println(choi);
}