package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
)

const RequestIDContextKey string = "requestID"

//RequestIDMiddleware generates unique Request ID and adds it to request context and response headers
func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    requestID := uuid.New().String()

    ctx := context.WithValue(r.Context(), RequestIDContextKey, requestID)
    r = r.WithContext(ctx)

    w.Header().Set("X-Request-ID", requestID)

    log.Printf("[%s] Incoming request: %s %s", requestID, r.Method, r.URL)

    next.ServeHTTP(w, r)
	})
}
