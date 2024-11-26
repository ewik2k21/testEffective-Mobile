package handler

import (
	"testEffective-Mobile/internal/services"

	"github.com/gin-gonic/gin"
)

type MusicLibraryHandler struct {
	libraryService services.MusicLibraryService
}

func NewMusicLibraryHandler(libraryService services.MusicLibraryService) *MusicLibraryHandler {
	return &MusicLibraryHandler{
		libraryService: libraryService,
	}
}

func (h *MusicLibraryHandler) DeleteSong(c *gin.Context) {

}

func (h *MusicLibraryHandler) UpdateSong(c *gin.Context) {

}

func (h *MusicLibraryHandler) AddSong(c *gin.Context) {

}

func (h *MusicLibraryHandler) GetAllMusicLibraryData(c *gin.Context) {

}

func (h *MusicLibraryHandler) GetSongText(c *gin.Context) {

}
