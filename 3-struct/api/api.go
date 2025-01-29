package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"struct/bins"
	"struct/config"
	"time"
)

const (
	apiBaseUrl = "https://api.jsonbin.io/v3/b"
)

type JsonBin struct {
	Bins bins.BinList
    Metadata struct {
        ID        string    `json:"id"`
        CreatedAt time.Time `json:"createdAt"`
    } `json:"metadata"`
}

type APIClient struct {
    apiKey config.Config
}

func NewAPIClient(key *config.Config) *APIClient {
    return &APIClient{
        apiKey: *key,
    }
}

func (key *APIClient) CreateBin(data *bins.BinList) (*JsonBin, error) {
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
	req.Header.Set("X-Master-Key",key.apiKey.Key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Ошибка при создание запроса клиента")
	}
	defer resp.Body.Close()

	 if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body)
        return nil, fmt.Errorf("Ошибка HTTP: %s, Тело ответа: %s", resp.Status, body)
    }

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Ошибка прочтения ответа")
	}
	var apiResp JsonBin
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, errors.New("Ошибка парсинга JSON данных")
	}
	apiResp.Bins = *data
	return &apiResp, nil
}

func SaveJsonBin(data *JsonBin) (*JsonBin, error) {
	var name string
	fmt.Println("Введите названия для JSON файла: ")
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
	return &JsonBin{
		Bins: data.Bins,
	}, nil
}