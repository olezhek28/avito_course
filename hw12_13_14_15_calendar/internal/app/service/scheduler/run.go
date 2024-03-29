package scheduler

import (
	"context"
	"encoding/json"
	"time"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/utils"
)

// 1 year ...
const eventTTL = time.Second * 60 * 60 * 24 * 365

// Run ...
func (s *Service) Run(ctx context.Context) {
	ticker := time.NewTicker(s.checkPeriod)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := s.handleEvents(ctx)
			if err != nil {
				s.logger.Error("failed to handle events: %s", err.Error())
			}
		}
	}
}

func (s *Service) handleEvents(ctx context.Context) error {
	s.logger.Info("%v : Handle events start ...\n", time.Now())

	events, err := s.getEvents(ctx)
	if err != nil {
		return err
	}
	if len(events) == 0 {
		s.logger.Info("No events.")
		return nil
	}

	err = s.sendEvent(events)
	if err != nil {
		return err
	}

	err = s.deleteOldEvents(ctx)
	if err != nil {
		return err
	}

	s.logger.Info("Handle events success ...")
	return nil
}

func (s *Service) getEvents(ctx context.Context) ([]*model.Event, error) {
	endDate := utils.RoundUpToMinutes(time.Now())
	startDate := endDate.Add(-s.checkPeriod)

	s.logger.Info("select events from %v to %v\n", startDate, endDate)

	return s.eventRepository.GetEventListByDate(ctx, startDate, endDate)
}

func (s *Service) sendEvent(event []*model.Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return s.rabbitProducer.Publish(data)
}

func (s *Service) deleteOldEvents(ctx context.Context) error {
	date := time.Now().Add(-eventTTL)

	return s.eventRepository.DeleteEventsBeforeDate(ctx, date)
}
