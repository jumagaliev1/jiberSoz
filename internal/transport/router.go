package transport

import "github.com/labstack/echo/v4"

func (s *Server) SetupRoutes() *echo.Group {
	v1 := s.App.Group("/api")

	return v1
}

