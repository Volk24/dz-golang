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
	"struct/file"
	"time"
)

const (
	apiBaseUrl = "https://api.jsonbin.io/v3/b"
)

type JsonBin struct {
	Bins     bins.BinList
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

func (key *APIClient) CreateBin(data []byte, dataBinlist *bins.BinList) (*JsonBin, string, error) {
	if data == nil {
		return nil, "", errors.New("Ошибка: данные пустые")
	}

	req, err := http.NewRequest("POST", apiBaseUrl, bytes.NewBuffer(data))
	if err != nil {
		return nil, "", errors.New("Ошибка при создание запроса")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", key.apiKey.Key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", errors.New("Ошибка при создание запроса клиента")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, "", fmt.Errorf("Ошибка HTTP: %s, Тело ответа: %s", resp.Status, body)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", errors.New("Ошибка прочтения ответа")
	}
	var apiResp JsonBin
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, "", errors.New("Ошибка парсинга JSON данных")
	}
	apiResp.Bins = *dataBinlist
	return &apiResp, apiResp.Metadata.ID, nil
}

func SaveJsonBin(data *JsonBin, name string) (*JsonBin, error) {
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
		Bins:     data.Bins,
		Metadata: data.Metadata,
	}, nil
}

func (key *APIClient) UpdateBin(data []byte, id string) (*JsonBin, error) {
	var file bins.Bin
	url := fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", id)

	newData := bins.Bin{
		ID:   "2",
		Name: "Dom",
	}

	dataJson, err := json.Marshal(newData)
	if err != nil {
		return nil, errors.New("Ошибка сереализаци данных ")
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(dataJson))
	if err != nil {
		return nil, errors.New("Ошибка при создание запроса")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", key.apiKey.Key)

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

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Бин успешно обновлен!")
		fmt.Printf("Ответ сервера: %s\n", body)
	} else {
		fmt.Printf("Ошибка: статус %d\n", resp.StatusCode)
		fmt.Printf("Сообщение: %s\n", body)
	}
	if err := json.Unmarshal(body, &file); err != nil {
		return nil, errors.New("Ошибка парсинга JSON данных")
	}
	updatedBin := &JsonBin{
		Bins: bins.BinList{
			Bins: []bins.Bin{file},
		},
		Metadata: struct {
			ID        string    `json:"id"`
			CreatedAt time.Time `json:"createdAt"`
		}{
			ID:        id,
			CreatedAt: time.Now(),
		},
	}

	return updatedBin, nil
}

func (key *APIClient) DeleteBin(data *JsonBin, id string) error {
	url := fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", id)

	dataJson, err := json.Marshal(data)
	if err != nil {
		return errors.New("Ошибка сереализаци данных ")
	}

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(dataJson))
	if err != nil {
		return errors.New("Ошибка при создание запроса")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", key.apiKey.Key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("Ошибка при создание запроса клиента")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Ошибка прочтения ответа")
	}

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Бин успешно удален!")
		fmt.Printf("Ответ сервера: %s\n", body)
	} else {
		fmt.Printf("Ошибка: статус %d\n", resp.StatusCode)
		fmt.Printf("Сообщение: %s\n", body)
	}
	return nil
}

func (key *APIClient) ReadBin(data *JsonBin, id string) error {
	url := fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", id)

	dataJson, err := json.Marshal(data)
	if err != nil {
		return errors.New("Ошибка сереализаци данных ")
	}

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(dataJson))
	if err != nil {
		return errors.New("Ошибка при создание запроса")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", key.apiKey.Key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("Ошибка при создание запроса клиента")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Ошибка прочтения ответа")
	}

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Бин успешно получен!")
		fmt.Printf("Ответ сервера: %s\n", body)
	} else {
		fmt.Printf("Ошибка: статус %d\n", resp.StatusCode)
		fmt.Printf("Сообщение: %s\n", body)
	}
	return nil
}

func List(name string) error {
	content, err := file.ReadFileJson(name)
	if err != nil {
		return errors.New("Ошибка чтение JSON файла")
	}

	var datas []JsonBin
	if err := json.Unmarshal(content, &datas); err != nil {
		return errors.New("Ошибка парсинга JSON файла")
	}

	for _, data := range datas {
		fmt.Printf("ID %s Name: %s\n", data.Metadata.ID, data.Bins.Bins[0].Name)
	}
	return nil
}
