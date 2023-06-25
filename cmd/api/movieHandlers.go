package main

import (
	"errors"
	"movies/models"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.logger.Println(errors.New("Invalid Id parameter"))
	}

	app.logger.Println("id is", id)

	movie := models.Movie{
		ID:          id,
		Title:       "Some movie",
		Description: "Some description",
		Year:        2000,
		Runtime:     200,
		Rating:      4,
		MPAARating:  "R",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		ReleaseDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
	}

	err = app.writeJson(w, http.StatusOK, movie, "movie")
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {}
