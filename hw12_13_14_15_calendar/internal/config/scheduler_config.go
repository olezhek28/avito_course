package config

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

// RabbitProducer ...
type RabbitProducer struct {
	DSN       string `json:"dsn"`
	QueueName string `json:"queue_name"`
}

// Scheduler ...
type Scheduler struct {
	CheckPeriodSec int64 `json:"check_period_sec"`
}

// SchedulerConfig ...
type SchedulerConfig struct {
	RabbitProducer *RabbitProducer `json:"rabbit_producer"`
	Scheduler      *Scheduler      `json:"scheduler"`
	DB             *DB             `json:"db"`
}

// NewSchedulerConfig ...
func NewSchedulerConfig(path string) (*SchedulerConfig, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &SchedulerConfig{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// GetRabbitProducerConfig ...
func (c *SchedulerConfig) GetRabbitProducerConfig() *RabbitProducer {
	return c.RabbitProducer
}

// GetSchedulerConfig ...
func (c *SchedulerConfig) GetSchedulerConfig() *Scheduler {
	return c.Scheduler
}

// GetDBConfig ...
func (c *SchedulerConfig) GetDBConfig() (*pgxpool.Config, error) {
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
