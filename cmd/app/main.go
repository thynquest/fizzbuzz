package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/thynquest/fizzbuzz/pkg/logging"
	"github.com/thynquest/fizzbuzz/server"
)

var mainTitle = "[fizzbuzz-api-main]"

const (
	envFile = "ENV_FILE"
)

func main() {
	envFile := os.Getenv(envFile)
	config, errConfig := server.LoadConfig(envFile)
	if errConfig != nil {
		panic(errConfig)
	}
	target := fmt.Sprintf("%s:%s", config.HOST, config.PORT)
	logging.Info(mainTitle, fmt.Sprintf("fizzbuzz server on %s", target))
	ctx := context.Background()
	APIServer := server.NewServer(target)
	server.DefineHandlers(APIServer)
	stop := make(chan error, 1)
	go APIServer.Start(stop)
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)
		<-sig
		stop <- fmt.Errorf("received Interrupt signal")
	}()
	<-stop
	logging.Info(mainTitle, "shutting down..")
	APIServer.ShutDown(ctx)
}
