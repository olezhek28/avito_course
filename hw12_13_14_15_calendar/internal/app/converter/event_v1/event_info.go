package event_v1

import (
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
)

// ToEventInfo ...
func ToEventInfo(event *desc.EventInfo) *model.EventInfo {
	date := event.GetDate().AsTime()

	return &model.EventInfo{
		Title: event.GetTitle(),
		Date:  &date,
		Owner: event.GetOwner(),
	}
}
