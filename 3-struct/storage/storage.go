package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"struct/bins"
)

func SaveBinListJson(bins bins.BinList) error {
	data, err := json.Marshal(bins)
	if err != nil {
		fmt.Println("Не удалось сериализовать данные в JSON")
	}
	err = WriteFile(data, "data.json")
	if err != nil {
		fmt.Println("Не удалось записать в файл")
		return err
	}
	fmt.Println("Данные успешно сохранены")
	return nil 
}

func WriteFile(content []byte, name string) error {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Запись успешна")
	return nil
}

func ReadBinListJson(bins bins.BinList) error {
	file, err := ReadFile("data.json")
	if err != nil {
		fmt.Println("Не удалось прочесть файл")
		return err
	}
	err = json.Unmarshal(file, &bins)
	if err != nil {
		fmt.Println("Не удалось разобрать файл JSON")
		return err
	}
	return nil
}

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
	}
	return data, nil
}