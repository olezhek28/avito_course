package event_v1

import (
	"database/sql"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
)

// ToUpdateEventInfo ...
func ToUpdateEventInfo(updateEventInfo *desc.UpdateEventRequest_UpdateEventInfo) *model.UpdateEventInfo {
	if updateEventInfo == nil {
		return nil
	}

	return &model.UpdateEventInfo{
		Title: sql.NullString{
			String: updateEventInfo.GetTitle().GetValue(),
			Valid:  updateEventInfo.GetTitle() != nil,
		},
		StartDate: sql.NullTime{
			Time:  updateEventInfo.GetStartDate().AsTime(),
			Valid: updateEventInfo.GetStartDate() != nil,
		},
		EndDate: sql.NullTime{
			Time:  updateEventInfo.GetEndDate().AsTime(),
			Valid: updateEventInfo.GetEndDate() != nil,
		},
		NotificationIntervalMin: sql.NullInt64{
			Int64: updateEventInfo.GetNotificationIntervalMin().GetValue(),
			Valid: updateEventInfo.GetNotificationIntervalMin() != nil,
		},
		Description: sql.NullString{
			String: updateEventInfo.GetDescription().GetValue(),
			Valid:  updateEventInfo.GetDescription() != nil,
		},
		OwnerID: sql.NullInt64{
			Int64: updateEventInfo.GetOwnerId().GetValue(),
			Valid: updateEventInfo.GetOwnerId() != nil,
		},
	}
}
