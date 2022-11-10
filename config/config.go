package config

import (
	"fmt"

	"github.com/horzions/pkg/app"
)

type Config struct {
	Port int `json:"port"`
}

func (c *Config) Load(a *app.App) {
	fmt.Println("load config")
}
