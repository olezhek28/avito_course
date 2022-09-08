package e2e_test

import (
	"context"
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
	. "github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const grpcHost = "localhost:7002"

var _ = Describe("Calendar", func() {
	var todayEvent, weekAgoEvent, monthAgoEvent, failedEventInfo *desc.EventInfo
	var grpcClient desc.EventServiceV1Client

	var eventRes1, eventRes2, eventRes3 *desc.CreateEventResponse

	ctx := context.Background()
	now := time.Now()
	weekAgo := now.AddDate(0, 0, -7)
	monthAgo := now.AddDate(0, -1, 1)

	BeforeEach(func() {
		todayEvent = &desc.EventInfo{
			Title:       gofakeit.JobTitle(),
			StartDate:   timestamppb.New(now),
			Description: &wrapperspb.StringValue{Value: gofakeit.Phrase()},
			OwnerId:     int64(gofakeit.Number(1, 100)),
		}

		weekAgoEvent = &desc.EventInfo{
			Title:       gofakeit.JobTitle(),
			StartDate:   timestamppb.New(weekAgo),
			Description: &wrapperspb.StringValue{Value: gofakeit.Phrase()},
			OwnerId:     int64(gofakeit.Number(1, 100)),
		}

		monthAgoEvent = &desc.EventInfo{
			Title:       gofakeit.JobTitle(),
			StartDate:   timestamppb.New(monthAgo),
			Description: &wrapperspb.StringValue{Value: gofakeit.Phrase()},
			OwnerId:     int64(gofakeit.Number(1, 100)),
		}

		failedEventInfo = &desc.EventInfo{
			Title:   gofakeit.JobTitle(),
			OwnerId: gofakeit.Int64(),
		}

		//nolint:staticcheck
		con, err := grpc.Dial(grpcHost, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("didn't connect: %s", err.Error())
		}

		grpcClient = desc.NewEventServiceV1Client(con)
	})

	Describe("CreateEvent", func() {
		It("success  add today event", func() {
			var err error
			eventRes1, err = grpcClient.CreateEvent(ctx, &desc.CreateEventRequest{
				EventInfo: todayEvent,
			})
			require.NoError(GinkgoT(), err)
		})

		It("success  add week ago event", func() {
			var err error
			eventRes2, err = grpcClient.CreateEvent(ctx, &desc.CreateEventRequest{
				EventInfo: weekAgoEvent,
			})
			require.NoError(GinkgoT(), err)
		})

		It("success  add month ago event", func() {
			var err error
			eventRes3, err = grpcClient.CreateEvent(ctx, &desc.CreateEventRequest{
				EventInfo: monthAgoEvent,
			})
			require.NoError(GinkgoT(), err)
		})

		It("no required field", func() {
			_, err := grpcClient.CreateEvent(ctx, &desc.CreateEventRequest{
				EventInfo: failedEventInfo,
			})
			require.Error(GinkgoT(), err)
		})
	})

	Describe("GetEventListForDay", func() {
		It("success result", func() {
			events, err := grpcClient.GetEventListForDay(ctx, &desc.GetEventListForDayRequest{
				Date: timestamppb.New(now),
			})
			require.NoError(GinkgoT(), err)
			require.Equal(GinkgoT(), 1, len(events.GetResult().GetEvents()))
		})
	})

	Describe("GetEventListForWeek", func() {
		It("success result", func() {
			events, err := grpcClient.GetEventListForWeek(ctx, &desc.GetEventListForWeekRequest{
				WeekStart: timestamppb.New(weekAgo),
			})
			require.NoError(GinkgoT(), err)
			require.Equal(GinkgoT(), 2, len(events.GetResult().GetEvents()))
		})
	})

	Describe("GetEventListForMonth", func() {
		It("success result", func() {
			events, err := grpcClient.GetEventListForMonth(ctx, &desc.GetEventListForMonthRequest{
				MonthStart: timestamppb.New(monthAgo),
			})
			require.NoError(GinkgoT(), err)
			require.Equal(GinkgoT(), 3, len(events.GetResult().GetEvents()))
		})
	})

	Describe("UpdateEvent", func() {
		It("success result", func() {
			_, err := grpcClient.UpdateEvent(ctx, &desc.UpdateEventRequest{
				Id: &wrapperspb.Int64Value{Value: eventRes1.GetResult().GetId()},
				UpdateEventInfo: &desc.UpdateEventRequest_UpdateEventInfo{
					StartDate: timestamppb.New(now.AddDate(0, 0, 1)),
				},
			})
			require.NoError(GinkgoT(), err)

			events, err := grpcClient.GetEventListForDay(ctx, &desc.GetEventListForDayRequest{
				Date: timestamppb.New(now),
			})
			require.NoError(GinkgoT(), err)
			require.Equal(GinkgoT(), 0, len(events.GetResult().GetEvents()))
		})
	})

	Describe("DeleteEvent", func() {
		It("success delete today event", func() {
			_, err := grpcClient.DeleteEvent(ctx, &desc.DeleteEventRequest{
				Id: &wrapperspb.Int64Value{Value: eventRes1.GetResult().GetId()},
			})
			require.NoError(GinkgoT(), err)
		})

		It("success delete week ago event", func() {
			_, err := grpcClient.DeleteEvent(ctx, &desc.DeleteEventRequest{
				Id: &wrapperspb.Int64Value{Value: eventRes2.GetResult().GetId()},
			})
			require.NoError(GinkgoT(), err)
		})

		It("success delete month ago event", func() {
			_, err := grpcClient.DeleteEvent(ctx, &desc.DeleteEventRequest{
				Id: &wrapperspb.Int64Value{Value: eventRes3.GetResult().GetId()},
			})
			require.NoError(GinkgoT(), err)
		})
	})
})
