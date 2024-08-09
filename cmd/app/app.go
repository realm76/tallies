package main

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"tallies/internal/app"
	"time"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)

	defer cancel()
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	appServer := app.NewServer(sugar)
	httpServer := &http.Server{
		Addr:    net.JoinHostPort("", "3000"),
		Handler: appServer,
	}

	go func() {
		log.Printf("server listening on %s", httpServer.Addr)

		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("server error: %v", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		<-ctx.Done()

		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)

		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.Printf("server shutdown error: %v", err)
		}
	}()

	wg.Wait()

	return nil
}
