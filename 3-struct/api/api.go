package api

import (
	"fmt"
	"struct/config"
)

func cfg() {
	cfg := &config.Config{
		Key: "KEY",
	}
	fmt.Println(cfg)//Сделано образно что бы go не ругался!
}