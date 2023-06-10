package service

import (
	"errors"
	"github.com/jumagaliev1/jiberSoz/internal/storage"
	"github.com/jumagaliev1/jiberSoz/internal/storage/redis"
	s32 "github.com/jumagaliev1/jiberSoz/internal/storage/s3"
	"github.com/spf13/viper"
)

type Service struct {
	Text ITextService
}

func New(repo *storage.Repository) (*Service, error) {
	if repo == nil {
		return nil, errors.New("NO repo")
	}
	s3 := s32.NewAmazonS3()
	cacheView := redis.New(viper.GetString("redis.view.uri"))
	cachePost := redis.New(viper.GetString("redis.post.uri"))
	textService := NewTextService(repo, s3, cacheView, cachePost)

	return &Service{
		Text: textService,
	}, nil
}
