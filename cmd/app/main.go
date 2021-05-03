package main

import (
	"context"
	"github.com/acelot/articles/internal/app"
	"github.com/acelot/articles/internal/http/handler"
	"github.com/docopt/docopt-go"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const usage = `Application

Usage:
  app [options] <primary-dsn>

Options:
  -h --help                     Show this screen.
  -l --listen=<address>         Server listen address [default: 0.0.0.0:80].
     --secondary-db=<dsn>       Secondary PostgreSQL connection string. If omitted primary database is used.
     --image-storage=<uri>      Images S3 storage URI [default: http://minioadmin:minioadmin@localhost:9000/images].
     --log-level=<string>       Logging level [default: debug].
	 --shutdown-timeout=<secs>  Shutdown timeout in seconds [default: 15].`

func main() {
	stderr := log.New(os.Stderr, "", 0)

	args, err := parseArgs()
	if err != nil {
		stderr.Fatalf("cannot parse args: %v", err)
	}

	env, err := app.NewEnv(args)
	if err != nil {
		stderr.Fatalf("cannot create environment: %v", err)
	}

	env.Logger.Info("app started")

	server := http.Server{
		Addr:    args.ListenAddress,
		Handler: handler.NewAppHandler(env),
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			env.Logger.Fatal("app stopped due error", zap.Error(err))
		}

		env.Logger.Info("app stopped gracefully")
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interrupt

	env.Logger.Warn("app interruption signal received")

	ctx, cancel := context.WithTimeout(context.Background(), args.ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		env.Logger.Fatal("app shutdown failed", zap.Error(err))
	}

	if err := env.Close(); err != nil {
		env.Logger.Fatal("app environment closing failed", zap.Error(err))
	}
}

func parseArgs() (*app.Args, error) {
	opts, err := docopt.ParseDoc(usage)
	if err != nil {
		return nil, err
	}

	args, err := app.NewArgs(&opts)
	if err != nil {
		return nil, err
	}

	return args, nil
}
