package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type MusicInfo struct {
	gorm.Model
	ID          uuid.UUID   `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Group       string      `gorm:"not_null"`
	Song        string      `gorm:"not_null"`
	SongDetails SongDetails `gorm:"not_null;foreignKey:MusicInfoId"`
}

type SongDetails struct {
	gorm.Model
	MusicInfoId uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	ReleaseDate string    `gorm:"not_null"`
	Text        string    `gorm:"not_null"`
	Link        string    `gorm:"not_null"`
}
