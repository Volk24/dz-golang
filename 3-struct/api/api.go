package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"struct/bins"
	"struct/config"
)

const apiBaseUrl = "https://api.jsonbin.io/v3/b"

var apiKey *config.Config

type Db interface {
	SerializeToJson(*bins.BinList) ([]byte, error)
	ParsingJson([]byte, *bins.BinList) error
}

type JsonBin struct {
	bins.BinList
	db Db
}

func (data *JsonBin) CreateBin() (string, error) {
	jsonData, err := data.db.SerializeToJson(&data.BinList)
	if err != nil {
		return "", errors.New("Ошибка сереализации JSON файла")
	}

	req, err := http.NewRequest(http.MethodPost, apiBaseUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", errors.New("Ошибка Http запроса")
	}

	if apiKey == nil {
		return "", errors.New("Конфигурация API не установлена")
	}
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("X-Master-Key", apiKey.Key)

	client := &http.Client{}
	resq, err := client.Do(req)
	if err != nil {
		return "", errors.New("Ошибка выполнение запроса")
	}
	defer resq.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return "", errors.New("Ошибка чтение ответа")
	}

	if resq.StatusCode != http.StatusOK && resq.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("Ошибка: неожиданный статус ответа %d: %s", resq.StatusCode, string(body))
	}

	var result struct {
		Metadata struct {
			ID string `json:"id"`
		} `json:"metadata"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", errors.New("Ошибка декодирование ответа")
	}
	return result.Metadata.ID, nil
}

func (vault JsonBin) AddBin(bin bins.Bin) {
	vault.Bins =append(vault.Bins, bin)
}