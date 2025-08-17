package server

import (
	"html/template"
	"net/http"
	"time"
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
	_ = r.Context().Value("claims").(*JWTClaims)
	renderTemplate(w, "index", nil)
}

func (uh *UIHandlers) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: "",
		Expires: time.Now(),
		Path: "/",
	})

	http.Redirect(w, r, "/login", http.StatusFound)
}
