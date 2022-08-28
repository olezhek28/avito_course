package event_v1

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository"
	repoMocks "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository/mocks"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/service/event"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/config"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/logger"
	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestImplementation_CreateEvent(t *testing.T) {
	type eventRepositoryMockFunc func(mc *minimock.Controller) repository.EventRepository

	type args struct {
		ctx context.Context
		req *desc.CreateEventRequest
	}

	var (
		mc                   = minimock.NewController(t)
		ctx                  = context.Background()
		title                = gofakeit.Phrase()
		startDate            = gofakeit.Date()
		endDate              = gofakeit.Date()
		notificationInterval = time.Duration(gofakeit.Number(1, 100))
		description          = gofakeit.Phrase()
		ownerID              = gofakeit.Int64()
		loggerConf           = &config.LoggerConf{
			ShowTime: gofakeit.Bool(),
		}

		repoErr = fmt.Errorf(gofakeit.Phrase())

		req = &desc.CreateEventRequest{
			EventInfo: &desc.EventInfo{
				Title:                title,
				StartDate:            timestamppb.New(startDate),
				EndDate:              timestamppb.New(endDate),
				NotificationInterval: durationpb.New(notificationInterval),
				Description:          &wrapperspb.StringValue{Value: description},
				OwnerId:              ownerID,
			},
		}

		eventInfoRepoReq = &model.EventInfo{
			Title: title,
			NotificationDate: sql.NullTime{
				Time:  startDate.Add(-notificationInterval),
				Valid: true,
			},
			StartDate: sql.NullTime{
				Time:  startDate,
				Valid: true,
			},
			EndDate: sql.NullTime{
				Time:  endDate,
				Valid: true,
			},
			NotificationInterval: &notificationInterval,
			Description: sql.NullString{
				String: description,
				Valid:  true,
			},
			OwnerID: ownerID,
		}
	)

	tests := []struct {
		name                string
		args                args
		want                *emptypb.Empty
		err                 error
		eventRepositoryMock eventRepositoryMockFunc
	}{
		{
			name: "positive case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: &emptypb.Empty{},
			err:  nil,
			eventRepositoryMock: func(mc *minimock.Controller) repository.EventRepository {
				mock := repoMocks.NewEventRepositoryMock(mc)
				mock.CreateEventMock.Expect(ctx, eventInfoRepoReq).Return(nil)
				return mock
			},
		},
		{
			name: "negative case - repository error",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  repoErr,
			eventRepositoryMock: func(mc *minimock.Controller) repository.EventRepository {
				mock := repoMocks.NewEventRepositoryMock(mc)
				mock.CreateEventMock.Expect(ctx, eventInfoRepoReq).Return(repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := newMockEventV1(Implementation{
				eventService: event.NewService(logger.New(loggerConf), tt.eventRepositoryMock(mc)),
			})

			res, err := api.CreateEvent(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.want, res)
			require.Equal(t, tt.err, err)
		})
	}
}
