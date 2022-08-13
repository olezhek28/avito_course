package event_v1

import (
	"database/sql"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
)

// ToUpdateEvent ...
func ToUpdateEvent(updateEvent *desc.UpdateEventRequest_UpdateEventInfo) *model.UpdateEvent {
	if updateEvent == nil {
		return nil
	}

	return &model.UpdateEvent{
		Title: sql.NullString{
			String: updateEvent.GetTitle().GetValue(),
			Valid:  updateEvent.GetTitle() != nil,
		},
		Date: sql.NullTime{
			Time:  updateEvent.GetDate().AsTime(),
			Valid: updateEvent.GetDate() != nil,
		},
		Owner: sql.NullString{
			String: updateEvent.GetOwner().GetValue(),
			Valid:  updateEvent.GetOwner() != nil,
		},
	}
}
