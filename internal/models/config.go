package models

import (
	"fmt"
	"os"
)

type Config struct {
	Enabled      bool
	DatabasePath string
	Port         string
}

func NewConfig() *Config {
	login := os.Getenv("GOSQLSERVERLOGIN")
	dsn := fmt.Sprintf("sqlserver://%s@localhost:1433?database=CarSales", login)

	return &Config{
		Enabled:      true,
		DatabasePath: dsn,
		Port:         "8000",
	}
}
