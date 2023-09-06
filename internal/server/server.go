package server

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"

	"testTask/internal/config"
)

type Server struct {
	e   *echo.Echo
	cfg *config.Config
}

func New(_ context.Context, cfg *config.Config) (*Server, error) {
	e := echo.New()
	handler := NewHandler()

	api := e.Group("/api")

	// Auth API routes
	api.GET("/items/limits", handler.GetLimits)
	api.POST("/items/process", handler.Process)

	return &Server{e: e, cfg: cfg}, nil
}

func (s *Server) Start() error {
	port := s.cfg.Port
	slog.Info("server started", "port", port)
	return s.e.Start(fmt.Sprintf(":%d", port))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}

func (s *Server) Port() (int, error) {
	listener := s.e.Listener
	if listener == nil {
		return 0, errors.New("server is not started")
	}

	addr := listener.Addr()
	if addr == nil {
		return 0, errors.New("server is not started")
	}

	return addr.(*net.TCPAddr).Port, nil
}
