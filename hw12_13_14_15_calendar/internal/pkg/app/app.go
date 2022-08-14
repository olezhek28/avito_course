package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	eventV1 "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/api/event_v1"
	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
	"google.golang.org/grpc"
)

const (
	httpPort = ":7000"
	grpcPort = ":7002"
)

type App struct {
	eventImpl       *eventV1.Implementation
	serviceProvider *serviceProvider

	pathConfig string

	grpcServer *grpc.Server
	mux        *runtime.ServeMux
}

// NewApp ...
func NewApp(ctx context.Context, pathConfig string) (*App, error) {
	a := &App{
		pathConfig: pathConfig,
	}
	err := a.initDeps(ctx)

	return a, err
}

func (a *App) Run() error {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	err := a.runGRPC(wg)
	if err != nil {
		return err
	}

	err = a.runPublicHTTP(wg)
	if err != nil {
		return err
	}

	wg.Wait()
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initServer,
		a.initGRPCServer,
		a.initPublicHTTPHandlers,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.pathConfig)
	return nil
}

func (a *App) initServer(ctx context.Context) error {
	a.eventImpl = eventV1.NewEventV1(
		a.serviceProvider.GetNoteService(ctx),
	)

	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.StreamInterceptor(grpcValidator.StreamServerInterceptor()),
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)

	desc.RegisterEventServiceV1Server(a.grpcServer, a.eventImpl)

	return nil
}

func (a *App) initPublicHTTPHandlers(ctx context.Context) error {
	a.mux = runtime.NewServeMux()

	// nolint
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterEventServiceV1HandlerFromEndpoint(ctx, a.mux, grpcPort, opts)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runGRPC(wg *sync.WaitGroup) error {
	list, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return err
	}

	go func() {
		defer wg.Done()

		if err = a.grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to process gRPC server: %s", err.Error())
		}
	}()

	log.Printf("Run gRPC server on %s port\n", grpcPort)
	return nil
}

func (a *App) runPublicHTTP(wg *sync.WaitGroup) error {
	go func() {
		defer wg.Done()

		if err := http.ListenAndServe(httpPort, a.mux); err != nil {
			log.Fatalf("failed to process muxer: %s", err.Error())
		}
	}()

	log.Printf("Run public http handler on %s port\n", httpPort)
	return nil
}
