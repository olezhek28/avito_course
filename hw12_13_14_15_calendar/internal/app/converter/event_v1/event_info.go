package event_v1

import (
	"database/sql"
	"time"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
)

// ToEventInfo ...
func ToEventInfo(eventInfo *desc.EventInfo) *model.EventInfo {
	var notificationInterval time.Duration
	if eventInfo.NotificationInterval != nil {
		notificationInterval = eventInfo.NotificationInterval.AsDuration()
	}

	return &model.EventInfo{
		Title: eventInfo.GetTitle(),
		StartDate: sql.NullTime{
			Time:  eventInfo.GetStartDate().AsTime(),
			Valid: eventInfo.GetStartDate() != nil,
		},
		EndDate: sql.NullTime{
			Time:  eventInfo.GetEndDate().AsTime(),
			Valid: eventInfo.GetEndDate() != nil,
		},
		NotificationInterval: &notificationInterval,
		Description: sql.NullString{
			String: eventInfo.GetDescription().GetValue(),
			Valid:  eventInfo.GetDescription() != nil,
		},
		OwnerID: eventInfo.GetOwnerId(),
	}
}
