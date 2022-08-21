package config

import (
	"encoding/json"
	"io/ioutil"
)

// RabbitConsumer ...
type RabbitConsumer struct {
	DSN       string `json:"dsn"`
	QueueName string `json:"queue_name"`
}

// SenderConfig ...
type SenderConfig struct {
	RabbitConsumer *RabbitConsumer `json:"rabbit_consumer"`
}

// NewSenderConfig ...
func NewSenderConfig(path string) (*SenderConfig, error) {
	file, err := ioutil.ReadFile(path)
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

// GetRabbitConsumerConfig ...
func (c *SenderConfig) GetRabbitConsumerConfig() (*RabbitConsumer, error) {
	return c.RabbitConsumer, nil
}
