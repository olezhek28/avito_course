package config

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	dbPassEscSeq = "{password}"
)

// LoggerConf ...
type LoggerConf struct {
	Level string `json:"level"`
}

// SourceConf ...
type SourceConf struct {
	SourceType string `json:"source_type"`
}

// DB ...
type DB struct {
	DSN                string `json:"dsn"`
	MaxOpenConnections int32  `json:"max_open_connections"`
}

// Config ...
type Config struct {
	Logger LoggerConf `json:"logger"`
	Source SourceConf `json:"source"`
	DB     DB         `json:"db"`
}

// New ...
func New(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// GetDbConfig ...
func (c *Config) GetDbConfig() (*pgxpool.Config, error) {
	password := "event-service-password"

	dbDsn := strings.ReplaceAll(c.DB.DSN, dbPassEscSeq, password)

	poolConfig, err := pgxpool.ParseConfig(dbDsn)
	if err != nil {
		return nil, err
	}
	poolConfig.ConnConfig.BuildStatementCache = nil
	poolConfig.ConnConfig.PreferSimpleProtocol = true
	poolConfig.MaxConns = c.DB.MaxOpenConnections

	return poolConfig, nil
}
