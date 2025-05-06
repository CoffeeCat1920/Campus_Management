package api

import (
	"net/http"
	"what/internal/database"

	"github.com/spf13/afero"
)

type Api interface {
	GetAllBooksHandler(w http.ResponseWriter, r *http.Request)
	AddBookHandler(w http.ResponseWriter, r *http.Request)
	DeleteBookHandler(w http.ResponseWriter, r *http.Request)
	EditBookHandler(w http.ResponseWriter, r *http.Request)
	ToggleBookHandler(w http.ResponseWriter, r *http.Request)
}

type api struct {
	db database.Service
	fs afero.Fs
}

func NewApi() Api {
	return &api{
		db: database.New(),
		fs: afero.OsFs{},
	}
}
