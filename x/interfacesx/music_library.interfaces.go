package interfacesx

import (
	"testEffective-Mobile/internal/model"

	"github.com/gofrs/uuid"
)

type SongAddRequest struct {
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type SongData struct {
	ID          uuid.UUID          `json:"id"`
	Group       string             `json:"group"`
	Song        string             `json:"song"`
	SongDetails *model.SongDetails `json:"song_details"`
}
