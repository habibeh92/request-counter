package rest

import (
	"context"
	"encoding/json"
	"github.com/habibeh92/request-counter/config"
	"github.com/habibeh92/request-counter/internal/repository/mock"
	"github.com/habibeh92/request-counter/internal/service"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCountHandler(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.New()
	limiter := service.NewRateLimiter(cfg.RateLimit)

	counter, err := service.NewRequestsCounter(cfg, mock.New())
	if err != nil {
		log.Fatal(err)
	}

	go counter.CleanUp(ctx)

	server := New(cfg, counter, limiter)
	handler := server.counterMiddleware(http.HandlerFunc(server.countHandler))

	for i := 0; i < 5; i++ {
		req, err := http.NewRequest("GET", "/count", nil)
		if err != nil {
			t.Fatal(err)
		}

		r := httptest.NewRecorder()

		handler.ServeHTTP(r, req)

		if r.Code != http.StatusOK {
			t.Errorf("Wrong status code, actual: %v expected: %v", r.Code, http.StatusOK)
		}

		value, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
		}

		var response Response
		err = json.Unmarshal(value, &response)
		if err != nil {
			t.Error(err)
		}
		count := int(response.Result.(map[string]interface{})["count"].(float64))
		if count != i+1 {
			t.Errorf("Wrong results, actual: %v expected: %v", count, i+1)
		}
	}
}
