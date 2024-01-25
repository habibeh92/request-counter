package rest

import (
	"encoding/json"
	"net/http"
	"time"
)

// SetupRoutes configures the HTTP routes for the server.
func (s *Server) SetupRoutes(mux *http.ServeMux) {
	mux.Handle("/count", s.counterMiddleware(http.HandlerFunc(s.countHandler)))
}

// countHandler handles HTTP requests to retrieve the count of the requests in last 60 seconds.
func (s *Server) countHandler(w http.ResponseWriter, r *http.Request) {
	if !s.limiter.Allow() {
		s.responseJson(w, Response{
			Status: http.StatusTooManyRequests,
			Result: "Rate limit exceeded",
		}, http.StatusTooManyRequests)
		return
	}
	defer s.limiter.Release()

	count := s.counter.Get(r.Context())

	// simulate 2 second process
	time.Sleep(2 * time.Second)

	response := Response{
		Status: http.StatusOK,
		Result: ResponseCount{
			Count: count,
		},
	}

	s.responseJson(w, response, http.StatusOK)
}

// responseJson sends a JSON response with the specified status code.
func (s *Server) responseJson(w http.ResponseWriter, response Response, status int) {
	resp, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(resp)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
