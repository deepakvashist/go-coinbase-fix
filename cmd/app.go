package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/deepakvashist/go-fix/client"
	"github.com/quickfixgo/quickfix"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	zap.ReplaceGlobals(logger)
	//nolint: errcheck // disable log sync error validation
	defer logger.Sync()

	cfg, err := os.Open("config/client.cfg")
	if err != nil {
		zap.L().Error("fix config file error", zap.Error(err))
		return
	}

	appSettings, err := quickfix.ParseSettings(cfg)
	if err != nil {
		zap.L().Error("quickfix parse settings error", zap.Error(err))
		return
	}

	fileLogFactory, err := quickfix.NewFileLogFactory(appSettings)
	if err != nil {
		zap.L().Error("quickfix new filelog factory error", zap.Error(err))
		return
	}

	app := client.NewClient()

	initiator, err := quickfix.NewInitiator(
		app,
		quickfix.NewMemoryStoreFactory(),
		appSettings,
		fileLogFactory,
	)
	if err != nil {
		zap.L().Error("quickfix initiator error", zap.Error(err))
		return
	}

	initiator.Start()

	ctx, ctxCancel := context.WithCancel(context.Background())
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		ctxCancel()
	}()
	for {
		select {
		case <-ctx.Done():
			zap.L().Info("stop fix initiator")
			initiator.Stop()
			return
		default:
		}
	}
}
