package config

import (
	"database/sql"
	"encoding/json"
	"os"
	"path/filepath"
)

type Database struct {
	User string `json:"user"`
	Password string `json:"password"`
	Name string `json:"name"`
	Host string `json:"host"`
	Port int `json:"port"`

}

type Config struct {
	Database Database `json:"database"`
}

func getPostgresDbConnection(maxOpenConnections int = 10) *sql.DB {

}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, err
	}

	config := &Config{}
	if err := json.NewDecoder(file).Decode(config); err != nil {
		return nil, err
	}
	if err := file.Close(); err != nil {
		return nil, err
	}
	return config, nil
}

