package main

import (
	"bankProject/banking"
	"fmt"
)

func main() {
	account := accounts.NewAccount("Pedro")
	account.Deposit(10)

	fmt.Println(account)

}
