package middleware

import (
	"context"
	"net/http"
	"strings"
)

const LanguageContextKey string = "language"

func LanguageMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		acceptLanguageHeader := r.Header.Get("Accept-Language")

		lang := strings.ToLower(strings.TrimSpace(acceptLanguageHeader))

		if lang != "en" && lang != "ru" && lang != "tg" {
			lang = "tg"
		}

		ctx := context.WithValue(r.Context(), LanguageContextKey, lang)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// Get Language from Context, if absent set to "tg"
func GetLanguageFromContext(ctx context.Context) string {
	language, ok := ctx.Value(LanguageContextKey).(string)
	if !ok || language == "" {
		return "tg"
	}

	return language
}
