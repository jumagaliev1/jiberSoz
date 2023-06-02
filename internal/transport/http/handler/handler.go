package handler

import (
	"errors"
	"github.com/jumagaliev1/jiberSoz/internal/service"
)

type Handler struct {
}

func New(service *service.Service) (*Handler, error) {
	if service == nil {
		return nil, errors.New("no given service")
	}

	//genHandler := NewGeneratorHandler(service)

	return &Handler{}, nil

}
