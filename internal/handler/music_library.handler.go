package handler

import (
	"net/http"
	"testEffective-Mobile/internal/services"
	"testEffective-Mobile/x/interfacesx"

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

func (h *MusicLibraryHandler) AddSong(c *gin.Context) {
	var songRequest interfacesx.SongAddRequest

	if err := c.ShouldBindJSON(&songRequest); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.libraryService.AddSong(&songRequest); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Song add successfully")

}

func (h *MusicLibraryHandler) DeleteSong(c *gin.Context) {
	songId := c.Param("songid")
	if songId == "" {
		c.JSON(http.StatusBadRequest, "param is null")
		return
	}

	if ok, err := h.libraryService.DeleteSong(songId); !ok || err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	} else if ok {
		c.JSON(http.StatusOK, "song was deleted")
	}

}

func (h *MusicLibraryHandler) UpdateSong(c *gin.Context) {
	songId := c.Param("songid")
	var songRequest interfacesx.SongAddRequest

	if err := c.ShouldBindJSON(&songRequest); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.libraryService.UpdateSong(&songRequest, songId); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Song update successfully")

}

func (h *MusicLibraryHandler) GetAllMusicLibraryData(c *gin.Context) {

	songs, err := h.libraryService.GetAllMusicLibraryData(c)

	if err != nil {
		c.JSON(http.StatusNotFound, "songs not found")
		return
	}

	c.JSON(http.StatusOK, songs)
}

func (h *MusicLibraryHandler) GetSongText(c *gin.Context) {

}
