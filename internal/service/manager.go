package service

import (
	"errors"
	pb "github.com/jumagaliev1/jiberSoz/hasher/proto"
	"github.com/jumagaliev1/jiberSoz/internal/storage"
	"github.com/jumagaliev1/jiberSoz/internal/storage/redis"
	s32 "github.com/jumagaliev1/jiberSoz/internal/storage/s3"
)

type Service struct {
	Text ITextService
}

func New(repo *storage.Repository, s3 *s32.AmazonS3, cacheView *redis.RedisClient, cachePost *redis.RedisClient, hasher pb.HashServiceClient) (*Service, error) {
	if repo == nil {
		return nil, errors.New("no given repo")
	}

	textService := NewTextService(repo, s3, cacheView, cachePost, hasher)

	return &Service{
		Text: textService,
	}, nil
}
