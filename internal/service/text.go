package service

import (
	"context"
	"github.com/jumagaliev1/jiberSoz/internal/model"
	"github.com/jumagaliev1/jiberSoz/internal/storage"
	"math/rand"
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
	text.Link = generateLink()

	return s.repo.Text.Create(ctx, text)
}

func (s *TextService) GetByLink(ctx context.Context, link string) (*model.Text, error) {
	return s.repo.Text.GetByLink(ctx, link)
}

func (s *TextService) GetByID(ctx context.Context, ID uint) (*model.Text, error) {
	return s.repo.Text.GetByID(ctx, ID)
}

func generateLink() string {
	b := make([]byte, lengthBytes)

	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
