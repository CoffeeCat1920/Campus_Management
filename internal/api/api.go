package api

import (
	"github.com/spf13/afero"
	"net/http"
	"what/internal/database"
)

type Api interface {
	GetAllBooksHandler(w http.ResponseWriter, r *http.Request)
	AddBookHandler(w http.ResponseWriter, r *http.Request)
	DeleteBookHandler(w http.ResponseWriter, r *http.Request)
	EditBookHandler(w http.ResponseWriter, r *http.Request)
	ToggleBookHandler(w http.ResponseWriter, r *http.Request)

	AddStudentHandler(w http.ResponseWriter, r *http.Request)

	AddBorrowHandler(w http.ResponseWriter, r *http.Request)
	BorrowFineHandler(w http.ResponseWriter, r *http.Request)
	ReturnBookHandler(w http.ResponseWriter, r *http.Request)

	RequestBorrowHandler(w http.ResponseWriter, r *http.Request)
	AcceptRequestHandler(w http.ResponseWriter, r *http.Request)
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
