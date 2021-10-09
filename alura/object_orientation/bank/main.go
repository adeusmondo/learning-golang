package main

import "fmt"

type CheckingAccount struct {
	holder        string
	agencyNumber  int64
	accountNumber int64
	balance       float64
}

func main() {

	holder := "Carlos"
	agencyNumber := 559
	accountNumber := 1032
	balance := 352.20

	ckAccountCarlos := CheckingAccount{
		holder:        holder,
		agencyNumber:  int64(agencyNumber),
		accountNumber: int64(accountNumber),
		balance:       balance,
	}
	ckAccountYasmin := CheckingAccount{
		"Yasmin",
		159,
		951,
		15998.87,
	}

	fmt.Println("Hello", holder,
		"Agency Number:", agencyNumber,
		"Account Number:", accountNumber,
		"Balance:", balance,
	)
	fmt.Println(ckAccountCarlos)
	fmt.Println(ckAccountYasmin)
}
