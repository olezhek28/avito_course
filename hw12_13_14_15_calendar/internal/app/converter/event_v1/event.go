package event_v1

import (
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
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
	var startDate *timestamppb.Timestamp
	if event.EventInfo.StartDate.Valid {
		startDate = timestamppb.New(event.EventInfo.StartDate.Time)
	}
	var endDate *timestamppb.Timestamp
	if event.EventInfo.EndDate.Valid {
		endDate = timestamppb.New(event.EventInfo.EndDate.Time)
	}

	var notificationInterval *durationpb.Duration
	if event.EventInfo.NotificationInterval != nil {
		notificationInterval = durationpb.New(*event.EventInfo.NotificationInterval)
	}

	var createdAt *timestamppb.Timestamp
	if event.CreatedAt.Valid {
		createdAt = timestamppb.New(event.CreatedAt.Time)
	}

	var updatedAt *timestamppb.Timestamp
	if event.UpdatedAt.Valid {
		updatedAt = timestamppb.New(event.UpdatedAt.Time)
	}

	var description *wrapperspb.StringValue
	if event.EventInfo.Description.Valid {
		description = &wrapperspb.StringValue{Value: event.EventInfo.Description.String}
	}

	return &desc.Event{
		Id: event.ID,
		EventInfo: &desc.EventInfo{
			Title:                event.EventInfo.Title,
			StartDate:            startDate,
			EndDate:              endDate,
			NotificationInterval: notificationInterval,
			Description:          description,
			OwnerId:              event.EventInfo.OwnerID,
		},
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
