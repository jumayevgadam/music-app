package models

// Same logic, we use methods ToStorage() and ToServer()
// for clearly using models in layers

// SongDetailDTO struct is
type SongDetailDTO struct {
	ReleaseDate string `json:"release_date" validate:"required"`
	Text        string `json:"text" validate:"required"`
	Link        string `json:"link" validate:"required"`
}

// SongDetailDAO struct is
type SongDetailDAO struct {
	ReleaseDate string `db:"release_date"`
	Text        string `db:"text"`
	Link        string `db:"link"`
}

// ToStorage is
func (s *SongDetailDTO) ToStorage() *SongDetailDAO {
	return &SongDetailDAO{
		ReleaseDate: s.ReleaseDate,
		Text:        s.Text,
		Link:        s.Link,
	}
}

// ToServer is
func (s *SongDetailDAO) ToServer() *SongDetailDTO {
	return &SongDetailDTO{
		ReleaseDate: s.ReleaseDate,
		Text:        s.Text,
		Link:        s.Link,
	}
}
