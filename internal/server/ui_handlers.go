package server

import (
	"html/template"
	"net/http"
)

// Pre template render for caching
var templateFiles, _ = GetTemplateFile("internal/server/static") 
var templates = template.Must(template.ParseFiles(templateFiles...))

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	w.WriteHeader(200)
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

type DashboardHandler struct {}

func (dh DashboardHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", nil)
}
