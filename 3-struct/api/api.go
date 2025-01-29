package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"struct/bins"
	"struct/config"
	"time"
)

const (
	apiBaseUrl = "https://api.jsonbin.io/v3/b"
)

var apiKey = config.Config{}

type MetaData struct {
	ParentId  string    `json:"parentId"`
	CreatedAt time.Time `json:"createdAt"`
}

type JsonBin struct {
	BinList  bins.BinList
	MetaData MetaData `json:"metaData"`
}

func CreateBin(data *bins.BinList) (*MetaData, error) {
	if data == nil {
		return nil, errors.New("Ошибка данные в BinList пустые")
	}
	dataJson, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("Ошибка сереализаци данных ")
	}

	req, err := http.NewRequest("POST", apiBaseUrl, bytes.NewBuffer(dataJson))
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

	metaDta := MetaData{
		ParentId:  string(body),
		CreatedAt: time.Now(),
	}

	if err := json.Unmarshal(body, &metaDta); err != nil {
		return nil, errors.New("Ошибка парсинга JSON данных")
	}
	return &MetaData{}, nil
}
