package rest

import (
	"context"
	"fmt"
	"github.com/habibeh92/request-counter/config"
	"github.com/habibeh92/request-counter/internal/service"
	"log"
	"net/http"
	"time"
)

type Server struct {
	cfg     *config.Config
	counter *service.RequestCounter
	limiter *service.RateLimiter
}

// New creates a new instance of the Server with the provided configuration and request counter.
func New(cfg *config.Config, counter *service.RequestCounter, limiter *service.RateLimiter) *Server {
	return &Server{cfg: cfg, counter: counter, limiter: limiter}
}

// Serve starts the HTTP server and listens for incoming requests.
func (s *Server) Serve(ctx context.Context) error {
	mux := http.NewServeMux()

	// Configure the routes for the server.
	s.SetupRoutes(mux)

	fmt.Println(fmt.Sprintf("Server is listening to port: %s", s.cfg.Http.Port))

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", s.cfg.Http.Port),
		Handler:           mux,
		ReadHeaderTimeout: time.Second * 5,
	}

	srvError := make(chan error)
	go func() {
		srvError <- srv.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		log.Println("rest server is shutting down...")
		return srv.Shutdown(ctx)
	case err := <-srvError:
		return err
	}
}
