package server

import (
	"github.com/balu6914/KYC-Match-API/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server defines the interface for the server operations
type Server interface {
	Start() error
}

// EchoServer implements the Server interface using Echo framework
type EchoServer struct {
	e       *echo.Echo
	handler handlers.KYCHandler // Changed from *handlers.KYCHandler to handlers.KYCHandler
}

// NewEchoServer creates a new instance of EchoServer
func NewEchoServer(handler handlers.KYCHandler) *EchoServer {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &EchoServer{
		e:       e,
		handler: handler,
	}
}

func (s *EchoServer) Start() error {
	s.e.POST("/match", s.handler.Match)
	return s.e.Start(":8080")
}
