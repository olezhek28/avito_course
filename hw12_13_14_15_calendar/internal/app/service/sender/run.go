package sender

import (
	"encoding/json"
	"fmt"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	"github.com/streadway/amqp"
)

func (s *Service) Run() error {
	msgChan, err := s.rabbitConsumer.Consume()
	if err != nil {
		return err
	}

	for msg := range msgChan {
		err = s.receiveEvents(msg)
		if err != nil {
			fmt.Printf("failed to receive events: %s/n", err.Error())
			continue
		}

		msg.Ack(false)
	}

	return nil
}

func (s *Service) receiveEvents(msg amqp.Delivery) error {
	fmt.Printf("Received a message: %s\n", msg.Body)

	var events []*model.Event
	err := json.Unmarshal(msg.Body, &events)
	if err != nil {
		return err
	}

	for _, event := range events {
		fmt.Printf("Event:  %d\n", event.ID)
		fmt.Printf(
			"Title: %s\n"+
				"StartDate: %s\n"+
				"EndDate: :%s\n"+
				"NotificationIntervalMin: %d\n"+
				"Description: %s\n"+
				"OwnerID: %d\n"+
				"CreatedAt: %s\n"+
				"UpdatedAt: %s\n\n",
			event.EventInfo.Title,
			event.EventInfo.StartDate,
			event.EventInfo.EndDate,
			event.EventInfo.NotificationIntervalMin.Int64,
			event.EventInfo.Description.String,
			event.EventInfo.OwnerID,
			event.CreatedAt,
			event.UpdatedAt,
		)

	}

	return nil
}
