package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct{
	*echo.Echo
}

func (s *Server) Run() {
	s.Use(middleware.Logger())
	
}