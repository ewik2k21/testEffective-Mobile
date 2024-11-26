package repository

import (
	"testEffective-Mobile/internal/model"
	"testEffective-Mobile/x/interfacesx"
	"time"

	"gorm.io/gorm"
)

type MusicLibraryRepository interface {
	DeleteSong(songId string) (bool, error)
	AddSong(songRequest *interfacesx.SongAddRequest) error
	GetAllMusicLibraryData() (*[]model.MusicInfo, error)
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

func (r *musicLibraryRepository) GetAllMusicLibraryData() (*[]model.MusicInfo, error) {
	songs := &[]model.MusicInfo{}

	if err := r.db.Preload("SongDetails").Find(songs).Error; err != nil {
		return nil, err
	}

	return songs, nil

}
