package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgx/stdlib"
	"os"
	"path/filepath"
)

var logLevels = map[string]int{
	"DEBUG": 10,
	"INFO":  20,
	"WARN":  30,
	"ERROR": 40,
	"FATAL": 50,
}

type Database struct {
	User     string `json:"user"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

type Config struct {
	Postgres   Database `json:"postgres"`
	Redis      Database `json:"redis"`
	LoggerFile string   `json:"logger"`
	LogLevel   string   `json:"log_level"`
}

func (d *Database) GetPostgresDbConnection() (*sql.DB, error) {
	connString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=disable",
		d.User, d.Name, d.Password, d.Host, d.Port)

	db, err := sql.Open("pgx", connString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *Database) GetRedisDbConnection() (redis.Conn, error) {
	connString := fmt.Sprintf("redis://%s:@%s:%d/0",
		d.User, d.Host, d.Port)

	redisConn, err := redis.DialURL(connString)
	if err != nil {
		return nil, err
	}

	return redisConn, nil
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

func (c *Config) GetLoggerDir() string {
	return c.LoggerFile
}

func (c *Config) GetLogLevel() int {
	return logLevels[c.LogLevel]
}
