package config

import (
	"os"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
	"gopkg.in/yaml.v3"
)

const (
	dbPassEscSeq = "{password}"
)

// LoggerConf ...
type LoggerConf struct {
	Level string `yaml:"level"`
}

// SourceConf ...
type SourceConf struct {
	SourceType string `yaml:"source_type"`
}

// DB ...
type DB struct {
	dsn                string `yaml:"dsn"`
	maxOpenConnections int32  `yaml:"max_open_connections"`
}

// Config ...
type Config struct {
	logger LoggerConf `yaml:"logger"`
	source SourceConf `yaml:"source"`
	db     DB         `yaml:"db"`
}

// New ...
func New(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	config := &Config{}
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// GetDbConfig ...
func (c *Config) GetDbConfig() (*pgxpool.Config, error) {
	password := "event-service-password"

	dbDsn := strings.ReplaceAll(c.db.DSN(), dbPassEscSeq, password)

	poolConfig, err := pgxpool.ParseConfig(dbDsn)
	if err != nil {
		return nil, err
	}
	poolConfig.ConnConfig.BuildStatementCache = nil
	poolConfig.ConnConfig.PreferSimpleProtocol = true
	poolConfig.MaxConns = c.db.MaxOpenConnections()

	return poolConfig, nil
}
