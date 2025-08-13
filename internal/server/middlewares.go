package server

import "net/http"

type Middlewere func (http.Handler) http.Handler

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func (w http.ResponseWriter, r *http.Request) {
			// Do authentication check here

			next.ServeHTTP(w, r)
		},
	)
}

func Role(roles []string) Middlewere {
	return func (next http.Handler) http.Handler {
		return http.HandlerFunc(
			func (w http.ResponseWriter, r *http.Request) {
				// Do admin role check here

				next.ServeHTTP(w, r)
			},
		)
	}
}

func Chain(h http.Handler, middleweres ...Middlewere) http.Handler {
	for _, m := range middleweres {
		h = m(h)
	}

	return h
}
