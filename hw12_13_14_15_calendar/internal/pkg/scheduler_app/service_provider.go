package calendar_app

import (
	"log"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/service/scheduler"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/config"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/pkg/rabbit"
)

type serviceProvider struct {
	rabbit     rabbit.Client
	configPath string
	config     *config.SchedulerConfig

	schedulerService *scheduler.Service
}

func newServiceProvider(configPath string) *serviceProvider {
	return &serviceProvider{
		configPath: configPath,
	}
}

// GetRabbit ...
func (s *serviceProvider) GetRabbit() rabbit.Client {
	if s.rabbit == nil {
		cfg, err := s.GetConfig().GetRabbitConfig()
		if err != nil {
			log.Fatalf("failed to get rabbit config: %s", err.Error())
		}

		rc, err := rabbit.NewClient(cfg)
		if err != nil {
			log.Fatalf("can`t connect to rabbit err: %s", err.Error())
		}
		s.rabbit = rc
	}

	return s.rabbit
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
		s.schedulerService = scheduler.NewService(s.GetRabbit())
	}

	return s.schedulerService
}
