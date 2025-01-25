package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	name := "data.json"
	file, err := os.Create(name)
	if err != nil {
		return nil, errors.New("Ошибка создание локально файла")
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("Ошибка сериализация данный")
	}

	if _, err := file.Write(jsonData); err != nil {
		return nil, errors.New("Ошибка записи в файл")
	}
	fmt.Printf("Данные сохранены в %s файл", name)
	return &bins.BinList{
		Bins: []bins.Bin{},
	}, nil
}

func ReadBinList(bin *bins.BinList) (*bins.BinList, error) {
	filename := "data.json"

	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("невозможно открыть файл")
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("чтения файла ")
	}

	if err := json.Unmarshal(data, bin); err != nil {
		return nil, errors.New("десериализации данных ")
	}
	return &bins.BinList{
		Bins: []bins.Bin{},
	}, nil

}
