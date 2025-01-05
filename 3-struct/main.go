package main

import (
	"errors"
	"fmt"
	"time"
)

type bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	bins []bin
}

func newBin(id, name string, private bool) (*bin, error) {
	if id == "" {
		return nil, errors.New("INVALID_ID")
	}

	if name == "" {
		return nil, errors.New("INVALID_NAME")
	}

	newBin := &bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
	return newBin, nil
}

func (bl *BinList) addBin(bin bin) {
	bl.bins = append(bl.bins, bin)
}

func main() {
	id := promptData("Введите id")
	private := promptBool("Приватный? (yes/no)")
	name := promptData("Введите Имя")

	binList := BinList{}

	newBin, err := newBin(id, name, private)
	if err != nil {
		fmt.Println(err)
		return
	}

	binList.addBin(*newBin)
	fmt.Printf("id: %s", "Name: %s", "Private: %t", "CreatedAt: %s\n", newBin.id, newBin.name, newBin.private, newBin.createdAt)

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