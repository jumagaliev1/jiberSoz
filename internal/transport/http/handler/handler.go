package handler

import (
	"errors"
	"github.com/jumagaliev1/jiberSoz/internal/service"
	"github.com/labstack/echo/v4"
)

type ITextHandler interface {
	Create(c echo.Context) error
	GetByLink(c echo.Context) error
}

type Handler struct {
	Text ITextHandler
}

func New(service *service.Service) (*Handler, error) {
	if service == nil {
		return nil, errors.New("no given service")
	}

	textHandler := NewTextHandler(service)

	return &Handler{
		Text: textHandler,
	}, nil

}
