package accounts

import (
	"strconv"

	"github.com/learning-golang/alura/object_orientation/bank/customers"
)

type CheckingAccount struct {
	Holder                      customers.Holder
	AgencyNumber, AccountNumber uint64
	balance                     float64
}

func (c *CheckingAccount) Withdraw(value float64) string {
	if value < 0 {
		return "Invalid value, only positive values!"
	}

	if value <= c.balance {
		c.balance -= value
		return "Success withdraw!"
	} else {
		return ("Cannot withdraw this ammount.. " + strconv.FormatFloat(value, 'f', 2, 64))
	}
}

func (c *CheckingAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		c.balance += value
		return "Deposito realizado com sucesso!", c.balance
	}

	return "O valor do depósito deve ser maior que 0!", c.balance
}

func (c *CheckingAccount) Transfer(value float64, destinationAccount *CheckingAccount) bool {
	if c.balance < value && value < 0 {
		return false
	}

	c.balance -= value
	destinationAccount.balance += value
	return true
}

func (c CheckingAccount) GetBalance() float64 {
	return c.balance
}
