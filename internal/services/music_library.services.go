package services

import (
	"testEffective-Mobile/internal/model"
	"testEffective-Mobile/internal/repository"
	"testEffective-Mobile/x/interfacesx"
)

type MusicLibraryService interface {
	DeleteSong(songId string) (bool, error)
	AddSong(songRequest *interfacesx.SongAddRequest) error
	GetAllMusicLibraryData() (*[]interfacesx.SongData, error)
	UpdateSong(songRequest *interfacesx.SongAddRequest, songId string) error
}

type musicLibraryService struct {
	musicLibraryRepo repository.MusicLibraryRepository
}

func NewMusicLibraryService(musicLibraryRepo repository.MusicLibraryRepository) MusicLibraryService {
	return &musicLibraryService{
		musicLibraryRepo: musicLibraryRepo,
	}
}

func (s *musicLibraryService) DeleteSong(songId string) (bool, error) {
	if ok, err := s.musicLibraryRepo.DeleteSong(songId); err != nil {
		return ok, err
	}
	return true, nil
}

func (s *musicLibraryService) AddSong(songRequest *interfacesx.SongAddRequest) error {

	if err := s.musicLibraryRepo.AddSong(songRequest); err != nil {
		return err
	}

	return nil
}

func (s *musicLibraryService) UpdateSong(songRequest *interfacesx.SongAddRequest, songId string) error {
	if err := s.musicLibraryRepo.UpdateSong(songRequest, songId); err != nil {
		return err
	}
	return nil
}

func (s *musicLibraryService) GetAllMusicLibraryData() (*[]interfacesx.SongData, error) {
	songs, err := s.musicLibraryRepo.GetAllMusicLibraryData()

	if err != nil {
		return nil, err
	}
	songsData := make([]interfacesx.SongData, len(*songs))

	for i, song := range *songs {
		songsData[i] = interfacesx.SongData{
			ID:    song.ID,
			Group: song.Group,
			Song:  song.Song,
			SongDetails: &model.SongDetails{
				MusicInfoId: song.SongDetails.MusicInfoId,
				ReleaseDate: song.SongDetails.ReleaseDate,
				Text:        song.SongDetails.Text,
				Link:        song.SongDetails.Link,
			},
		}
	}
	return &songsData, nil
}
