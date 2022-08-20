package event_v1

import (
	"database/sql"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
)

// ToEventInfo ...
func ToEventInfo(eventInfo *desc.EventInfo) *model.EventInfo {
	startDate := eventInfo.GetStartDate().AsTime()
	endDate := eventInfo.GetEndDate().AsTime()

	return &model.EventInfo{
		Title:     eventInfo.GetTitle(),
		StartDate: &startDate,
		EndDate:   &endDate,
		NotificationIntervalMin: sql.NullInt64{
			Int64: eventInfo.GetNotificationIntervalMin().GetValue(),
			Valid: eventInfo.GetNotificationIntervalMin() != nil,
		},
		Description: sql.NullString{
			String: eventInfo.GetDescription().GetValue(),
			Valid:  eventInfo.GetDescription() != nil,
		},
		OwnerID: eventInfo.GetOwnerId(),
	}
}
