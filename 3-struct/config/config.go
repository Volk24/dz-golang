package config

import (
	"fmt"

	"github.com/qiangxue/go-env"
)

type Config struct {
	Key string
}

func readEnv(key string) (*Config, error) {
	if key == "" {
		return nil, fmt.Errorf("Имя ENV не может быть пусты")
	}
	if err := env.Load(key); err != nil {
		return nil, fmt.Errorf("Ошибка загрузки ENV %s: %w", key, err)
	}
	return &Config{
		Key: key,
	}, nil
}