package event_v1

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository"
	repoMocks "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository/mocks"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/service/event"
	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestImplementation_CreateEvent(t *testing.T) {
	type eventRepositoryMockFunc func(mc *minimock.Controller) repository.EventRepository

	type args struct {
		ctx context.Context
		req *desc.CreateEventRequest
	}

	var (
		mc         = minimock.NewController(t)
		ctx        = context.Background()
		eventTitle = gofakeit.Phrase()
		eventDate  = gofakeit.Date()
		eventOwner = gofakeit.Name()

		repoErr = fmt.Errorf(gofakeit.Phrase())

		req = &desc.CreateEventRequest{
			EventInfo: &desc.EventInfo{
				Title: eventTitle,
				Date:  timestamppb.New(eventDate),
				Owner: eventOwner,
			},
		}

		eventInfoRepoReq = &model.EventInfo{
			Title: eventTitle,
			Date:  &eventDate,
			Owner: eventOwner,
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
				eventService: event.NewService(tt.eventRepositoryMock(mc)),
			})

			res, err := api.CreateEvent(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, res)
			assert.Equal(t, tt.err, err)
		})
	}
}
