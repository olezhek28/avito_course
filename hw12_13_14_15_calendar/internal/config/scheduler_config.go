package config

import (
	"encoding/json"
	"io/ioutil"
)

// Rabbit ...
type Rabbit struct {
	DSN       string `json:"dsn"`
	QueueName string `json:"queue_name"`
}

// SchedulerConfig ...
type SchedulerConfig struct {
	Rabbit *Rabbit `json:"rabbit"`
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

// GetRabbitConfig ...
func (c *SchedulerConfig) GetRabbitConfig() (*Rabbit, error) {
	return c.Rabbit, nil
}
