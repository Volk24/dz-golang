package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		input, err := userInputOperation()
		if err != nil {
			fmt.Print(err)
			continue
		}

		number, err := numberInput()
		if err != nil {
			fmt.Print(err)
			continue
		}

		operatic, err := translateInt(number)
		if err != nil {
			fmt.Print(err)
			continue
		} else {
			fmt.Print(input, operatic)
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

func numberInput() (string, error) {
	fmt.Print("Введите числа через (,): ")
	scanner, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	if scanner == "" {
		return "", errors.New("Ввод чисел не правильный")
	} else {
		return scanner, nil
	}
}

func translateInt(scanner string) ([]int, error) {
	scanner = strings.TrimSpace(scanner)
	number := strings.Split(scanner, ",")
	var nums []int
	for _, str := range number {
		str = strings.TrimSpace(str)
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, errors.New("Не удалось перевести строку в число")
		} else {
			nums = append(nums, num)
		}
	}
	return nums, nil
}

func calculator(userInput string, nums []int) (int, error) {}