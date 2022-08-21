package scheduler

import (
	"context"
	"time"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
)

func (s *Service) Run(ctx context.Context) error {
	ticker := time.NewTicker(1 * time.Second)

	var count int64 = 0

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			s.SendEvent(&model.Event{
				ID: count,
			})
			count++
		}
	}

	return nil
}
