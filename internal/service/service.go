package service

import (
	"context"
	"github.com/jumagaliev1/jiberSoz/internal/model"
)

type ITextService interface {
	Create(ctx context.Context, request model.TextRequest) (*model.Text, error)
	GetByLink(ctx context.Context, link string) (*model.TextResponse, error)
	GetByID(ctx context.Context, ID uint) (*model.TextResponse, error)
}
