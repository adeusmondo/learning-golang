package accounts

import (
	"strconv"

	"github.com/learning-golang/alura/object_orientation/bank/customers"
)

type SavingAccount struct {
	Holder                                 customers.Holder
	AgencyNumber, AccountNumber, Operation uint64
	balance                                float64
}

func (s *SavingAccount) Withdraw(value float64) string {
	if value < 0 {
		return "Invalid value, only positive values!"
	}

	if value <= s.balance {
		s.balance -= value
		return "Success withdraw!"
	} else {
		return ("Cannot withdraw this ammount.. " + strconv.FormatFloat(value, 'f', 2, 64))
	}
}

func (s *SavingAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		s.balance += value
		return "Deposito realizado com sucesso!", s.balance
	}

	return "O valor do depósito deve ser maior que 0!", s.balance
}

func (s SavingAccount) GetBalance() float64 {
	return s.balance
}
