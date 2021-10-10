package main

import(
  "fmt"
  //"strings"
)

type testStruct struct{
  prop1 int
}

func createStruct() *testStruct {  // ?? 
  test := testStruct{};
  fmt.Println("create함수의 test 주소 : ", &test)
  return &test;
}

func (receiver *testStruct) testMethod(arg1 int){
  //receiver는 관찰자로 바로 사용이 가능하다.
  receiver.prop1 = arg1;
  fmt.Println("&receiver :" ,&receiver)
}


func main() {
  var a = 5; //일반 변수
  //var b* 등 에스터리스크 붙여서 변수선언 불가
  var b = &a; //b became pointer for a 
  fmt.Println(b) //address of a
  fmt.Println(*b) //value of a 
  //b = 15; b의 타입은 엄밀히 말해 *int 이다 (포인터) 따라서 일반 값으로 초기화 불가능
  *b = 15;  //b는 a를 참조하는 포인터이다. *b로 a값에 접근해 변경이 가능
  fmt.Println(a)
  //포인터 활용 함수
  add(&a, 15)
  fmt.Println(a);
  
  //Struct 와 pointer
  testinstance := createStruct();
  fmt.Println("testinstance의 주소 : ", &testinstance)
  testinstance.testMethod(12);
  fmt.Println(testinstance);

  testinstance2 := &testinstance;
  (*testinstance2).testMethod(15);
  fmt.Println(testinstance);
}

func add(a *int, addition int) {
  *a  = *a + addition;
}

/*
0xc00018c000
5
15
30
create함수의 test 주소 :  &{0}     //
testinstance의 주소 :  0xc000182020   
&receiver : 0xc000182028
&{12}
&receiver : 0xc000182030
&{15}
*/
