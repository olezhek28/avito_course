package app

import (
	"context"
	"log"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository"
	dbRepository "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository/db"
	memoryRepository "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository/memory"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/service/event"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/config"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/pkg/db"
)

type sourceType int64

const (
	dbSource sourceType = iota
	memorySource
)

type serviceProvider struct {
	db             db.Client
	configPath     string
	config         *config.Config
	dataSourceType sourceType

	// repositories
	eventRepository repository.EventRepository

	// services
	eventService *event.Service
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
func (s *serviceProvider) GetConfig() *config.Config {
	if s.config == nil {
		cfg, err := config.New(s.configPath)
		if err != nil {
			log.Fatalf("failed to get config: %s", err.Error())
		}

		s.config = cfg
	}

	return s.config
}

// GetEventRepository ...
func (s *serviceProvider) GetEventRepository(ctx context.Context) repository.EventRepository {
	if s.eventRepository == nil {
		if s.dataSourceType == dbSource {
			s.eventRepository = dbRepository.NewEventRepository(s.GetDB(ctx))
		} else if s.dataSourceType == memorySource {
			s.eventRepository = memoryRepository.NewEventRepository()
		}
	}

	return s.eventRepository
}

// GetEventService ...
func (s *serviceProvider) GetEventService(ctx context.Context) *event.Service {
	if s.eventService == nil {
		s.eventService = event.NewService(
			s.GetEventRepository(ctx),
		)
	}

	return s.eventService
}
