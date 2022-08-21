package config

import (
	"encoding/json"
	"io/ioutil"
)

// RabbitProducer ...
type RabbitProducer struct {
	DSN       string `json:"dsn"`
	QueueName string `json:"queue_name"`
}

// SchedulerConfig ...
type SchedulerConfig struct {
	RabbitProducer *RabbitProducer `json:"rabbit_producer"`
}

// NewSchedulerConfig ...
func NewSchedulerConfig(path string) (*SchedulerConfig, error) {
	file, err := ioutil.ReadFile(path)
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
func (c *SchedulerConfig) GetRabbitProducerConfig() (*RabbitProducer, error) {
	return c.RabbitProducer, nil
}
