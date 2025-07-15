package middleware

import (
	"fmt"
	"log" // Or your structured logger (slog)
	"net/http"
	"runtime/debug" // For printing stack trace
)

func PanicRecoveryMiddleware(respondWithError func(w http.ResponseWriter, r *http.Request, err error)) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rvr := recover(); rvr != nil {
					// Log the panic with a stack trace. This is crucial for debugging.
					// Use your structured logger (slog) if available.
					log.Printf("PANIC: %v\n%s\n", rvr, debug.Stack()) // Using standard log for simplicity, replace with slog if you have it injected

					// Respond with a generic 500 Internal Server Error to the client.
					// Do NOT expose the panic message or stack trace to the client.
					// Use your utility function for standardized error responses.
          respondWithError(w, r, fmt.Errorf("Internal Server Error. Please try again later"))
				}
			}()

			// Call the next handler in the chain
			next.ServeHTTP(w, r)
		})
	}
}
