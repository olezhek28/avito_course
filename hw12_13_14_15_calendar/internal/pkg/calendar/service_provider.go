package calendar

import (
	"context"
	"log"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository"
	dbRepository "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository/db_repository"
	memoryRepository "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository/memory_repository"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/service/event"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/config"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/logger"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/pkg/db"
)

type serviceProvider struct {
	db         db.Client
	configPath string
	config     *config.CalendarConfig
	logger     *logger.Logger

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
		cfg, err := s.GetConfig().GetDBConfig()
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
func (s *serviceProvider) GetConfig() *config.CalendarConfig {
	if s.config == nil {
		cfg, err := config.NewCalendarConfig(s.configPath)
		if err != nil {
			log.Fatalf("failed to get config: %s", err.Error())
		}

		s.config = cfg
	}

	return s.config
}

func (s *serviceProvider) GetLogger() *logger.Logger {
	if s.logger == nil {
		s.logger = logger.New(s.GetConfig().GetLoggerConfig())
	}

	return s.logger
}

// GetEventRepository ...
func (s *serviceProvider) GetEventRepository(ctx context.Context) repository.EventRepository {
	if s.eventRepository == nil {
		if s.GetConfig().GetSourceConfig().SourceType == model.DBSource {
			s.eventRepository = dbRepository.NewEventRepository(s.GetDB(ctx))
		} else if s.GetConfig().GetSourceConfig().SourceType == model.MemorySource {
			s.eventRepository = memoryRepository.NewEventRepository()
		}
	}

	return s.eventRepository
}

// GetEventService ...
func (s *serviceProvider) GetEventService(ctx context.Context) *event.Service {
	if s.eventService == nil {
		s.eventService = event.NewService(
			s.GetLogger(),
			s.GetEventRepository(ctx),
		)
	}

	return s.eventService
}
