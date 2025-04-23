package server

import (
	"net/http"

	"github.com/balu6914/KYC-Match-API/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server interface {
	Start() error
}

type EchoServer struct {
	e       *echo.Echo
	handler handlers.KYCHandler
}

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
	// Add root route
	s.e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Welcome to KYC Match API!",
		})
	})

	// Existing KYC match route
	s.e.POST("/match", s.handler.Match)

	// Handle favicon.ico
	s.e.GET("/favicon.ico", func(c echo.Context) error {
		return c.NoContent(http.StatusNoContent)
	})

	return s.e.Start(":8080")
}
