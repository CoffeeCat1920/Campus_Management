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

	r.HandleFunc("/books", auth.AuthLibrarian(api.AddBookHandler)).Methods("POST")
	r.HandleFunc("/books/{id}", auth.AuthLibrarian(api.DeleteBookHandler)).Methods("DELETE")
	r.HandleFunc("/books/{id}", auth.AuthLibrarian(api.EditBookHandler)).Methods("PATCH")
	r.HandleFunc("/toggle_books/{id}", auth.AuthLibrarian(api.ToggleBookHandler)).Methods("PATCH")

	// Users
	r.HandleFunc("/login_user", api.LoginUserHandler).Methods("POST")

	r.HandleFunc("/student", auth.AuthAdmin(api.AddStudentHandler)).Methods("POST")
	r.HandleFunc("/student/{id}", auth.AuthAdmin(api.DeleteStudentHandler)).Methods("DELETE")
	r.HandleFunc("/all_students", api.GetAllStudentsHandler).Methods("GET")

	// Students
	r.HandleFunc("/add_user", auth.AuthLibrarian(api.DeleteStudentHandler)).Methods("POST")

	// Borrow
	r.HandleFunc("/borrow", auth.AuthLibrarian(api.AddBorrowHandler)).Methods("POST")
	r.HandleFunc("/borrow_fine", auth.AuthLibrarian(api.BorrowFineHandler)).Methods("GET")
	r.HandleFunc("/return_book/{id}", auth.AuthLibrarian(api.ReturnBookHandler)).Methods("POST")

	// Request
	r.HandleFunc("/request", auth.AuthStudent(api.RequestBorrowHandler)).Methods("POST")
	r.HandleFunc("/accept_request/{id}", auth.AuthStudent(api.AcceptRequestHandler)).Methods("POST")

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
