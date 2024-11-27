package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		inpud, err := userInputOperation()
		if err != nil {
			fmt.Print(err)
			continue
		}

		number, err := numberInput()
		if err != nil {
			fmt.Print(err)
			continue
		} else {
			fmt.Println(inpud, number)
			break
		}
	}
}

func userInputOperation() (string, error) {
	var operations string
	fmt.Print("Выберете операцию (AVG/SUM/MED): ")
	fmt.Scan(&operations)
	if operations == "AVG" || operations == "SUM" || operations == "MED" {
		return operations, nil
	} else {
		return "", errors.New("Такой операции нет! Попробуйте заново")
	}
}

func numberInput() ([]string, error) {
	fmt.Print("Введите числа через (,): ")
	numbers, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	numbers = strings.TrimSpace(numbers)
	if numbers == "" {
		return nil, errors.New("Ввод чисел не правильный")
	} else {
		num := strings.Split(numbers, ",")
		for v := range num {
			num[v] = strings.TrimSpace(num[v])
		}
		return num, nil
	}
}