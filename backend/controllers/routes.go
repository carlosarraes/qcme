package controllers

import (
	"database/sql"
	"net/http"

	"github.com/carlosarraes/qcmeback/models"
	"github.com/carlosarraes/qcmeback/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	DSN string
	DB  models.Data
}

func (a *App) Connect() (*sql.DB, error) {
	db, err := models.OpenDB(a.DSN)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (a *App) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(utils.Cors())

	mux.Get("/qrcodeme/{name}", a.GenerateQRCode)

	return mux
}
