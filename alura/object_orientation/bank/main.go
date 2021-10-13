package main

import (
	"fmt"

	"github.com/learning-golang/alura/object_orientation/bank/accounts"
	"github.com/learning-golang/alura/object_orientation/bank/customers"
)

type checkAccount interface {
	Withdraw(value float64) string
}

func PayBillet(account checkAccount, value float64) {
	account.Withdraw(value)
}

func main() {
	clienteBruno := customers.Holder{
		Name: "Bruno",
	}
	contaBruno := accounts.SavingAccount{
		Holder: clienteBruno,
	}
	contaBruno.Deposit(2000)
	fmt.Println(contaBruno.GetBalance())
	// contaBruno.Withdraw(1999)
	fmt.Println(contaBruno.GetBalance())
	PayBillet(&contaBruno, 500)
	fmt.Println(contaBruno.GetBalance())

}
