package db_repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository/table"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/pkg/db"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/utils"
)

type eventRepository struct {
	db db.Client
}

func NewEventRepository(db db.Client) repository.EventRepository {
	return &eventRepository{
		db: db,
	}
}

// CreateEvent ...
func (r *eventRepository) CreateEvent(ctx context.Context, eventInfo *model.EventInfo) error {
	builder := sq.Insert(table.Event).
		PlaceholderFormat(sq.Dollar).
		Columns("title", "notification_date", "start_date", "end_date", "notification_interval", "description", "owner_id").
		Values(eventInfo.Title, eventInfo.NotificationDate, eventInfo.StartDate, eventInfo.EndDate, eventInfo.NotificationInterval.Nanoseconds(), eventInfo.Description, eventInfo.OwnerID)

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "repository.CreateEvent",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, v...)

	return err
}

// UpdateEvent ...
func (r *eventRepository) UpdateEvent(ctx context.Context, eventID int64, updateEventInfo *model.UpdateEventInfo) error {
	builder := sq.Update(table.Event).
		PlaceholderFormat(sq.Dollar).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": eventID})

	if updateEventInfo.Title.Valid {
		builder = builder.Set("title", updateEventInfo.Title.String)
	}
	if updateEventInfo.NotificationDate.Valid {
		builder = builder.Set("notification_date", updateEventInfo.NotificationDate.Time)
	}
	if updateEventInfo.StartDate.Valid {
		builder = builder.Set("start_date", updateEventInfo.StartDate.Time)
	}
	if updateEventInfo.EndDate.Valid {
		builder = builder.Set("end_date", updateEventInfo.EndDate.Time)
	}
	if updateEventInfo.NotificationInterval != nil {
		builder = builder.Set("notification_interval", updateEventInfo.NotificationInterval.Nanoseconds())
	}
	if updateEventInfo.Description.Valid {
		builder = builder.Set("description", updateEventInfo.Description.String)
	}
	if updateEventInfo.OwnerID.Valid {
		builder = builder.Set("owner_id", updateEventInfo.OwnerID.Int64)
	}

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "repository.UpdateEvent",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, v...)

	return err
}

// DeleteEvent ...
func (r *eventRepository) DeleteEvent(ctx context.Context, eventID int64) error {
	builder := sq.Delete(table.Event).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": eventID})

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "repository.DeleteEvent",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, v...)

	return err
}

// GetEventListForDay ...
func (r *eventRepository) GetEventListForDay(ctx context.Context, date time.Time) ([]*model.Event, error) {
	builder := sq.Select("id, title, start_date, end_date, notification_interval, description, owner_id, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		From(table.Event).
		Where(sq.GtOrEq{"start_date": utils.BeginningOfDay(date)}).
		Where(sq.LtOrEq{"start_date": utils.EndOfDay(date)})

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "repository.GetEventListForDay",
		QueryRaw: query,
	}

	var res []*model.Event
	err = r.db.DB().SelectContext(ctx, &res, q, v...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetEventListForWeek ...
func (r *eventRepository) GetEventListForWeek(ctx context.Context, weekStart time.Time) ([]*model.Event, error) {
	builder := sq.Select("id, title, start_date, end_date, notification_interval, description, owner_id, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		From(table.Event).
		Where(sq.GtOrEq{"start_date": utils.BeginningOfDay(weekStart)}).
		Where(sq.LtOrEq{"start_date": utils.EndOfDay(weekStart).AddDate(0, 0, utils.DaysInWeek)})

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "repository.GetEventListForWeek",
		QueryRaw: query,
	}

	var res []*model.Event
	err = r.db.DB().SelectContext(ctx, &res, q, v...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetEventListForMonth ...
func (r *eventRepository) GetEventListForMonth(ctx context.Context, monthStart time.Time) ([]*model.Event, error) {
	builder := sq.Select("id, title, start_date, end_date, notification_interval, description, owner_id, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		From(table.Event).
		Where(sq.GtOrEq{"start_date": utils.BeginningOfDay(monthStart)}).
		Where(sq.LtOrEq{"start_date": utils.EndOfDay(monthStart).AddDate(0, 0, utils.DaysInMonth)})

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "repository.GetEventListForWeek",
		QueryRaw: query,
	}

	var res []*model.Event
	err = r.db.DB().SelectContext(ctx, &res, q, v...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetEventListByDate ...
func (r *eventRepository) GetEventListByDate(ctx context.Context, startDate time.Time, endDate time.Time) ([]*model.Event, error) {
	builder := sq.Select("id, title, notification_date, start_date, end_date, notification_interval, description, owner_id, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		From(table.Event).
		Where(sq.Or{
			sq.And{
				sq.Gt{"start_date": startDate},
				sq.LtOrEq{"start_date": endDate},
			},
			sq.And{
				sq.Gt{"notification_date": startDate},
				sq.LtOrEq{"notification_date": endDate},
			},
		})

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "repository.GetEventListByDate",
		QueryRaw: query,
	}

	var res []*model.Event
	err = r.db.DB().SelectContext(ctx, &res, q, v...)
	if err != nil {
		return nil, err
	}

	return res, nil
}
