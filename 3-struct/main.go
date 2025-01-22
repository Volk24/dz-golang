package main

import (
	"fmt"
)

func main() {

}

// func CreateAccount(vault *api.JsonBin) {
// 	id := promptData("Введите Id")
// 	private := promptBool("Будет аккаунт приватным (yes/no)")
// 	name := promptData("Введите имя")

// 	myAccount, err := bins.NewBin(id, name, private)
// 	if err != nil {
// 		log.Panicln(err)
// 		return
// 	}
// 	vault.AddBin(*myAccount)
// }

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
