package calendar_app

import (
	"log"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/service/scheduler"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/config"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/pkg/rabbit"
)

type serviceProvider struct {
	rabbitProducer rabbit.Producer
	configPath     string
	config         *config.SchedulerConfig

	schedulerService *scheduler.Service
}

func newServiceProvider(configPath string) *serviceProvider {
	return &serviceProvider{
		configPath: configPath,
	}
}

// GetRabbitProducer ...
func (s *serviceProvider) GetRabbitProducer() rabbit.Producer {
	if s.rabbitProducer == nil {
		cfg, err := s.GetConfig().GetRabbitProducerConfig()
		if err != nil {
			log.Fatalf("failed to get rabbit producer config: %s", err.Error())
		}

		rp, err := rabbit.NewProducer(cfg)
		if err != nil {
			log.Fatalf("can`t connect to rabbit producer err: %s", err.Error())
		}
		s.rabbitProducer = rp
	}

	return s.rabbitProducer
}

// GetConfig ...
func (s *serviceProvider) GetConfig() *config.SchedulerConfig {
	if s.config == nil {
		cfg, err := config.NewSchedulerConfig(s.configPath)
		if err != nil {
			log.Fatalf("failed to get config: %s", err.Error())
		}

		s.config = cfg
	}

	return s.config
}

// GetSchedulerService ...
func (s *serviceProvider) GetSchedulerService() *scheduler.Service {
	if s.schedulerService == nil {
		s.schedulerService = scheduler.NewService(s.GetRabbitProducer())
	}

	return s.schedulerService
}
