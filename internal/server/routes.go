package server

import (
	"encoding/json"
	"log"
	"net/http"
	"what/internal/api"
	"what/internal/auth"

	"github.com/gorilla/mux"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()

	r.Use(s.corsMiddleware)

	auth := auth.NewAuthService()
	api := api.NewApi(auth)

	// Books
	r.HandleFunc("/all_books", api.GetAllBooksHandler).Methods("GET")
	r.HandleFunc("/student/all_books", auth.AuthStudent(api.BooksExcept)).Methods("GET")

	r.HandleFunc("/books", auth.AuthLibrarian(api.AddBookHandler)).Methods("POST")
	r.HandleFunc("/books/{id}", auth.AuthLibrarian(api.GetBookHandler)).Methods("GET")
	r.HandleFunc("/books/{id}", auth.AuthLibrarian(api.DeleteBookHandler)).Methods("DELETE")
	r.HandleFunc("/books/{id}", auth.AuthLibrarian(api.EditBookHandler)).Methods("PATCH")
	r.HandleFunc("/toggle_books/{id}", auth.AuthLibrarian(api.ToggleBookHandler)).Methods("PATCH")

	// Users
	r.HandleFunc("/login", api.LoginUserHandler).Methods("POST")
	r.HandleFunc("/logout", api.LogoutUserHandler).Methods("POST")

	// Librarian
	r.HandleFunc("/all_librarians", api.GetAllLibrariansHandler).Methods("GET")

	r.HandleFunc("/librarian/data", api.LoginLibrarianDataHandler).Methods("GET")
	r.HandleFunc("/librarian", api.AddLibrarianHandler).Methods("POST")
	r.HandleFunc("/librarian/{id}", auth.AuthAdmin(api.EditLibrarianHandler)).Methods("PATCH")
	r.HandleFunc("/librarian/{id}", auth.AuthAdmin(api.DeleteLibrarianHandler)).Methods("DELETE")
	r.HandleFunc("/librarian/{id}", auth.AuthAdmin(api.GetLibrarianHandler)).Methods("GET")

	// Students
	r.HandleFunc("/all_students", api.GetAllStudentsHandler).Methods("GET")

	r.HandleFunc("/student/data", api.LoginStudentDataHandler).Methods("GET")
	r.HandleFunc("/student", auth.AuthAdmin(api.AddStudentHandler)).Methods("POST")
	r.HandleFunc("/student/{id}", auth.AuthAdmin(api.EditStudentHandler)).Methods("PATCH")
	r.HandleFunc("/student/{id}", auth.AuthAdmin(api.DeleteStudentHandler)).Methods("DELETE")
	r.HandleFunc("/student/{id}", auth.AuthAdmin(api.GetStudentHandler)).Methods("GET")

	r.HandleFunc("/student/nob/{id}", auth.AuthUser(api.NumberOfBorrowHandler)).Methods("GET")

	// Borrow
	r.HandleFunc("/borrow/{id}", api.GetBorrowByUserHandler).Methods("GET")
	r.HandleFunc("/borrow", auth.AuthLibrarian(api.AddBorrowHandler)).Methods("POST")
	r.HandleFunc("/borrow_fine", auth.AuthLibrarian(api.BorrowFineHandler)).Methods("GET")
	r.HandleFunc("/return_book/{id}", api.ReturnBookHandler).Methods("POST")

	// Request
	r.HandleFunc("/request", auth.AuthStudent(api.RequestBorrowHandler)).Methods("POST")
	r.HandleFunc("/accept_request/{id}", auth.AuthLibrarian(api.AcceptRequestHandler)).Methods("POST")
	r.HandleFunc("/decline_request/{id}", auth.AuthLibrarian(api.DeclineRequestHandler)).Methods("POST")
	r.HandleFunc("/requests/{id}", auth.AuthLibrarian(api.GetRequestByUserHandler)).Methods("GET")

	// Admin
	r.HandleFunc("/admin/login", api.LoginAdminHandler).Methods("POST")
	r.HandleFunc("/admin/logout", api.LogoutAdminHandler).Methods("POST")
	r.HandleFunc("/admin/data", api.LoginAdminDataHandler).Methods("GET")

	r.HandleFunc("/health", s.healthHandler)

	return r
}

func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Wildcard allows all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Credentials not allowed with wildcard origins

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
