package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

// Get return one movie and error, if any
func (m *DBModel) Get(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, description, year, releaseDate, runtime, rating, mpaaRating, createdAt, updatedAt from movies where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie

	err := row.Scan(
		&movie.ID, &movie.Title, &movie.Description, &movie.Year, &movie.ReleaseDate, &movie.Rating, &movie.Runtime, &movie.MPAARating, &movie.CreatedAt, &movie.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

// All return all movies and error, if any
func (m *DBModel) All(id int) ([]*Movie, error) {
	return nil, nil
}
