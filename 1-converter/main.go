package main

import (
	"errors"
	"fmt"
)

const usdEur float64 = 0.94691
const usdRub float64 = 100.0
const eurRub float64 = usdEur * usdRub

func main() {
	for {
		input, err := userInput()
		if err != nil {
			fmt.Println(err)
			continue
		}
		numbers, err := numberInput()
		if err != nil {
			fmt.Println(err)
			continue
		}

		currency, err := targetCurrency(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = calculatorCurrencies(numbers, input, currency)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			break
		}
	}
}

func userInput() (string, error) {
	var currency string
	fmt.Print("Введите название валюты (USD/EUR/RUB):")
	fmt.Scan(&currency)
	if currency == "Usd" || currency == "usd" || currency == "Eur" || currency == "eur" || currency == "Rub" || currency == "rub" {
		return currency, nil
	} else {
		return "", errors.New("Такой валюты нет! Введите заново")
	}
}

func numberInput() (float64, error) {
	var number float64
	fmt.Print("Введите число:")
	fmt.Scan(&number)
	if number > 0 {
		return number, nil
	} else {
		return 0, errors.New("Такое число не существует! Введите заново")
	}
}

func targetCurrency(currencies string) (string, error) {
	switch {
	case currencies == "Usd" || currencies == "usd":
		fmt.Println("Выберете подходящею валюту (EUR, RUB)")
		fmt.Scan(&currencies)
	case currencies == "Eur" || currencies == "eur":
		fmt.Println("Выберете подходящею валюту (USD, RUB)")
		fmt.Scan(&currencies)
	case currencies == "Rub" || currencies == "rub":
		fmt.Println("Выберете подходящею валюту (USD, EUR)")
		fmt.Scan(&currencies)
	default:
		return "", errors.New("Нет подходящий валюты")
	}
	return currencies, nil
}

func calculatorCurrencies(number float64, user, targetCurrency string)  error {
	switch {
	case user == "Usd" || user== "usd" && targetCurrency == "EUR" || targetCurrency == "eur":
		fmt.Println(number * usdEur)
	case user == "Usd" || user == "usd" && targetCurrency == "Rub" || targetCurrency == "rub":
		fmt.Println(number * usdRub)
	case user == "Eur" || user == "eur" && targetCurrency == "Rub" || targetCurrency == "rub":
		fmt.Println(number * eurRub)
	default:
		return errors.New("Неправильные вычисления")
	}
	return nil
}