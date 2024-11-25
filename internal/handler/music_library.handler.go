package handler

import "testEffective-Mobile/internal/services"

type MusicLibraryHandler struct {
	libraryService services.MusicLibraryService
}

func NewMusicLibraryHandler(libraryService services.MusicLibraryService) *MusicLibraryHandler {
	return &MusicLibraryHandler{
		libraryService: libraryService,
	}
}
