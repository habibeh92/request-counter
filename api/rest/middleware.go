package rest

import (
	"net/http"
)

// counterMiddleware a middleware for incrementing the requests count on each incoming HTTP request before passing it to the next handler.
func (s *Server) counterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Increment the requests count.
		s.counter.Add()

		next.ServeHTTP(w, r)
	})
}
