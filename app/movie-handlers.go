package main

import (
	"errors"
	"net/http"
	"server/app/models"
	"strconv"
	"time"

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

	movie := models.Movie{
		ID:          id,
		Title:       "title",
		Description: "description",
		Year:        2021,
		ReleaseDate: time.Date(2020, 01, 01, 01, 0, 0, 0, time.Local),
		Runtime:     100,
		Rating:      5,
		MPAARating:  "PG-13",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	error = app.WriteJson(w, http.StatusOK, movie, "movie")
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {

}
