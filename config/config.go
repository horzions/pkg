package config

import (
	"github.com/horzions/pkg/database"
	"github.com/horzions/pkg/serve"
)

type Config struct {
	Serve    *serve.Serve
	Database *database.Database
}

func NewConfig() *Config {
	return &Config{}
}
