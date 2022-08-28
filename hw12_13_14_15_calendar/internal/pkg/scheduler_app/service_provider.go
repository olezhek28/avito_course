package calendar_app

import (
	"context"
	"log"
	"time"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository"
	dbRepository "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository/db"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/service/scheduler"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/config"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/pkg/db"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/pkg/rabbit"
)

type serviceProvider struct {
	db             db.Client
	rabbitProducer rabbit.Producer
	configPath     string
	config         *config.SchedulerConfig

	// repositories
	eventRepository repository.EventRepository

	// services
	schedulerService *scheduler.Service
}

func newServiceProvider(configPath string) *serviceProvider {
	return &serviceProvider{
		configPath: configPath,
	}
}

// GetDB ...
func (s *serviceProvider) GetDB(ctx context.Context) db.Client {
	if s.db == nil {
		cfg, err := s.GetConfig().GetDbConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

		dbc, err := db.NewClient(ctx, cfg)
		if err != nil {
			log.Fatalf("can`t connect to db err: %s", err.Error())
		}
		s.db = dbc
	}

	return s.db
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

// GetRabbitProducer ...
func (s *serviceProvider) GetRabbitProducer() rabbit.Producer {
	if s.rabbitProducer == nil {
		rp, err := rabbit.NewProducer(s.GetConfig().GetRabbitProducerConfig())
		if err != nil {
			log.Fatalf("can`t connect to rabbit producer err: %s", err.Error())
		}
		s.rabbitProducer = rp
	}

	return s.rabbitProducer
}

// GetEventRepository ...
func (s *serviceProvider) GetEventRepository(ctx context.Context) repository.EventRepository {
	if s.eventRepository == nil {
		s.eventRepository = dbRepository.NewEventRepository(s.GetDB(ctx))
	}

	return s.eventRepository
}

// GetSchedulerService ...
func (s *serviceProvider) GetSchedulerService(ctx context.Context) *scheduler.Service {
	if s.schedulerService == nil {
		s.schedulerService = scheduler.NewService(
			s.GetRabbitProducer(),
			s.GetEventRepository(ctx),
			time.Duration(s.GetConfig().GetSchedulerConfig().CheckPeriodSec)*time.Second,
		)
	}

	return s.schedulerService
}
