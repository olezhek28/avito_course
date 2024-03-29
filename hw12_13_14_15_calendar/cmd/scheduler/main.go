package main

import (
	"context"
	"flag"
	"log"
	"time"

	schedulerApp "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/pkg/scheduler"
)

var pathConfig string

func init() {
	flag.StringVar(&pathConfig, "config", "./scheduler_config.json", "Path to configuration file")
	time.Local = time.UTC
}

func main() {
	flag.Parse()

	ctx := context.Background()

	a, err := schedulerApp.NewApp(ctx, pathConfig)
	if err != nil {
		log.Fatalf("Can't create app: %s", err.Error())
	}

	err = a.Run(ctx)
	if err != nil {
		log.Fatalf("Can't run app: %s", err.Error())
	}
}
