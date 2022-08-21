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

// CalendarConfig ...
type CalendarConfig struct {
	Logger LoggerConf `json:"logger"`
	Source SourceConf `json:"source"`
	DB     DB         `json:"db"`
}

// NewCalendarConfig ...
func NewCalendarConfig(path string) (*CalendarConfig, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &CalendarConfig{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// GetDbConfig ...
func (c *CalendarConfig) GetDbConfig() (*pgxpool.Config, error) {
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
