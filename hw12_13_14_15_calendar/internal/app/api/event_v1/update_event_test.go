package event_v1

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model/err"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository"
	repoMocks "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository/mocks"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/service/event"
	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestImplementation_UpdateEvent(t *testing.T) {
	type eventRepositoryMockFunc func(mc *minimock.Controller) repository.EventRepository

	type args struct {
		ctx context.Context
		req *desc.UpdateEventRequest
	}

	var (
		mc         = minimock.NewController(t)
		ctx        = context.Background()
		eventID    = gofakeit.Int64()
		eventTitle = gofakeit.Phrase()
		eventDate  = gofakeit.Date()
		eventOwner = gofakeit.Name()

		repoErr               = fmt.Errorf(gofakeit.Phrase())
		invalidIDError        = status.Error(codes.InvalidArgument, err.ErrInvalidEventID)
		invalidArgumentsError = status.Error(codes.InvalidArgument, err.ErrInvalidArguments)

		req = &desc.UpdateEventRequest{
			Id: &wrapperspb.Int64Value{Value: eventID},
			UpdateEventInfo: &desc.UpdateEventRequest_UpdateEventInfo{
				Title: &wrapperspb.StringValue{Value: eventTitle},
				Date:  timestamppb.New(eventDate),
				Owner: &wrapperspb.StringValue{Value: eventOwner},
			},
		}

		reqWithInvalidID = &desc.UpdateEventRequest{
			Id: nil,
			UpdateEventInfo: &desc.UpdateEventRequest_UpdateEventInfo{
				Title: &wrapperspb.StringValue{Value: eventTitle},
				Date:  timestamppb.New(eventDate),
				Owner: &wrapperspb.StringValue{Value: eventOwner},
			},
		}

		reqWithInvalidArguments = &desc.UpdateEventRequest{
			Id:              nil,
			UpdateEventInfo: nil,
		}

		updateEventInfoRepoReq = &model.UpdateEventInfo{
			Title: sql.NullString{String: eventTitle, Valid: true},
			Date:  sql.NullTime{Time: eventDate, Valid: true},
			Owner: sql.NullString{String: eventOwner, Valid: true},
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
				mock.UpdateEventMock.Expect(ctx, eventID, updateEventInfoRepoReq).Return(nil)
				return mock
			},
		},
		{
			name: "negative case - invalid event id",
			args: args{
				ctx: ctx,
				req: reqWithInvalidID,
			},
			want: nil,
			err:  invalidIDError,
			eventRepositoryMock: func(mc *minimock.Controller) repository.EventRepository {
				return nil
			},
		},
		{
			name: "negative case - invalid arguments",
			args: args{
				ctx: ctx,
				req: reqWithInvalidArguments,
			},
			want: nil,
			err:  invalidArgumentsError,
			eventRepositoryMock: func(mc *minimock.Controller) repository.EventRepository {
				return nil
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
				mock.UpdateEventMock.Expect(ctx, eventID, updateEventInfoRepoReq).Return(repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := newMockEventV1(Implementation{
				eventService: event.NewService(tt.eventRepositoryMock(mc)),
			})

			res, err := api.UpdateEvent(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, res)
			assert.Equal(t, tt.err, err)
		})
	}
}
