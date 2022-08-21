package scheduler

import (
	"encoding/json"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
)

// SendEvent ...
func (s *Service) SendEvent(event *model.Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return s.rabbitClient.Publish(data)
}
