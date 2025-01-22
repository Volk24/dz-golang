package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"struct/bins"
)

func SerializeToJson(bin *bins.BinList) ([]byte, error) {
	jsonData, err := json.Marshal(bin)
	if err != nil {
		return nil, fmt.Errorf("Ошибка сериализации JSON: %w", err)
	}
	return jsonData, nil
}

func WriteFile(content []byte, name string) error {
	file, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("Ошибка создание файла: %w", err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return fmt.Errorf("Файл %s не существует", name)
	}
	if err != nil {
		return fmt.Errorf("Ошибка записи в файл: %w", err)
	}
	fmt.Println("Запись успешна")
	return nil
}

func SaveBinListJson(bins *bins.BinList) error {
	data, err := SerializeToJson(bins)
	if err != nil {
		return err
	}
	if err := WriteFile(data, "data.json"); err != nil {
		return fmt.Errorf("Ошибка записи файла: %w", err)
	}
	fmt.Println("Данные успешно сохранены")
	return nil
}


func ReadBinListJson(bins *bins.BinList) error {
	file, err := ReadFile("data.json")
	if err != nil {
		return err
	}
	err = ParsingJson(file, bins)
	if err != nil {
		return err
	}
	return nil
}

func ParsingJson(data []byte, bins *bins.BinList) error {
	err := json.Unmarshal(data, bins)
	if err != nil {
		return fmt.Errorf("Ошибка разбора JSON: %w", err)
	}
	return nil
}

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return nil, fmt.Errorf("Файл %s не существует", name)
	}
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения файла %s: %w", name, err)
	}
	return data, nil
}
