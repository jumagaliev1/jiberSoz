package transport

import "github.com/labstack/echo/v4"

func (s *Server) SetupRoutes() *echo.Group {
	v1 := s.App.Group("/api")

	v1.POST("/create", s.Handler.Text.Create)
	v1.GET("/text/:link", s.Handler.Text.GetByLink)

	return v1
}
