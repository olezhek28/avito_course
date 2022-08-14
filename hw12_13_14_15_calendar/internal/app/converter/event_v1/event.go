package event_v1

import (
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToDescEvents ...
func ToDescEvents(events []*model.Event) []*desc.Event {
	res := make([]*desc.Event, 0, len(events))
	for _, event := range events {
		res = append(res, ToDescEvent(event))
	}

	return res
}

// ToDescEvent ...
func ToDescEvent(event *model.Event) *desc.Event {
	var date *timestamppb.Timestamp
	if event.Date != nil {
		date = timestamppb.New(*event.Date)
	}

	return &desc.Event{
		Id: event.ID,
		EventInfo: &desc.EventInfo{
			Title: event.Title,
			Date:  date,
			Owner: event.Owner,
		},
	}
}
