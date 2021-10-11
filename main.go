package main

import(
  "errors"
  "fmt"
  //"strings"
)

type Dict map[string]string; //GO는 타입을 직접만들 수 있음
//여기서 타입은 타입이름 뒤에오는 것에 별명을 의미한다고 볼수있음 
//type Money int // 코드에서 Money(1) = 1 int 와 같은 의미를 가짐

var (
  errorNotFound = errors.New("Not found");
  errorBookAlreadyExists = errors.New("book already exists");
  errorUnexpected = errors.New("unexpected error")
)

//타입은 method를 가질 수 있음
//+dictionary는 자동 참조임 따라서 Dict이라고 적어도 *Dict 으로 사용됨

//**Method

//Search
func (a Dict) findBook(bookName string) (string, error){
  author, exists := a[bookName]; //map interfaces는 인덱스 searching시 2가지 를 return함 value, isExist 
  if exists {
    return author, nil;
  }
  return "", errorNotFound;
}
//Add
func (a Dict) addBook(bookName, author string) (error){ 
  _, err := a.findBook(bookName)
  if err == errorNotFound {
    a[bookName] = author;
    return nil
  } else if err==errorBookAlreadyExists{
    return errorBookAlreadyExists; 
  } else {
    return errorUnexpected;
  }
}
//Update
func (d Dict) UpdateAuthor(bookName, newAuthor string) error {
  _, err := d.findBook(bookName)
  if err == nil {
    d[bookName] = newAuthor;
    return nil;
  } else if err == errorNotFound {
    return err;
  } else {
    return errorUnexpected;
  }
}
//Delete 
func (d Dict) DeleteBook(bookName string) error{
  _, err := d.findBook(bookName)
  if err == nil {
    delete(d, bookName)
    return nil;
  } else if err == errorNotFound {
    return err;
  } else {
    return errorUnexpected;
  }
}



func main() {
  const myBookName = "beatCoin"
  const myAuthor1 = "jun";
  const myAuthor2 = "choi";
  //Dictionary
  bookStore := Dict{};
  //Add
  err := bookStore.addBook(myBookName, myAuthor1)
  if err != nil {
    fmt.Println(err)
  } else{
    fmt.Println("after add book : ", bookStore);
  }
  //update 
  err = bookStore.UpdateAuthor(myBookName, myAuthor2);
  if err != nil {
    fmt.Println(err)
  } else{
    fmt.Println("after update book : ", bookStore);
  }
  //delete
  err = bookStore.DeleteBook(myBookName);
  if err != nil {
    fmt.Println(err)
  } else{
    fmt.Println("after delete book : ", bookStore);
  }
  /*
  after add book :  map[beatCoin:jun]
  after update book :  map[beatCoin:choi]
  after delete book :  map[]
  */
}