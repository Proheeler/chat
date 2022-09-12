package postgres

import (
	"chat/internal/types"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresStorage struct {
	db *gorm.DB
}

func NewPostgresStorageStorage() *PostgresStorage {
	dsn := "host=localhost user=postgres password=dev dbname=chat port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	db.AutoMigrate(&types.Client{}, &types.Media{}, &types.Message{}, &types.Room{})
	return &PostgresStorage{
		db: db,
	}
}
