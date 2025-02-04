package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"struct/bins"
)


func SaveBinListJson(data *bins.Bin) (*bins.BinList, error) {
	var name string
	fmt.Println("Введите названия для сохранение локального JSON файла: ")
	fmt.Scan(&name)

	file, err := os.Create(name)
	if err != nil {
		return nil, errors.New("Ошибка создание локально файла")
	}

	defer file.Close()

	jsonEncod := json.NewEncoder(file)
	if err := jsonEncod.Encode(data); err != nil {
		return nil, errors.New("при записи JSON-данных")
	}

	fmt.Printf("Данные сохранены в %s файл", name)
	return &bins.BinList{
		Bins: []bins.Bin{*data},
	}, nil
}

func ReadBinList(bin *bins.BinList) error {
	filename := "data.json"

	file, err := os.Open(filename)
	if err != nil {
		return errors.New("невозможно открыть файл")
	}

	defer file.Close()

	jsonDecode := json.NewDecoder(file)
	if err := jsonDecode.Decode(bin); err != nil {
		return errors.New("декодирования")
	}

	fmt.Printf("Успешно прочитано: %v\n", bin)

	return nil
}
