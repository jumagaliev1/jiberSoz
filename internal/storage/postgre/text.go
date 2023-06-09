package postgre

import (
	"context"
	"github.com/jumagaliev1/jiberSoz/internal/model"
	"gorm.io/gorm"
)

type TextRepository struct {
	DB *gorm.DB
}

func NewTextRepository(DB *gorm.DB) *TextRepository {
	return &TextRepository{
		DB: DB,
	}
}

func (r *TextRepository) Create(ctx context.Context, text model.Text) (*model.Text, error) {
	if err := r.DB.WithContext(ctx).Create(&text).Error; err != nil {
		return nil, err
	}

	return &text, nil
}

func (r *TextRepository) GetByLink(ctx context.Context, link string) (*model.Text, error) {
	var text model.Text

	if err := r.DB.WithContext(ctx).Where("link = ?", link).Find(&text).Error; err != nil {
		return nil, err
	}

	return &text, nil
}

func (r *TextRepository) GetByID(ctx context.Context, ID uint) (*model.Text, error) {
	var text model.Text

	if err := r.DB.WithContext(ctx).Where("id = ?", ID).Find(&text).Error; err != nil {
		return nil, err
	}

	return &text, nil
}
