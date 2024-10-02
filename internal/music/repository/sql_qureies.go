package repository

// SQL Queries for songs
const (
	// addSongQuery is
	addSongQuery = `
		INSERT INTO songs (group, title, release_date, text, link)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`
)
