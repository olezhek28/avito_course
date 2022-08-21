package sender

import (
	"encoding/json"
	"log"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
)

func (s *Service) Run() error {
	msgChan, err := s.rabbitConsumer.Consume()
	if err != nil {
		return err
	}

	for m := range msgChan {
		log.Printf("Received a message: %s", m.Body)
		event := &model.Event{}
		json.Unmarshal(m.Body, &event)
		log.Printf("Event: %+v", event)

		m.Ack(false)
	}

	return nil
}
