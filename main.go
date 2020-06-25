package main

import (
  "fmt"
)

type Account struct {
  owner string
  balance int
}

func NewAccount(owner string) *Account{
    account := Account {owner:owner, balance: 0}
    return &account
}
func main(){
  account := NewAccount("Hamin")
  fmt.Println(account)
  }