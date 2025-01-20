package main

import (
	"fmt"
	"struct/bins"
	"struct/file"
	"struct/storage"
)

type Assembly interface {
	bins.Bins
	file.Files
	storage.Storage
}

func main() {
	id := promptData("Введите id")
	private := promptBool("Приватный? (yes/no)")
	name := promptData("Введите Имя")

	binList := bins.BinList{}

	newBin, err := bins.NewBin(id, name, private)
	if err != nil {
		logError(err)
		return
	}

	binList.AddBin(*newBin)
	fmt.Printf("id: %s, Name: %s, Private: %t, CreatedAt: %s\n", newBin.ID, newBin.Name, newBin.Private, newBin.CreatedAt)

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
