package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"struct/api"
	"struct/bins"
	"struct/config"
	"struct/file"
	"struct/storage"

	"github.com/joho/godotenv"
)

func main() {
	create := flag.Bool("create", false, "Создать новый бин")
	update := flag.Bool("update", false, "Обновить бин")
	delete := flag.Bool("delete", false, "Удалить бин")
	get := flag.Bool("get", false, "Выводит bin по id")
	list := flag.Bool("list", false, "Создать новый бин")
	files := flag.String("file", "", "Ввод имя файла, которые нужно отправить в JSON BIN")
	id := flag.String("id", "", "ID пользователя на сервере JSON BIN")
	name := flag.String("name", "", "Название файла")

	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}
	apiKey, err := config.ReadEnv("API_KEY")
	if err != nil {
		log.Fatalf("Ошибка чтения API_KEY: %v", err)
	}
	apiClient := api.NewAPIClient(apiKey)

	if *create {
		vault := bins.BinList{}
		bin, err := CreateAccount(&vault)
		if err != nil {
			logError(err)
		}
		saveBin, err := storage.SaveBinListJson(bin)
		if err != nil {
			logError(err)
		}

		fileData, err := file.ReadFileJson(*files)
		if err != nil {
			log.Fatalf("Ошибка чтения %s файла", *files)
		}
		jsonBin, _, err := apiClient.CreateBin(fileData, saveBin)
		if err != nil {
			logError(err)
		}
		_, err = api.SaveJsonBin(jsonBin, *name)
		if err != nil {
			logError(err)
		}
	} else if *update {
		fileData, err := file.ReadFileJson(*files)
		if err != nil {
			log.Fatalf("Ошибка чтения %s файла", *files)
		}
		updateJson, err := apiClient.UpdateBin(fileData, *id)
		if err != nil {
			logError(err)
		}
		fmt.Println(updateJson)
	} else if *delete {
		data := &api.JsonBin{}
		if err := apiClient.DeleteBin(data, *id); err != nil {
			logError(err)
		}
		if err := file.DeleteFiles(); err != nil {
			logError(err)
		}
	} else if *get {
		data := &api.JsonBin{}
		if err := apiClient.ReadBin(data, *id); err != nil {
			logError(err)
		}
	} else if *list {
		var name string
		if err := api.List(name); err != nil {
			logError(err)
		}
	} else {
		os.Exit(1)
	}
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
