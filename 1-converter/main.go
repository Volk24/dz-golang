package main

import "fmt"

	const UsdEur float64 = 0.94691
	const UsdRub float64 = 100.0
	const EurRub float64 = UsdEur * UsdRub

func main() {
	inpud := userInpud()
}

func userInpud() (string) {
	var currency string
	fmt.Print("Введите название валюты:")
	fmt.Scan(&currency)
	return currency
}

func calculateCurrencies(num int, currency1, currency2 string) (int) {
	
}