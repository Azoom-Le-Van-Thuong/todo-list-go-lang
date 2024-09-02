package storage

import (
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewMySqlStore(db *gorm.DB) *Store {
	return &Store{db}
}
