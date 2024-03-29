package sender

import (
	"encoding/json"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	"github.com/streadway/amqp"
)

// Run ...
func (s *Service) Run() error {
	msgChan, err := s.rabbitConsumer.Consume()
	if err != nil {
		return err
	}

	for msg := range msgChan {
		err = s.receiveEvents(msg)
		if err != nil {
			s.logger.Error("failed to receive events: %s/n", err.Error())
			continue
		}

		msg.Ack(false)
	}

	return nil
}

func (s *Service) receiveEvents(msg amqp.Delivery) error {
	s.logger.Info("Received a message: %s\n", msg.Body)

	var events []*model.Event
	err := json.Unmarshal(msg.Body, &events)
	if err != nil {
		return err
	}

	for _, event := range events {
		s.logger.Info("Event:  %d\n", event.ID)
		s.logger.Info(
			"Title: %s\n"+
				"StartDate: %v\n"+
				"EndDate: :%v\n"+
				"NotificationInterval: %v\n"+
				"Description: %s\n"+
				"OwnerID: %d\n"+
				"CreatedAt: %v\n"+
				"UpdatedAt: %v\n\n",
			event.EventInfo.Title,
			event.EventInfo.StartDate,
			event.EventInfo.EndDate,
			event.EventInfo.NotificationInterval,
			event.EventInfo.Description.String,
			event.EventInfo.OwnerID,
			event.CreatedAt,
			event.UpdatedAt,
		)
	}

	return nil
}
