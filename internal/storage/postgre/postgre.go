package postgre

import (
	"context"
	"github.com/jumagaliev1/jiberSoz/internal/model"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Dial(ctx context.Context) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(viper.GetString("postgres.uri")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Text{})

	return db, err
}
