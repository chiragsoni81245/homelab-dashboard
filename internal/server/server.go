// Its a package containing http server build using standard net/http
package server

import (
	"fmt"
	"homelab-dashboard/internal/config"
	"homelab-dashboard/internal/logger"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type Server struct {
	Router  http.Handler	
	Address string
}

type JSON map[string]any

var JWT_SECRET []byte = []byte(config.App.Server.SecretKey) 

func JWT_SECRET_FUNC (*jwt.Token) (any, error) {
	return JWT_SECRET, nil
}

type JWTClaims struct {
	UserId   uint   `json:"user_id"`
	Username string `json:"username"`	
	Roles    []string `json:"roles"`
	jwt.RegisteredClaims
}


func NewServer() *Server {

	router := http.NewServeMux()

	// Static file server to server files from internal/server/static directory
	staticHandler := http.FileServer(http.FS(staticFS))
	router.Handle("GET /static/", staticHandler)

	// Auth API routes
	authAPI := AuthAPIHandlers{}
	apiBasePath := "/api/v1"
	router.HandleFunc(fmt.Sprintf("POST %s/auth/login", apiBasePath), authAPI.LoginHandler)

	// UI routes
	ui := UIHandlers{}
	router.HandleFunc("GET /login", ui.LoginHandler)
	router.HandleFunc("GET /{$}", Chain(ui.DashboardHandler, AuthMiddleware))



	return &Server{
		Router: router,
		Address: fmt.Sprintf("0.0.0.0:%d", config.App.Server.Port),
	} 
}

func (s *Server) Start() error {
	logger.Log.Info(fmt.Sprintf("Listening on %s", s.Address))
	return http.ListenAndServe(s.Address, s.Router)
}
