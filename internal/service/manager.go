package service

import (
	"errors"
	"github.com/jumagaliev1/jiberSoz/internal/storage"
	s32 "github.com/jumagaliev1/jiberSoz/internal/storage/s3"
)

type Service struct {
	Text ITextService
}

func New(repo *storage.Repository) (*Service, error) {
	if repo == nil {
		return nil, errors.New("NO repo")
	}
	s3 := s32.NewAmazonS3()

	textService := NewTextService(repo, s3)

	return &Service{
		Text: textService,
	}, nil
}
