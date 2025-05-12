package api

import (
	"net/http"
	"what/internal/auth"
	"what/internal/database"

	"github.com/spf13/afero"
)

type Api interface {
	// Book Api
	GetAllBooksHandler(w http.ResponseWriter, r *http.Request)
	AddBookHandler(w http.ResponseWriter, r *http.Request)
	DeleteBookHandler(w http.ResponseWriter, r *http.Request)
	EditBookHandler(w http.ResponseWriter, r *http.Request)
	ToggleBookHandler(w http.ResponseWriter, r *http.Request)
	GetBookHandler(w http.ResponseWriter, r *http.Request)

	// User

	// Student Handler
	GetAllStudentsHandler(w http.ResponseWriter, r *http.Request)
	AddStudentHandler(w http.ResponseWriter, r *http.Request)
	GetStudentHandler(w http.ResponseWriter, r *http.Request)
	LoginUserHandler(w http.ResponseWriter, r *http.Request)
	EditStudentHandler(w http.ResponseWriter, r *http.Request)
	LoginStudentDataHandler(w http.ResponseWriter, r *http.Request)

	// Librarian Handler
	GetAllLibrariansHandler(w http.ResponseWriter, r *http.Request)
	AddLibrarianHandler(w http.ResponseWriter, r *http.Request)
	GetLibrarianHandler(w http.ResponseWriter, r *http.Request)
	LoginLibrarianHandler(w http.ResponseWriter, r *http.Request)
	EditLibrarianHandler(w http.ResponseWriter, r *http.Request)
	DeleteLibrarianHandler(w http.ResponseWriter, r *http.Request)
	LoginLibrarianDataHandler(w http.ResponseWriter, r *http.Request)

	// Borrow Handler
	AddBorrowHandler(w http.ResponseWriter, r *http.Request)
	BorrowFineHandler(w http.ResponseWriter, r *http.Request)
	ReturnBookHandler(w http.ResponseWriter, r *http.Request)

	RequestBorrowHandler(w http.ResponseWriter, r *http.Request)
	AcceptRequestHandler(w http.ResponseWriter, r *http.Request)
	DeleteStudentHandler(w http.ResponseWriter, r *http.Request)

	// Admin Handler
	LoginAdminHandler(w http.ResponseWriter, r *http.Request)
	LogoutAdminHandler(w http.ResponseWriter, r *http.Request)
	LoginAdminDataHandler(w http.ResponseWriter, r *http.Request)
}

type api struct {
	auth auth.AuthService
	db   database.Service
	fs   afero.Fs
}

func NewApi(auth auth.AuthService) Api {
	return &api{
		db:   database.New(),
		fs:   afero.OsFs{},
		auth: auth,
	}
}
