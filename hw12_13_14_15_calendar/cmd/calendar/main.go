package main

import (
	"context"
	"flag"
	"log"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/pkg/app"
)

var pathConfig string

func init() {
	flag.StringVar(&pathConfig, "config", "./config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	ctx := context.Background()

	a, err := app.NewApp(ctx, pathConfig)
	if err != nil {
		log.Fatalf("Can't create app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("Can't run app: %s", err.Error())
	}
}

//func main() {
//	flag.Parse()
//
//	if flag.Arg(0) == "version" {
//		printVersion()
//		return
//	}
//
//	config := NewConfig()
//	logg := logger.New(config.Logger.Level)
//
//	storage := memorystorage.New()
//	calendar := app.New(logg, storage)
//
//	server := internalhttp.NewServer(logg, calendar)
//
//	ctx, cancel := signal.NotifyContext(context.Background(),
//		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
//	defer cancel()
//
//	go func() {
//		<-ctx.Done()
//
//		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
//		defer cancel()
//
//		if err := server.Stop(ctx); err != nil {
//			logg.Error("failed to stop http server: " + err.Error())
//		}
//	}()
//
//	logg.Info("calendar is running...")
//
//	if err := server.Start(ctx); err != nil {
//		logg.Error("failed to start http server: " + err.Error())
//		cancel()
//		os.Exit(1) //nolint:gocritic
//	}
//}
