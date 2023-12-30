package service

import (
	"context"
	"github.com/habibeh92/request-counter/config"
	"github.com/habibeh92/request-counter/internal/repository/mock"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {

	cfg := config.New()

	tests := []struct {
		name  string
		count int
	}{
		{
			name:  "single",
			count: 1,
		},
		{
			name:  "multi",
			count: 10,
		},
		// more test cases
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			counter, err := NewRequestsCounter(cfg, mock.New())
			if err != nil {
				t.Fatal(err)
			}

			for i := 0; i < test.count; i++ {
				counter.Add()
			}

			count := counter.Get(context.Background())
			if count != test.count {
				t.Errorf("Counter count is not equal to expected, expected: %d, actual: %d", test.count, count)
			}
		})
	}

}

func TestCleanUpMemory(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.New()

	tests := []struct {
		name     string
		count    int
		expected int
		wait     int
	}{
		{
			name:     "test clean up with 3 requests",
			count:    3,
			expected: 2,
			wait:     1,
		},
		// more test cases
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			counter, err := NewRequestsCounter(cfg, mock.New())
			if err != nil {
				t.Fatal(err)
			}

			cfg.RequestsTimeLimit = test.count

			for i := 0; i < test.count; i++ {
				counter.Add()
				time.Sleep(time.Duration(test.wait) * time.Second)
			}

			err = counter.cleanUpMemory()
			if err != nil {
				t.Fatal(err)
			}

			count := counter.Get(ctx)
			if count != test.expected {
				t.Errorf("Counter count is not equal to expected, expected: %d, actual: %d", test.expected, count)
			}
		})
	}

}
