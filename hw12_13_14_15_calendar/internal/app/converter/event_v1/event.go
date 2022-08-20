package event_v1

import (
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
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
	if event.EventInfo.StartDate != nil {
		startDate = timestamppb.New(*event.EventInfo.StartDate)
	}
	var endDate *timestamppb.Timestamp
	if event.EventInfo.EndDate != nil {
		endDate = timestamppb.New(*event.EventInfo.EndDate)
	}

	var createdAt *timestamppb.Timestamp
	if event.CreatedAt != nil {
		createdAt = timestamppb.New(*event.CreatedAt)
	}

	var updatedAt *timestamppb.Timestamp
	if event.UpdatedAt != nil {
		updatedAt = timestamppb.New(*event.UpdatedAt)
	}

	var description *wrapperspb.StringValue
	if event.EventInfo.Description.Valid {
		description = &wrapperspb.StringValue{Value: event.EventInfo.Description.String}
	}

	return &desc.Event{
		Id: event.ID,
		EventInfo: &desc.EventInfo{
			Title:       event.EventInfo.Title,
			StartDate:   startDate,
			EndDate:     endDate,
			Description: description,
			OwnerId:     event.EventInfo.OwnerID,
		},
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
