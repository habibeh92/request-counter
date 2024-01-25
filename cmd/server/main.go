package main

import (
	"context"
	"github.com/habibeh92/request-counter/api/rest"
	"github.com/habibeh92/request-counter/config"
	repository "github.com/habibeh92/request-counter/internal/repository/file"
	"github.com/habibeh92/request-counter/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg := config.New()

	// Open or create a file for storing requests. If the file does not exist, it will be created.
	file, err := os.OpenFile("data/requests.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	counter, err := service.NewRequestsCounter(cfg, repository.New(file))
	if err != nil {
		log.Fatal(err)
	}

	limiter := service.NewRateLimiter(cfg.RateLimit)

	// Start a goroutine for periodic cleanup of old request data.
	go counter.CleanUp(ctx)

	server := rest.New(cfg, counter, limiter)

	// Start serving HTTP requests and handle any errors that may occur.
	if err = server.Serve(ctx); err != nil {
		log.Fatal(err)
	}
}
