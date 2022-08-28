package calendar_app

import (
	"log"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/service/sender"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/config"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/pkg/rabbit"
)

type serviceProvider struct {
	rabbitConsumer rabbit.Consumer
	configPath     string
	config         *config.SenderConfig

	senderService *sender.Service
}

func newServiceProvider(configPath string) *serviceProvider {
	return &serviceProvider{
		configPath: configPath,
	}
}

// GetRabbitConsumer ...
func (s *serviceProvider) GetRabbitConsumer() rabbit.Consumer {
	if s.rabbitConsumer == nil {
		rc, err := rabbit.NewConsumer(s.GetConfig().GetRabbitConsumerConfig())
		if err != nil {
			log.Fatalf("can`t connect to rabbit consumer err: %s", err.Error())
		}
		s.rabbitConsumer = rc
	}

	return s.rabbitConsumer
}

// GetConfig ...
func (s *serviceProvider) GetConfig() *config.SenderConfig {
	if s.config == nil {
		cfg, err := config.NewSenderConfig(s.configPath)
		if err != nil {
			log.Fatalf("failed to get config: %s", err.Error())
		}

		s.config = cfg
	}

	return s.config
}

// GetSenderService ...
func (s *serviceProvider) GetSenderService() *sender.Service {
	if s.senderService == nil {
		s.senderService = sender.NewService(s.GetRabbitConsumer())
	}

	return s.senderService
}
