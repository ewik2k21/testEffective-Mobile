package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type MusicInfo struct {
	gorm.Model
	ID          uuid.UUID
	Group       string
	Song        string
	SongDetails SongDetails
}

type SongDetails struct {
	gorm.Model
	ReleaseDate string
	Text        string
	Link        string
}
