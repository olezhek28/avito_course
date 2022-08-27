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
			//msg.Ack(false)
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
