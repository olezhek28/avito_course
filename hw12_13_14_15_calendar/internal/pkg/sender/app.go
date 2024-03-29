package calendar_app

import (
	"context"
	"log"
	"sync"
)

// App ...
type App struct {
	serviceProvider *serviceProvider

	pathConfig string
}

// NewApp ...
func NewApp(ctx context.Context, pathConfig string) (*App, error) {
	a := &App{
		pathConfig: pathConfig,
	}
	err := a.initDeps(ctx)

	return a, err
}

// Run ...
func (a *App) Run() error {
	defer func() {
		a.serviceProvider.rabbitConsumer.Close()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	err := a.runSenderService(wg)

	wg.Wait()
	return err
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
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

func (a *App) runSenderService(wg *sync.WaitGroup) error {
	go func() {
		defer wg.Done()

		if err := a.serviceProvider.GetSenderService().Run(); err != nil {
			log.Fatalf("failed to process sender service: %s", err.Error())
		}
	}()

	log.Printf("Run sender service ...\n")
	return nil
}
