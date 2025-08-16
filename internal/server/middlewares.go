package server

import (
	"context"
	"homelab-dashboard/internal/logger"
	"net/http"
	"slices"
)

type Middlewere func (http.HandlerFunc) http.HandlerFunc

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func (w http.ResponseWriter, r *http.Request) {
			// Do authentication check here
			cookies := r.CookiesNamed("token")
			if len(cookies) == 0 {
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}

			tokenString := cookies[0].Value
			claims, err := ParseJWT(tokenString)
			if err != nil {
				logger.Log.Error(err)
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}
			
			ctx := context.WithValue(r.Context(), "claims", claims)
			next(w, r.WithContext(ctx))
		},
	)
}

func Role(roles []string) Middlewere {
	return func (next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(
			func (w http.ResponseWriter, r *http.Request) {
				// Do admin role check here
				cookies := r.CookiesNamed("token")
				if len(cookies) == 0 {
					http.Redirect(w, r, "/login", http.StatusFound)
					return
				}

				tokenString := cookies[0].Value
				claims, err := ParseJWT(tokenString)
				if err != nil {
					logger.Log.Error(err)
					http.Redirect(w, r, "/login", http.StatusFound)
					return
				}

				for _, requiredRole := range roles {
					if !slices.Contains(claims.Roles, requiredRole) {
						http.Redirect(w, r, "/login", http.StatusFound)
						return
					}
				}

				ctx := context.WithValue(r.Context(), "claims", claims)
				next(w, r.WithContext(ctx))
			},
		)
	}
}

func Chain(h http.HandlerFunc, middleweres ...Middlewere) http.HandlerFunc {
	for _, m := range middleweres {
		h = m(h)
	}

	return h
}
