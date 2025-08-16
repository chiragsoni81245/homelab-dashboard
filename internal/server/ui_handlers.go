package server

import (
	"homelab-dashboard/internal/logger"
	"html/template"
	"net/http"
)

// Pre template render for caching
var templates = template.Must(template.ParseFS(templateFS, "templates/*.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// --------- UI Handlers -------------

type UIHandlers struct{}

func (uh *UIHandlers) LoginHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login", nil)
}

func (uh *UIHandlers) DashboardHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(*JWTClaims)
	logger.Log.Info(claims)
	renderTemplate(w, "index", nil)
}
