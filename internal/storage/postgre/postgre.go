package postgre

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Dial(ctx context.Context) (*gorm.DB, error) {
	//db, err := gorm.Open(postgres.Open(viper.GetString("postgres.uri")), &gorm.Config{})
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate()

	return db, err
}
