package services

import "testEffective-Mobile/internal/repository"

type MusicLibraryService interface {
}

type musicLibraryService struct {
	musicLibraryRepo repository.MusicLibraryRepository
}

func NewMusicLibraryService(musicLibraryRepo repository.MusicLibraryRepository) MusicLibraryService {
	return &musicLibraryService{
		musicLibraryRepo: musicLibraryRepo,
	}
}
