package server

import (
	"fmt"
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
				w.WriteHeader(401)
				WriteJson(w, JSON{"error": "Unauthorised access 1"})
				return
			}

			tokenString := cookies[0].Value
			_, err := ParseJWT(tokenString)
			if err != nil {
				logger.Log.Error(err)
				w.WriteHeader(401)
				WriteJson(w, JSON{"error": "Unauthorised access 2"})
				return
			}

			next(w, r)
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
					w.WriteHeader(401)
					WriteJson(w, JSON{"error": "Unauthorised access"})
					return
				}

				tokenString := cookies[0].Value
				claims, err := ParseJWT(tokenString)
				if err != nil {
					logger.Log.Error(err)
					w.WriteHeader(401)
					WriteJson(w, JSON{"error": "Unauthorised access"})
					return
				}

				for _, requiredRole := range roles {
					if !slices.Contains(claims.Roles, requiredRole) {
						w.WriteHeader(401)
						WriteJson(w, JSON{"error": fmt.Sprintf("%s access is required for this feature", requiredRole)})
						return
					}
				}

				next(w, r)
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
