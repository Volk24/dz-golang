package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"struct/bins"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func SaveBinListJson(data *bins.Bin) (*bins.BinList, error) {
	var name string
	fmt.Println("Введите название для Json файла (.json):")
	fmt.Scan(&name)

	file, err := os.Create(name)
	if err != nil {
		return nil, errors.New("Ошибка создание локально файла")
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("Ошибка сериализация данный")
	}

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(jsonData); err != nil {
		return nil, errors.New("Ошибка при записи в JSON файл")
	}
	fmt.Printf("Данные сохранены в %s файл", name)
	return &bins.BinList{
		Bins: []bins.Bin{},
	}, nil
}
