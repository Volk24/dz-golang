package main

import (
	"fmt"
	"log"
	"struct/api"
	"struct/bins"
)

func main() {
	bin := &api.JsonBin{}
	data := CreateAccount(bin)
	fmt.Println(data)
}

func CreateAccount(vault *api.JsonBin) *bins.Bin {
	id := promptData("Введите Id")
	private := promptBool("Будет аккаунт приватным (yes/no)")
	name := promptData("Введите имя")

	myAccount, err := bins.NewBin(id, name, private)
	if err != nil {
		log.Panicln(err)
		return nil
	}
	vault.AddBin(*myAccount)
	return myAccount
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

func promptBool(prompt string) bool {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res == "yes"
}

func logError(err error) {
	fmt.Printf("Ошибка: %s \n", err)
}
