package server

import (
	"your_project/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server interface {
	Start() error
}

type EchoServer struct {
	e       *echo.Echo
	handler *handlers.KYCHandler
}

func NewEchoServer(handler *handlers.KYCHandler) *EchoServer {
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
//	