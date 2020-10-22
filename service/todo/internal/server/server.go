package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nasjp/microservice-todo/service/todo/internal/grpc"
)

const port = 3000

func Run() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}

func run() error {
	termCh := make(chan os.Signal, 1)
	signal.Notify(termCh, syscall.SIGTERM, syscall.SIGINT)

	s := grpc.NewServer(port)
	errCh := make(chan error, 1)

	go func() {
		errCh <- s.Start()
	}()

	select {
	case <-termCh:
		if err := s.Stop(); err != nil {
			return err
		}

		return nil
	case <-errCh:
		return nil
	}
}
