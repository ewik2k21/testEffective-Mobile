package repository

import (
	"testEffective-Mobile/internal/model"
	"testEffective-Mobile/x/interfacesx"
	"time"

	filter "github.com/ActiveChooN/gin-gorm-filter"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MusicLibraryRepository interface {
	DeleteSong(songId string) (bool, error)
	AddSong(songRequest *interfacesx.SongAddRequest) error
	GetAllMusicLibraryData(c *gin.Context) (*[]model.MusicInfo, error)
	UpdateSong(songRequest *interfacesx.SongAddRequest, songId string) error
	GetSongTextById(c *gin.Context, songId string) (*string, error)
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

	if err := r.db.Model(songs).Scopes(filter.FilterByQuery(c, filter.ALL)).Find(&songs).Error; err != nil {
		return nil, err
	}

	return songs, nil

}
func (r *musicLibraryRepository) GetSongTextById(c *gin.Context, songId string) (*string, error) {
	var songDetails model.SongDetails
	if err := r.db.Model(songDetails).Scopes(filter.FilterByQuery(c, filter.PAGINATE)).Where("music_info_id = ?", songId).First(&songDetails).Error; err != nil {
		return nil, err
	}

	return &songDetails.Text, nil
}
