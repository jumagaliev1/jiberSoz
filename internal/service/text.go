package service

import (
	"context"
	"errors"
	"github.com/jumagaliev1/jiberSoz/internal/model"
	"github.com/jumagaliev1/jiberSoz/internal/storage"
	"math/rand"
	"time"
)

var (
	lengthBytes = 8
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type TextService struct {
	repo *storage.Repository
}

func NewTextService(repo *storage.Repository) *TextService {
	return &TextService{repo: repo}
}

func (s *TextService) Create(ctx context.Context, request model.TextRequest) (*model.Text, error) {
	text := request.ToText()
	text.CreatedAt = time.Now()
	text.ExpiresAt = text.CreatedAt.AddDate(0, 0, request.Day)
	text.Link = generateLink()

	return s.repo.Text.Create(ctx, text)
}

func (s *TextService) GetByLink(ctx context.Context, link string) (*model.Text, error) {
	text, err := s.repo.Text.GetByLink(ctx, link)
	if err != nil {
		return nil, err
	}

	if checkExpired(text) {
		return nil, errors.New("text's expired")
	}

	return text, nil
}

func (s *TextService) GetByID(ctx context.Context, ID uint) (*model.Text, error) {
	text, err := s.repo.Text.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	if checkExpired(text) {
		return nil, errors.New("text's expired")
	}

	return text, nil
}

func generateLink() string {
	b := make([]byte, lengthBytes)

	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func checkExpired(text *model.Text) bool {
	if text.ExpiresAt.Unix() < time.Now().Unix() {
		return true
	}

	return false
}
