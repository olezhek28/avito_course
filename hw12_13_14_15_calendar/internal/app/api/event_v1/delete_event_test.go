package event_v1

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model/err"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository"
	repoMocks "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository/mocks"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/service/event"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/logger"
	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestImplementation_DeleteEvent(t *testing.T) {
	type eventRepositoryMockFunc func(mc *minimock.Controller) repository.EventRepository

	type args struct {
		ctx context.Context
		req *desc.DeleteEventRequest
	}

	var (
		mc  = minimock.NewController(t)
		ctx = context.Background()

		eventID = gofakeit.Int64()

		repoErr        = fmt.Errorf(gofakeit.Phrase())
		invalidIDError = status.Error(codes.InvalidArgument, err.ErrInvalidEventID)

		req = &desc.DeleteEventRequest{
			Id: &wrapperspb.Int64Value{Value: eventID},
		}

		invalidReq = &desc.DeleteEventRequest{
			Id: nil,
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
				mock.DeleteEventMock.Expect(ctx, eventID).Return(nil)
				return mock
			},
		},
		{
			name: "negative case - invalid event id",
			args: args{
				ctx: ctx,
				req: invalidReq,
			},
			want: nil,
			err:  invalidIDError,
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
				mock.DeleteEventMock.Expect(ctx, eventID).Return(repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := newMockEventV1(Implementation{
				eventService: event.NewService(logger.New(), tt.eventRepositoryMock(mc)),
			})

			res, err := api.DeleteEvent(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.want, res)
			require.Equal(t, tt.err, err)
		})
	}
}
