package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var calculator = map[string] func(string,[]int) (int,error) {
	"AVG": avg,
	"SUM": sum,
	"MED": med,
}

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
		}

		calculators, err := calculator[input](input, operatic)
		if err != nil {
			fmt.Print(err)
			continue
		} else {
			fmt.Println(calculators)
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

func avg(userInput string, nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("Некорректно набранное число")
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	result := sum / len(nums)
	return result, nil
}

func sum(userInput string, nums []int) (int, error) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum, nil
}

func med(userInput string, nums []int) (int, error) {
	sort.Ints(nums)
	result := len(nums)
	if result%2 == 1 {
		return nums[result/2], nil
	} else {
		mid1 := nums[result/2-1]
		mid2 := nums[result/2]
		return (mid1 + mid2) / 2, nil
	}
}
