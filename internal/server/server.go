package server

import (
	"github.com/labstack/echo/v4"

	"osamikoin/internal/api"
)

type Server struct {
	*echo.Echo
}
func New() Server {
	return Server{echo.New()}
}
func (s *Server) Run() {
	s.GET("/", api.Home)

	s.POST("/send", api.Send)
	s.Logger.Panic(s.Start(":2020"))
}
