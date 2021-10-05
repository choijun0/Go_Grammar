package main

import(
  "fmt"
  //"strings"
)

func main() {
  fmt.Println(checkAdult(16), checkAdultSwitch(16));
}

//If else
func checkAdult(age int) (isOk bool){
  if koreanAge := age + 2; koreanAge < 18{
    return false;
  } else if koreanAge < 19{ //koreanAge는 사실상 if문 밖에서 선언한 것과 같은 효과를 낸다.
    fmt.Println("that's close bro hahaha")
    return true
  } else {
    return true
  }
}

func checkAdultSwitch(age int) bool {
  switch koreanAge := age + 2; koreanAge{
    case 18:
      return false
  }
  return true 
}