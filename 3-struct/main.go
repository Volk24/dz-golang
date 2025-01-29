package main

import (
	"fmt"
	"log"
	"struct/api"
	"struct/bins"
	"struct/config"
	"struct/storage"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}
	apiKey, err := config.ReadEnv("API_KEY")
	if err != nil {
		log.Fatalf("Ошибка чтения API_KEY: %v", err)
	}
	apiClient := api.NewAPIClient(apiKey)
	vault := bins.BinList{}
	bin, err := CreateAccount(&vault)
	if err != nil {
		logError(err)
	}
	localBin, err := storage.SaveBinListJson(bin)
	if err != nil {
		logError(err)
	}
	jsonBin, err := apiClient.CreateBin(localBin)
	if err != nil {
		logError(err)
	}
	jsonBins, err :=api.SaveJsonBin(jsonBin)
	if err != nil {
		logError(err)
	}
	fmt.Println(jsonBins)
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
