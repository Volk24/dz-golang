package main

import (
	"fmt"
	"log"
	"struct/api"
	"struct/bins"
	"struct/storage"

	"github.com/joho/godotenv"
)

func main() {
	vault := bins.BinList{}
	err := godotenv.Load()
	if err != nil {
		logError(err)
	}
	bin, err := CreateAccount(&vault)
	if err != nil {
		logError(err)
	}
	localBin, err := storage.SaveBinListJson(bin)
	if err != nil {
		logError(err)
	}
	metaData, err := api.CreateBin(localBin)
	if err != nil {
		logError(err)
	}
	fmt.Println(metaData)
}

func CreateAccount(vault *bins.BinList) (*bins.Bin, error) {
	id := promptData("Введите Id")
	private := promptBool("Хотите создать приватный аккаунт (yes/no)")
	name := promptData("Введите имя")

	myAccount, err := bins.NewBin(id, name, private)
	if err != nil {
		log.Panicln(err)
	}
	vault.AddBin(*myAccount)
	return myAccount, nil

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
