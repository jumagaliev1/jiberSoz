package storage

import (
	"context"
	"github.com/jumagaliev1/jiberSoz/internal/storage/postgre"
)

type Repository struct {
}

func New(ctx context.Context) (*Repository, error) {
	_, err := postgre.Dial(ctx)
	if err != nil {
		return nil, err
	}
	return &Repository{}, nil
}
