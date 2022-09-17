package config

import (
	"encoding/json"
	"os"
)

// RabbitConsumer ...
type RabbitConsumer struct {
	DSN       string `json:"dsn"`
	QueueName string `json:"queue_name"`
}

// SenderConfig ...
type SenderConfig struct {
	Logger         *LoggerConf     `json:"logger"`
	RabbitConsumer *RabbitConsumer `json:"rabbit_consumer"`
}

// NewSenderConfig ...
func NewSenderConfig(path string) (*SenderConfig, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &SenderConfig{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// GetLoggerConfig ...
func (c *SenderConfig) GetLoggerConfig() *LoggerConf {
	return c.Logger
}

// GetRabbitConsumerConfig ...
func (c *SenderConfig) GetRabbitConsumerConfig() *RabbitConsumer {
	return c.RabbitConsumer
}
