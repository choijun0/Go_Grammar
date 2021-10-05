package main

import(
  "fmt"
  "strings"
)

func main() {
  fmt.Println("Loop")
  fmt.Println("defalutFor:", defaultFor(1,2,3,4,5))
  fmt.Println("forEeach:", forEeachInGO(1,2,3,4,5,6,7,8))
}

//naked return function
func PringUppercase(name string)(upperName string){
  defer fmt.Println("I', done~") //read code after the return is finished
  upperName = strings.ToUpper(name);
  return
}

//Loop in GO
func defaultFor(numbers ...int) (sum int){ //numbers is out of Array
  sum = 0
  for i:=0; i<len(numbers); i++ {
    sum += numbers[i]
  }
  return
}

func forEeachInGO(numbers ...int) (sum int){
  sum = 0;
  for _, number := range numbers{ //first return value is index of Array
    sum += number;
  }
  return
}