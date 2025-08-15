package server

import (
	"net/http"
)


//------------------ Authentication Handler ------------------

type AuthHandler struct {}

func (ah AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

