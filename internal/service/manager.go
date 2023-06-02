package service

import (
	"errors"
	"github.com/jumagaliev1/jiberSoz/internal/storage"
)

type Service struct {
	Text ITextService
}

func New(repo *storage.Repository) (*Service, error) {
	if repo == nil {
		return nil, errors.New("NO repo")
	}

	textService := NewTextService(repo)

	return &Service{
		Text: textService,
	}, nil
}
