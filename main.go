package main

import(
  "fmt"
  //"strings"
)

func main() {
  a := 2
  b := &a 
  *b = 5 //애스터리스크를 사용하면 해당 변수의 값을 주소로하는 변수에 접근 하는 개념 
  //즉 특정변수의 메모리 주소를 가진 변수가 해당 변수의 포인터이다.
  fmt.Println(a);
  
}
