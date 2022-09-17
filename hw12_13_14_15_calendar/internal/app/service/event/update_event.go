package event

import (
	"context"
	"database/sql"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model/err"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateEvent ...
func (s *Service) UpdateEvent(ctx context.Context, eventID sql.NullInt64, updateEventInfo *model.UpdateEventInfo) error {
	if updateEventInfo == nil {
		s.logger.Error(err.ErrInvalidArguments)
		return status.Error(codes.InvalidArgument, err.ErrInvalidArguments)
	}
	if !eventID.Valid {
		s.logger.Error(err.ErrInvalidEventID)
		return status.Error(codes.InvalidArgument, err.ErrInvalidEventID)
	}

	updateEventInfo.NotificationDate = sql.NullTime{
		Time:  updateEventInfo.StartDate.Time,
		Valid: updateEventInfo.StartDate.Valid,
	}
	if updateEventInfo.NotificationInterval != nil {
		updateEventInfo.NotificationDate.Time = updateEventInfo.StartDate.Time.Add(-*updateEventInfo.NotificationInterval)
	}

	return s.eventRepository.UpdateEvent(ctx, eventID.Int64, updateEventInfo)
}
