package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, error := strconv.Atoi(params.ByName("id"))
	if error != nil {
		app.logger.Print(errors.New("invalid id"))
		app.ErrorJSON(w, error)
		return
	}

	app.logger.Print("id: ", id)

	movie, error := app.models.DB.Get(id)

	// movie := models.Movie{
	// 	ID:          id,
	// 	Title:       "title",
	// 	Description: "description",
	// 	Year:        2021,
	// 	ReleaseDate: time.Date(2020, 01, 01, 01, 0, 0, 0, time.Local),
	// 	Runtime:     100,
	// 	Rating:      5,
	// 	MPAARating:  "PG-13",
	// 	CreatedAt:   time.Now(),
	// 	UpdatedAt:   time.Now(),
	// }

	error = app.WriteJson(w, http.StatusOK, movie, "movie")
	if error != nil {
		app.ErrorJSON(w, error)
		return
	}
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, error := app.models.DB.All()
	if error != nil {
		app.ErrorJSON(w, error)
		return
	}
	error = app.WriteJson(w, http.StatusOK, movies, "movies")
	if error != nil {
		app.ErrorJSON(w, error)
		return
	}
}
