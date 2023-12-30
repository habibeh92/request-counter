package service

import (
	"context"
	"github.com/habibeh92/request-counter/config"
	"github.com/habibeh92/request-counter/internal/repository"
	"log"
	"strconv"
	"sync"
	"time"
)

type RequestCounter struct {
	cfg         *config.Config
	repo        repository.Request
	requestData []string

	mu sync.RWMutex
}

// NewRequestsCounter get new instance of the RequestCounter
func NewRequestsCounter(cfg *config.Config, repo repository.Request) (*RequestCounter, error) {
	c := &RequestCounter{cfg: cfg, repo: repo}
	err := c.loadRequestData()

	return c, err
}

// Add a function to add the current time unix to the counter storage
func (c *RequestCounter) Add() {
	c.mu.Lock()
	defer c.mu.Unlock()

	t := strconv.Itoa(int(time.Now().Unix()))

	c.requestData = append(c.requestData, t)
}

// Get fetch the count of requests (in the configured timeframe)
func (c *RequestCounter) Get(ctx context.Context) int {
	return len(c.requestData)
}

// CleanUp remove the expired requests
func (c *RequestCounter) CleanUp(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(100 * time.Millisecond):
			oldLen := len(c.requestData)
			if oldLen == 0 {
				continue
			}
			err := c.cleanUpMemory()
			if err != nil {
				log.Println(err)
				continue
			}

			if oldLen == len(c.requestData) {
				continue
			}

			err = c.repo.Sync(c.requestData)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
}

// cleanUpMemory remove the expired requests from the memory storage
func (c *RequestCounter) cleanUpMemory() error {
	for k, value := range c.requestData {
		unix, err := strconv.Atoi(value)
		if err != nil {
			continue
		}
		unixTimestamp := int64(unix)
		if time.Since(time.Unix(unixTimestamp, 0)) >= time.Duration(c.cfg.RequestsTimeLimit)*time.Second {
			if k > len(c.requestData)-1 {
				continue
			}
			c.mu.Lock()

			if k == len(c.requestData)-1 {
				c.requestData = c.requestData[:k]
				c.mu.Unlock()
				continue
			}

			c.requestData = append(c.requestData[:k], c.requestData[k+1:]...)
			c.mu.Unlock()
		}
	}

	return nil
}

// loadRequestData load the requests from the repo (file) to memory storage
func (c *RequestCounter) loadRequestData() error {
	var err error
	c.requestData, err = c.repo.All()

	return err
}
