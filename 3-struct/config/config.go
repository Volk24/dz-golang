package config

import (
	"fmt"
	"os"
)

type Config struct {
	Key string
}

func readEnv(key string) (*Config, error) {
	if key == "" {
		return nil, fmt.Errorf("Имя ENV не может быть пусты")
	}

	envKey := os.Getenv(key)
	if envKey == "" {
		return nil, fmt.Errorf("переменная окружения %s не найдена или пуста", envKey)
	}

	return &Config{
		Key: key,
	}, nil
}
