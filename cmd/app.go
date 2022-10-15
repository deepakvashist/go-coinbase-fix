package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/deepakvashist/go-coinbase-fix/client"
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

	fixLogFactory := quickfix.NewScreenLogFactory()
	fixConfigFile, _ := os.Open("config/client.cfg")
	fixSettings, _ := quickfix.ParseSettings(fixConfigFile)

	app := client.NewClient()

	fixAcceptor, err := quickfix.NewInitiator(
		app,
		quickfix.NewMemoryStoreFactory(),
		fixSettings,
		fixLogFactory,
	)
	if err != nil {
		zap.L().Error("quickfix initiator error", zap.Error(err))
		return
	}

	err = fixAcceptor.Start()
	if err != nil {
		zap.L().Error("quickfix initiator start error", zap.Error(err))
		return
	}

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
			zap.L().Info("closing fix initiator")
			fixAcceptor.Stop()
			return
		default:
		}
	}
}
