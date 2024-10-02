package models

// We use in this project 'DAO' and 'DTO' models
// Easily separate them with tags

// DTO is
type DTO struct {
	ID          int    `json:"id" validate:"required"`
	Group       string `json:"group" validate:"required"`
	Title       string `json:"title" validate:"required"`
	ReleaseDate string `json:"release_date" validate:"required"`
	Text        string `json:"text" validate:"required"`
	Link        string `json:"link" validate:"required"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// DAO is
type DAO struct {
	ID          int    `db:"id"`
	Group       string `db:"group"`
	Title       string `db:"title"`
	ReleaseDate string `db:"release_date"`
	Text        string `db:"text"`
	Link        string `db:"link"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}

// Now we will use in service and handler layer DTO models
// In repository layer we only use DAO model

// ToStorage is
func (d *DTO) ToStorage() *DAO {
	return &DAO{
		ID:          d.ID,
		Group:       d.Group,
		Title:       d.Title,
		ReleaseDate: d.ReleaseDate,
		Text:        d.Text,
		Link:        d.Link,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}

// ToServer is
func (d *DAO) ToServer() *DTO {
	return &DTO{
		ID:          d.ID,
		Group:       d.Group,
		Title:       d.Title,
		ReleaseDate: d.ReleaseDate,
		Text:        d.Text,
		Link:        d.Link,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}
