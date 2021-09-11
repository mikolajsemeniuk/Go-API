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
			rating, 
			runtime, 
			mpaa_rating,
			created_at,
			updated_at 
		FROM 
			movies
		WHERE 
			id = $1
	`

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

	// get genres if any
	query = `
		SELECT
			mg.id,
			mg.movie_id,
			mg.genre_id,
			g.genre_name
		FROM
			movies_genres mg
		LEFT JOIN
			genres g
		ON
			(g.id = mg.genre_id)
		WHERE
			mg.movie_id = $1
	`

	rows, _ := m.DB.QueryContext(context, query, id)
	defer rows.Close()

	genres := make(map[int]string)

	for rows.Next() {
		var mg MovieGenre
		error := rows.Scan(
			&mg.ID,
			&mg.MovieId,
			&mg.GenreId,
			&mg.Genre.GenreName,
		)
		if error != nil {
			return nil, error
		}
		genres[mg.ID] = mg.Genre.GenreName
	}

	movie.MovieGenre = genres

	return &movie, nil
}

func (m *DBModel) All() ([]*Movie, error) {
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
	SELECT
		id,
		title,
		description,
		year,
		release_date,
		rating,
		runtime,
		mpaa_rating,
		created_at,
		updated_at
	FROM
		movies
	ORDER BY
		title
	`

	rows, error := m.DB.QueryContext(context, query)
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	var movies []*Movie

	for rows.Next() {
		var movie Movie
		error := rows.Scan(
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

		genreQuery := `
			SELECT
				mg.id,
				mg.movie_id,
				mg.genre_id,
				g.genre_name
			FROM
				movies_genres mg
			LEFT JOIN
				genres g
			ON
				(g.id = mg.genre_id)
			WHERE
				mg.movie_id = $1
		`

		genreRows, _ := m.DB.QueryContext(context, genreQuery, movie.ID)

		genres := make(map[int]string)

		for genreRows.Next() {
			var mg MovieGenre
			error := genreRows.Scan(
				&mg.ID,
				&mg.MovieId,
				&mg.GenreId,
				&mg.Genre.GenreName,
			)
			if error != nil {
				return nil, error
			}
			genres[mg.ID] = mg.Genre.GenreName
		}
		genreRows.Close()

		movie.MovieGenre = genres
		movies = append(movies, &movie)
	}

	return movies, nil
}
