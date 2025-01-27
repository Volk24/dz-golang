package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"struct/bins"
	"struct/config"
)

const apiBaseUrl = "https://api.jsonbin.io/v3/b"

var apiKey = &config.Config{Key: "KEY"}

type JsonBin struct {
	JsonBins []bins.BinList
	MetaData struct {
		ID       string `json:"id"`
	} `json:"metaData"`
}

func CreateBin(bin *bins.BinList) (*JsonBin, error) {
	jsonData, err := json.Marshal(bin)
	if err != nil {
		return nil, errors.New("Ошибка преобразование в JSON формат")
	}
	req, err := http.NewRequest("POST", apiBaseUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, errors.New("Ошибка при создание запроса")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", apiKey.Key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Ошибка при создание запроса клиента")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Ошибка прочтения ответа")
	}
	var jsonBin JsonBin
	if err := json.Unmarshal(body, &jsonBin); err != nil {
		return nil, errors.New("Ошибка парсинга JSON данных")
	}
	return &JsonBin{
		JsonBins: []bins.BinList{},
		MetaData: jsonBin.MetaData,
}, nil
}
func (vault *JsonBin) AddBin(bin bins.BinList) {
	vault.JsonBins = append(vault.JsonBins, bin)
}

