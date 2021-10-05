package main

import(
  "fmt"
  "strings"
)

func main() {
  const constName string = "choi";
  //constName = "jun" -> ERR
  var flexibleName string = "choi";
  flexibleName = "choi2";
  fmt.Println("constName : "+constName,"/ flexibleName : "+flexibleName);

  sayName("JunYoung");

  fmt.Println(lengthOfName("Jun"));

  //return couple of values 
  var val1 int
  var val2 string
  var val3 string
  //다중 return 값 전달
  val1, val2 = multipleReturn("choi")
  //받지않을 리턴값은 _ 로표기
  _, val3 =  multipleReturn("choi")
  //shortcut 이렇게 적으면 할당값에 따라 변수의 type을 자동으로 지정
  nameLen, upperName := multipleReturn("choi")
  fmt.Println(nameLen, upperName);

  //shortcut으로 선언된 변수는 기본적으로 var형
  nameLen = 5;
  fmt.Println(nameLen);

  fmt.Println(val1, val2);
  fmt.Println(val3);
}

func sayName(name string) { //return nothing
  fmt.Println(name);
}

func lengthOfName(name string) int {
  return len(name);
}

func multipleReturn(name string) (int, string){
  return len(name), strings.ToUpper(name);
}

func infiniteArg(words ...string) {
  fmt.Println(words);
}