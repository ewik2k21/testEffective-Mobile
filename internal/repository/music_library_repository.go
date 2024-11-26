package repository

import (
	"strconv"
	"testEffective-Mobile/internal/model"
	"testEffective-Mobile/x/interfacesx"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MusicLibraryRepository interface {
	DeleteSong(songId string) (bool, error)
	AddSong(songRequest *interfacesx.SongAddRequest) error
	GetAllMusicLibraryData(c *gin.Context) (*[]model.MusicInfo, error)
	UpdateSong(songRequest *interfacesx.SongAddRequest, songId string) error
}

type musicLibraryRepository struct {
	db *gorm.DB
}

func NewMusicLibraryRepository(db *gorm.DB) MusicLibraryRepository {
	return &musicLibraryRepository{db}
}

func (r *musicLibraryRepository) DeleteSong(songId string) (bool, error) {
	if err := r.db.Model(&model.MusicInfo{}).Where("id = ?", songId).UpdateColumn("deleted_at", time.Now()).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *musicLibraryRepository) AddSong(songRequest *interfacesx.SongAddRequest) error {
	song := &model.MusicInfo{
		Group: songRequest.Group,
		Song:  songRequest.Song,
		SongDetails: model.SongDetails{
			ReleaseDate: songRequest.ReleaseDate,
			Text:        songRequest.Text,
			Link:        songRequest.Link,
		},
	}

	if err := r.db.Create(&song).Error; err != nil {
		return err
	}
	return nil
}

func (r *musicLibraryRepository) UpdateSong(songRequest *interfacesx.SongAddRequest, songId string) error {
	var song model.MusicInfo
	var songDetails model.SongDetails

	if err := r.db.Where("id = ?", songId).First(&song).Error; err != nil {
		return err
	}
	if err := r.db.Where("music_info_id = ?", songId).First(&songDetails).Error; err != nil {
		return err
	}
	song.Group = songRequest.Group
	song.Song = songRequest.Song
	songDetails.ReleaseDate = songRequest.ReleaseDate
	songDetails.Text = songRequest.Text
	songDetails.Link = songRequest.Link

	if err := r.db.Save(&song).Error; err != nil {
		return err
	}
	if err := r.db.Save(&songDetails).Error; err != nil {
		return err
	}
	return nil
}

func (r *musicLibraryRepository) GetAllMusicLibraryData(c *gin.Context) (*[]model.MusicInfo, error) {
	songs := &[]model.MusicInfo{}

	if err := r.db.Scopes(Paginate(c)).Find(&songs).Error; err != nil {
		return nil, err
	}

	return songs, nil

}

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageStr := c.DefaultQuery("page", "1")
		page, _ := strconv.Atoi(pageStr)

		if page <= 0 {
			page = 1
		}

		pageSizeStr := c.DefaultQuery("page_size", "10")
		pageSize, _ := strconv.Atoi(pageSizeStr)

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
