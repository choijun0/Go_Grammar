package main

import(
  "fmt"
  //"strings"
)

type person struct{
  name string
  age int
}

func main() {
  //Array
  nameArray := [2]string{"a","b"}
  fmt.Println(nameArray[0])

  //Slice
  names := []string{"a", "b", "c"} //slice(type) : 크기제한이 없는 array
  names = append(names, "d") //새로운 요소를 추가한 새로운 배열을 리턴
  fmt.Println(names)

  //map
  choi := map[string]string{"name" : "choi", "age" : "22"} 
  fmt.Println(choi)

  //struct
  junYoung := person{"jun", 22}
  changHoon := person{name : "changHoon", age : 20}
  fmt.Println(junYoung.name, changHoon)
}
