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
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestImplementation_GetEventListForDay(t *testing.T) {
	type eventRepositoryMockFunc func(mc *minimock.Controller) repository.EventRepository

	type args struct {
		ctx context.Context
		req *desc.GetEventListForDayRequest
	}

	var (
		mc         = minimock.NewController(t)
		ctx        = context.Background()
		eventID    = gofakeit.Int64()
		eventTitle = gofakeit.Phrase()
		eventDate  = gofakeit.Date()
		eventOwner = gofakeit.Name()
		createdAt  = gofakeit.Date()
		updatedAt  = gofakeit.Date()

		repoErr = fmt.Errorf(gofakeit.Phrase())

		req = &desc.GetEventListForDayRequest{
			Date: timestamppb.New(eventDate),
		}

		eventsRepoRes = []*model.Event{
			{
				ID: eventID,
				EventInfo: &model.EventInfo{
					Title: eventTitle,
					Date:  &eventDate,
					Owner: eventOwner,
				},
				CreatedAt: &createdAt,
				UpdatedAt: &updatedAt,
			},
		}

		expectRes = &desc.GetEventListForDayResponse{
			Result: &desc.GetEventListForDayResponse_Result{
				Events: []*desc.Event{
					{
						Id: eventID,
						EventInfo: &desc.EventInfo{
							Title: eventTitle,
							Date:  timestamppb.New(eventDate),
							Owner: eventOwner,
						},
						CreatedAt: timestamppb.New(createdAt),
						UpdatedAt: timestamppb.New(updatedAt),
					},
				},
			},
		}
	)

	tests := []struct {
		name                string
		args                args
		want                *desc.GetEventListForDayResponse
		err                 error
		eventRepositoryMock eventRepositoryMockFunc
	}{
		{
			name: "positive case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: expectRes,
			err:  nil,
			eventRepositoryMock: func(mc *minimock.Controller) repository.EventRepository {
				mock := repoMocks.NewEventRepositoryMock(mc)
				mock.GetEventListForDayMock.Expect(ctx, eventDate).Return(eventsRepoRes, nil)
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
				mock.GetEventListForDayMock.Expect(ctx, eventDate).Return(nil, repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := newMockEventV1(Implementation{
				eventService: event.NewService(tt.eventRepositoryMock(mc)),
			})

			res, err := api.GetEventListForDay(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, res)
			assert.Equal(t, tt.err, err)
		})
	}
}
