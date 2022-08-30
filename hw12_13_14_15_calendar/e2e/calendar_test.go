package e2e_test

import (
	"context"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
	. "github.com/onsi/ginkgo/v2"
	//. "github.com/onsi/gomega"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const grpcHost = "localhost:7002"

var _ = Describe("Calendar", func() {
	var eventInfo, failedEventInfo *desc.EventInfo
	var grpcClient desc.EventServiceV1Client
	ctx := context.Background()

	BeforeEach(func() {
		eventInfo = &desc.EventInfo{
			Title:       gofakeit.JobTitle(),
			StartDate:   timestamppb.New(gofakeit.Date()),
			Description: &wrapperspb.StringValue{Value: gofakeit.Phrase()},
			OwnerId:     gofakeit.Int64(),
		}

		failedEventInfo = &desc.EventInfo{
			Title:   gofakeit.JobTitle(),
			OwnerId: gofakeit.Int64(),
		}

		con, err := grpc.Dial(grpcHost, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("didn't connect: %s", err.Error())
		}

		grpcClient = desc.NewEventServiceV1Client(con)
	})

	Describe("CreateEvent", func() {
		It("success result", func() {
			_, err := grpcClient.CreateEvent(ctx, &desc.CreateEventRequest{
				EventInfo: eventInfo,
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

})
