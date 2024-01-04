package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/thynquest/fizzbuzz/pkg/logging"
)

const serverTitle = "[fizzbuzz-api-server]"

type fizzbuzzServer struct {
	server *http.Server
}

func NewServer(addr string) *fizzbuzzServer {
	return &fizzbuzzServer{
		server: &http.Server{
			Addr: addr,
		},
	}
}

func (f *fizzbuzzServer) Start(errChan chan error) {
	if f.server.Addr == "" {
		logging.Error(serverTitle, "server address is missing in configuration")
		errChan <- errors.New("server address is missing in configuration")
		return
	}

	if err := f.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logging.Error(serverTitle, fmt.Sprintf("error when starting the server: %s", err.Error()))
		errChan <- err
		return
	}
}

func (f *fizzbuzzServer) ShutDown(ctx context.Context) {
	if err := f.server.Shutdown(ctx); err != nil {
		logging.Fatal(serverTitle, err.Error())
		return
	}
}
