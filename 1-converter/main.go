package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	usdEur float64 = 0.94691
	usdRub float64 = 100.0
	eurRub float64 = usdEur * usdRub
)

var exchangeRates = map[string]map[string]float64 {
	"USD" : {
		"EUR" : usdEur,
		"RUB" : usdRub,
	},
	"EUR" : {
		"RUB" : eurRub,
	},
	"RUB" : {
		"USD" : 1 / usdRub,
		"EUR" : 1 / usdEur,
	},
}

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
	if _, valid := exchangeRates[currency]; valid {
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

func targetCurrency(input string) (string, error) {
	var target string
	fmt.Printf("Выберите валюту для обменя на %s (доступные варианты:", input)
	for currency := range exchangeRates[input] {
		fmt.Printf("%s", currency)
	}
	fmt.Println("):")
	fmt.Scan(&target)
	target = strings.ToUpper(input)
	if _, valid := exchangeRates[input][target]; valid {
		return target, nil
	}
	return "", errors.New("Неправильный выбор валюты!")
}


func calculatorCurrencies(number float64, user, targetCurrency string)  error {
	if rate, ok := exchangeRates[user][targetCurrency]; ok {
		fmt.Printf("Результат: %.2f %s\n", number*rate, targetCurrency)
		return nil
	}
	return errors.New("Неправильные вычисления")
}