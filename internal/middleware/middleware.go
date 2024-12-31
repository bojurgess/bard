package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func ChainMiddleware(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := range middlewares {
		handler = middlewares[len(middlewares)-1-i](handler)
	}

	return handler
}
