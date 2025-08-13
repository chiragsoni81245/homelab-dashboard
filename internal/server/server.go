package server

import (
	"fmt"
	"homelab-dashboard/internal/config"
	"net/http"
)

type Server struct {
	Router  http.Handler	
	Address string
}

func NewServer() *Server {

	router := http.NewServeMux()

	// Static file server to server files from internal/static directory
	staticHandler := http.FileServer(http.Dir("internal/static"))
	router.Handle("GET /static/", staticHandler)

	// UI routes
	router.Handle("GET /", DashboardHandler{})

	// API routes
	router.Handle("GET /api/auth/", AuthHandler{})


	return &Server{
		Router: router,
		Address: fmt.Sprintf("0.0.0.0:%d", config.App.Server.Port),
	} 
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.Address, s.Router)
}
