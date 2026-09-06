package http

import (
	"booking-service/config"
	"booking-service/internal/adapters/http/handler"
	"booking-service/internal/adapters/http/middleware"
	"booking-service/internal/domain/ports"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
	engine     *gin.Engine
}

func NewServer(cfg config.API, jwtCfg config.JWT, userService ports.UserService) *Server {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	v1 := engine.Group("/api/v1")

	handler.NewAuthHandler(userService).RegisterRoutes(v1)

	users := v1.Group("/")
	users.Use(middleware.Auth(jwtCfg.Secret))
	handler.NewUserHandler(userService).RegisterRoutes(users)

	return &Server{
		engine: engine,
		httpServer: &http.Server{
			Addr:         ":" + cfg.Port,
			Handler:      engine,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
		},
	}
}

func (s *Server) Start() {
	log.Printf("HTTP server listening on %s", s.httpServer.Addr)
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("HTTP server error: %v", err)
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
