package main

import "fmt"

func main() {
	holder := "Carlos"
	agencyNumber := 559
	accountNumber := 1032
	balance := 352.20

	fmt.Println("Hello", holder,
		"Agency Number:", agencyNumber,
		"Account Number:", accountNumber,
		"Balance:", balance,
	)

	holder2 := "Yasmin"
	agencyNumber2 := 159
	accountNumber2 := 357
	balance2 := 159159.15
	fmt.Println("Hello", holder2,
		"Agency Number:", agencyNumber2,
		"Account Number:", accountNumber2,
		"Balance:", balance2,
	)
}
