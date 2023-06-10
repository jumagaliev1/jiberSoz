package service

import (
	"context"
	"errors"
	"github.com/jumagaliev1/jiberSoz/internal/model"
	"github.com/jumagaliev1/jiberSoz/internal/storage"
	"github.com/jumagaliev1/jiberSoz/internal/storage/redis"
	"github.com/jumagaliev1/jiberSoz/internal/storage/s3"
	"github.com/labstack/gommon/log"
	"math/rand"
	"strconv"
	"time"
)

var (
	lengthBytes = 8
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	minView     = 10
)

type TextService struct {
	repo      *storage.Repository
	s3        *s3.AmazonS3
	cacheView *redis.RedisClient
	cachePost *redis.RedisClient
}

func NewTextService(repo *storage.Repository, amazonS3 *s3.AmazonS3, cacheView *redis.RedisClient, cachePost *redis.RedisClient) *TextService {
	return &TextService{
		repo:      repo,
		s3:        amazonS3,
		cacheView: cacheView,
		cachePost: cachePost,
	}
}

func (s *TextService) Create(ctx context.Context, request model.TextRequest) (*model.Text, error) {
	message := request.Message

	text := request.ToText()
	text.CreatedAt = time.Now()
	text.ExpiresAt = text.CreatedAt.AddDate(0, 0, request.Day)
	text.Link = generateLink()

	err := s.s3.Upload(text.Link, message)
	if err != nil {
		return nil, err
	}

	err = s.cacheView.Set(ctx, text.Link, 0)
	if err != nil {
		return nil, err
	}

	return s.repo.Text.Create(ctx, text)
}

func (s *TextService) GetByLink(ctx context.Context, link string) (*model.TextResponse, error) {
	text, err := s.repo.Text.GetByLink(ctx, link)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	view, err := s.cacheView.Get(ctx, text.Link)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if checkExpired(text) {
		return nil, errors.New("text's expired")
	}

	viewCount, err := strconv.Atoi(view.(string))
	if err != nil {
		return nil, err
	}

	message, err := s.getMessage(ctx, *text, viewCount)
	if err != nil {
		return nil, err
	}

	if viewCount == minView {
		err := s.cachePost.Set(ctx, text.Link, message)
		if err != nil {
			return nil, err
		}
	}

	viewCount++
	if err := s.cacheView.Set(ctx, text.Link, view); err != nil {
		return nil, err
	}

	response := text.ToTextResponse()
	response.Message = message

	return &response, nil
}

func (s *TextService) GetByID(ctx context.Context, ID uint) (*model.TextResponse, error) {
	text, err := s.repo.Text.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	view, err := s.cacheView.Get(ctx, text.Link)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if checkExpired(text) {
		return nil, errors.New("text's expired")
	}

	viewCount, err := strconv.Atoi(view.(string))
	if err != nil {
		return nil, err
	}

	message, err := s.getMessage(ctx, *text, viewCount)
	if err != nil {
		return nil, err
	}

	if viewCount == minView {
		err := s.cachePost.Set(ctx, text.Link, message)
		if err != nil {
			return nil, err
		}
	}

	viewCount++
	if err := s.cacheView.Set(ctx, text.Link, view); err != nil {
		return nil, err
	}

	response := text.ToTextResponse()
	response.Message = message

	return &response, nil
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

func (s *TextService) getMessage(ctx context.Context, text model.Text, viewCount int) (string, error) {
	var message string

	if viewCount > minView {
		res, err := s.cachePost.Get(ctx, text.Link)
		if err != nil {
			return "", err
		}
		message = res.(string)
	} else {
		res, err := s.s3.Download(text.Link)
		if err != nil {
			return "", err
		}
		message = res
	}

	return message, nil
}