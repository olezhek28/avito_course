package config

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	dbPassEscSeq = "{password}"
	password     = "event-service-password"
)

// LoggerConf ...
type LoggerConf struct {
	ShowTime bool `json:"show_time"`
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
	Logger *LoggerConf `json:"logger"`
	Source *SourceConf `json:"source"`
	DB     *DB         `json:"db"`
}

// NewCalendarConfig ...
func NewCalendarConfig(path string) (*CalendarConfig, error) {
	file, err := os.ReadFile(path)
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

// GetLoggerConfig ...
func (c *CalendarConfig) GetLoggerConfig() *LoggerConf {
	return c.Logger
}

// GetSourceConfig ...
func (c *CalendarConfig) GetSourceConfig() *SourceConf {
	return c.Source
}

// GetDBConfig ...
func (c *CalendarConfig) GetDBConfig() (*pgxpool.Config, error) {
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
