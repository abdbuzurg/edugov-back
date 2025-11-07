package middleware

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// CORSConfig defines the configuration for the CORS middleware.
type CORSConfig struct {
	AllowedOrigins   []string      // List of allowed origins (e.g., ["http://localhost:3000", "https://yourfrontend.com"])
	AllowedMethods   []string      // List of allowed HTTP methods (e.g., ["GET", "POST", "PUT", "DELETE", "OPTIONS"])
	AllowedHeaders   []string      // List of allowed request headers (e.g., ["Content-Type", "Authorization"])
	ExposedHeaders   []string      // List of headers that can be exposed to the browser
	AllowCredentials bool          // Whether to allow credentials (cookies, HTTP auth). If true, AllowedOrigins cannot contain "*"
	MaxAge           time.Duration // How long the results of a preflight request can be cached (e.g., 6 hours)
}

// DefaultCORSConfig provides a sensible default configuration
// for development, allowing common methods and headers from all origins,
// but it's crucial to narrow down AllowedOrigins in production.
var DefaultCORSConfig = CORSConfig{
	AllowedOrigins:   []string{"http://10.154.20.112", "http://farzonagon.tj", "*"}, // WARNING: Use specific origins in production!
	AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization"},
	ExposedHeaders:   []string{"Link"}, // Example: If you send 'Link' header for pagination
	AllowCredentials: true,
	MaxAge:           time.Hour * 6, // Cache preflight for 6 hours
}

// CORSMiddleware creates a middleware that applies CORS policies.
// It takes a CORSConfig struct to define the specific rules.
func CORSMiddleware(config CORSConfig) func(http.Handler) http.Handler {
	// Pre-join headers/methods into comma-separated strings for efficiency,
	// as these strings will be set on every response.
	allowedMethods := strings.Join(config.AllowedMethods, ", ")
	allowedHeaders := strings.Join(config.AllowedHeaders, ", ")
	exposedHeaders := strings.Join(config.ExposedHeaders, ", ")
	// Convert MaxAge duration to a string representing seconds, as required by the header.
	maxAge := strconv.FormatInt(int64(config.MaxAge.Seconds()), 10)

	// Return the actual middleware function, which takes an http.Handler and returns an http.Handler.
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the 'Origin' header from the incoming request. This header indicates
			// the origin of the request (scheme, host, and port).
			origin := r.Header.Get("Origin")

			// Determine if the request's origin is allowed based on the CORS configuration.
			isOriginAllowed := false
			if origin != "" { // Only proceed if an Origin header is actually present.
				if len(config.AllowedOrigins) == 1 && config.AllowedOrigins[0] == "*" {
					// If the configuration explicitly allows all origins ("*").
					isOriginAllowed = true
				} else {
					// Iterate through the list of specifically allowed origins.
					for _, allowedOrigin := range config.AllowedOrigins {
						if origin == allowedOrigin {
							isOriginAllowed = true
							break // Found a match, no need to check further.
						}
					}
				}
			}

			// Set the 'Access-Control-Allow-Origin' header.
			// This header indicates whether the response can be shared with the requesting origin.
			if isOriginAllowed {
				if config.AllowCredentials && (len(config.AllowedOrigins) == 1 && config.AllowedOrigins[0] == "*") {
					// If credentials (cookies, HTTP auth) are allowed, 'Access-Control-Allow-Origin'
					// cannot be "*". It must be the specific origin from the request.
					w.Header().Set("Access-Control-Allow-Origin", origin)
				} else if len(config.AllowedOrigins) == 1 && config.AllowedOrigins[0] == "*" {
					// If no credentials are allowed, and "*" is configured, then "*" is fine.
					w.Header().Set("Access-Control-Allow-Origin", "*")
				} else {
					// If specific origins are configured, set it to the matched origin.
					w.Header().Set("Access-Control-Allow-Origin", origin)
				}
			} else if origin != "" {
				// If the origin is not allowed and an Origin header was present,
				// do not set any Access-Control-Allow-Origin header.
				// You might log this for debugging.
				log.Printf("CORS: Request from disallowed origin: %s", origin)
				// Optionally, you could explicitly block the request here with http.StatusForbidden
				// if you want stricter control beyond just CORS headers.
				// http.Error(w, "Forbidden", http.StatusForbidden)
				// return
			}

			// Set other common CORS headers for both preflight and actual requests.
			// These tell the browser what methods and headers are allowed for the actual request.
			w.Header().Set("Access-Control-Allow-Methods", allowedMethods)
			w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)

			// If credentials are allowed, set the 'Access-Control-Allow-Credentials' header.
			if config.AllowCredentials {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}
			// If there are headers to be exposed to the client, set 'Access-Control-Expose-Headers'.
			if exposedHeaders != "" {
				w.Header().Set("Access-Control-Expose-Headers", exposedHeaders)
			}

			// Handle CORS preflight requests (HTTP OPTIONS method).
			// Browsers send these requests to check if the actual request is safe to send.
			if r.Method == "OPTIONS" {
				// Set 'Access-Control-Max-Age' to indicate how long the preflight results can be cached.
				w.Header().Set("Access-Control-Max-Age", maxAge)
				// Respond with 200 OK for a successful preflight.
				w.WriteHeader(http.StatusOK)
				return // Terminate the request processing here, as preflight is complete.
			}

			// For all other HTTP methods (GET, POST, PUT, etc.),
			// pass the request to the next handler in the chain.
			next.ServeHTTP(w, r)
		})
	}
}
