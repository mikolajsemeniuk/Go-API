package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Get(id int) (*Movie, error) {
	context, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
			SELECT
				id, 
				title, 
				description, 
				year, 
				release_date, 
				runtime, 
				mpaa_rating,
				created_at,
				updated_at,
			FROM
				movies
			WHERE
				id = $1`

	row := m.DB.QueryRowContext(context, query, id)

	var movie Movie

	error := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Rating,
		&movie.Runtime,
		&movie.MPAARating,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)

	if error != nil {
		return nil, error
	}

	return &movie, nil
}

func (m *DBModel) All(id int) ([]*Movie, error) {
	return nil, nil
}
