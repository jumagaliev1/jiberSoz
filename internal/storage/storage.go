package storage

import (
	"context"
	"github.com/jumagaliev1/jiberSoz/internal/model"
	"github.com/jumagaliev1/jiberSoz/internal/storage/postgre"
)

type ITextRepository interface {
	Create(ctx context.Context, text model.Text) (*model.Text, error)
	GetByLink(ctx context.Context, link string) (*model.Text, error)
	GetByID(ctx context.Context, ID uint) (*model.Text, error)
}

type Repository struct {
	Text ITextRepository
}

func New(ctx context.Context) (*Repository, error) {
	db, err := postgre.Dial(ctx)
	if err != nil {
		return nil, err
	}

	textRepo := postgre.NewTextRepository(db)

	return &Repository{
		Text: textRepo,
	}, nil
}
