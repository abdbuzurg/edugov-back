package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func CreateMiddlewareStack(middlewares ...Middleware) Middleware {
  return func(next http.Handler) http.Handler {
    for i := len(middlewares) - 1; i >= 0; i-- {
      middleware := middlewares[i]
      next = middleware(next)
    }

    return next
  }
}
