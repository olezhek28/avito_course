package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/utils"
)

func (s *Service) Run(ctx context.Context) {
	ticker := time.NewTicker(s.checkPeriod)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := s.handleEvents(ctx)
			if err != nil {
				fmt.Printf("failed to handle events: %s", err.Error())
			}
		}
	}
}

func (s *Service) handleEvents(ctx context.Context) error {
	fmt.Printf("%v : Handle events start ...\n", time.Now())

	events, err := s.getEvents(ctx)
	if err != nil {
		return err
	}
	if len(events) == 0 {
		fmt.Println("No events.")
		return nil
	}

	err = s.sendEvent(events)
	if err != nil {
		return err
	}

	fmt.Println("Handle events success ...")
	return nil
}

func (s *Service) getEvents(ctx context.Context) ([]*model.Event, error) {
	endDate := utils.RoundUpToMinutes(time.Now())
	startDate := endDate.Add(-s.checkPeriod)

	return s.eventRepository.GetEventListByDate(ctx, startDate, endDate)
}

func (s *Service) sendEvent(event []*model.Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return s.rabbitProducer.Publish(data)
}
