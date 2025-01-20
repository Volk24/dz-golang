package file

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Files interface {
	ReadFileJson(string) ([]byte, error)
}

func ReadFileJson(name string) ([]byte, error) {
	var file string
	fmt.Print("Введите название JSON файла: ")
	fmt.Scan(&file)
	if !strings.HasSuffix(file, ".json") {
		return nil, errors.New("Файл не имеете расширение JSON")
	}

	filename, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("Ошибка открытия файла: %v", err)
	}
	defer filename.Close()

	data, err := io.ReadAll(filename)
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения файла: %v", err)
	}
	return data, nil
}