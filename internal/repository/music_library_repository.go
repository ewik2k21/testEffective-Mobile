package repository

import "gorm.io/gorm"

type MusicLibraryRepository interface {
}

type musicLibraryRepository struct {
	db *gorm.DB
}

func NewMusicLibraryRepository(db *gorm.DB) MusicLibraryRepository {
	return &musicLibraryRepository{db}
}
