package main

import(
  "fmt"
  //"strings"
)

type Account struct{ //small scare name of property is private property
  name string
  money int
}

func CreateAccount(name string) *Account{
  account := Account{name : name, money : 0}
  return &account;
}

//automatically be method of Account because of struct receiver
func (receiver *Account) Deposit(amount int){
  receiver.money = amount;
}

func main() {

  //Method
  accountChoi := CreateAccount("choi");
  fmt.Println(accountChoi);

  //use Deposit as method !!
  accountChoi.Deposit(1000);
  fmt.Println(accountChoi);
}
