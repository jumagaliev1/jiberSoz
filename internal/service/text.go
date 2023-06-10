package service

import (
	"context"
	"errors"
	"github.com/jumagaliev1/jiberSoz/internal/model"
	"github.com/jumagaliev1/jiberSoz/internal/storage"
	"github.com/jumagaliev1/jiberSoz/internal/storage/s3"
	"github.com/labstack/gommon/log"
	"math/rand"
	"time"
)

var (
	lengthBytes = 8
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type TextService struct {
	repo *storage.Repository
	s3   *s3.AmazonS3
}

func NewTextService(repo *storage.Repository, amazonS3 *s3.AmazonS3) *TextService {
	return &TextService{
		repo: repo,
		s3:   amazonS3,
	}
}

func (s *TextService) Create(ctx context.Context, request model.TextRequest) (*model.Text, error) {
	text := request.ToText()
	text.CreatedAt = time.Now()
	text.ExpiresAt = text.CreatedAt.AddDate(0, 0, request.Day)
	text.Link = generateLink()

	err := s.s3.Upload(text.Link, text.Message)
	if err != nil {
		return nil, err
	}

	return s.repo.Text.Create(ctx, text)
}

func (s *TextService) GetByLink(ctx context.Context, link string) (*model.Text, error) {
	text, err := s.repo.Text.GetByLink(ctx, link)
	if err != nil {
		return nil, err
	}
	log.Info(text.Message)
	if checkExpired(text) {
		return nil, errors.New("text's expired")
	}

	ans, err := s.s3.Download(text.Link)
	if err != nil {
		return nil, err
	}
	text.Message = ans

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
