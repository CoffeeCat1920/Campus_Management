package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	modals "what/internal/models"

	"github.com/gorilla/mux"
)

func (api *api) AddLibrarianHandler(w http.ResponseWriter, r *http.Request) {
	info := &struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		http.Error(w, "Wrong Format", http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	librarian := modals.NewLibrarian(info.Name, info.Password)
	err = api.db.AddUser(librarian)
	if err != nil {
		http.Error(w, "Can't add librarian to db", http.StatusInternalServerError)
		fmt.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *api) GetLibrarianHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["id"]

	librarian, err := api.db.GetLibrarianFromUUID(uuid)
	if err != nil {
		http.Error(w, "Can't Find Librarian", http.StatusNotFound)
		fmt.Print(err)
		return
	}

	jsonData, err := json.Marshal(librarian)
	if err != nil {
		http.Error(w, "Can't Marshall json", http.StatusInternalServerError)
		fmt.Printf("Can't Marshall librarian cause, %v", err)
		return
	}

	fmt.Print(librarian)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (api *api) GetAllLibrariansHandler(w http.ResponseWriter, r *http.Request) {
	librarians, err := api.db.GetAllLibrarians()
	if err != nil {
		http.Error(w, "Can't get students", http.StatusInternalServerError)
		fmt.Printf("Can't get students cause, %v", err)
		return
	}

	fmt.Print(librarians)

	jsonData, err := json.Marshal(librarians)
	if err != nil {
		http.Error(w, "Can't Marshall json", http.StatusInternalServerError)
		fmt.Printf("Can't Marshall students cause, %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
func (api *api) EditLibrarianHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["id"]

	info := &struct {
		Name        string `json:"name"`
		NewPassword string `json:"new_password"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		http.Error(w, "Wrong Format", http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	_, err = api.db.GetLibrarianFromUUID(uuid)
	if err != nil {
		http.Error(w, "Can't Find User", http.StatusNotFound)
		fmt.Print(err)
		return
	}

	err = api.db.UpdateUserFromUUID(uuid, info.Name, info.NewPassword)
	if err != nil {
		http.Error(w, "Can't Edit user", http.StatusInternalServerError)
		fmt.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *api) DeleteLibrarianHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["id"]

	err := api.db.DeleteLibrarianFromUUID(uuid)
	if err != nil {
		http.Error(w, "Can't find user", http.StatusNotFound)
		fmt.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *api) AcceptRequestHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["id"]

	info := &struct {
		Days int
	}{}

	err := json.NewDecoder(r.Body).Decode(&info)

	request, err := api.db.GetRequest(uuid)
	if err != nil {
		http.Error(w, "Can't get the request from database", http.StatusBadRequest)
		fmt.Printf("Can't get the request from database cause, %v", err)
		return
	}

	err = api.db.DeleteRequest(uuid)
	if err != nil {
		http.Error(w, "Can't delete the request from database", http.StatusBadRequest)
		fmt.Printf("Can't delete the request from database, %v", err)
		fmt.Print(err)
		return
	}

	rentedBooks, err := api.db.NumberOfRentedBooks(request.UserId)
	if err != nil {
		http.Error(w, "\nCan't get number of rented books from database", http.StatusBadRequest)
		fmt.Printf("\nCan't get number of rented books from database, %v", err)
		return
	}

	if rentedBooks >= 3 {
		http.Error(w, "number of rented book is larger", http.StatusBadRequest)
		fmt.Print("number of rented book is larger it is")
		return
	}

	borrow := modals.NewBorrow(request.BookId, request.UserId, info.Days)

	err = api.db.AddBorrow(borrow)
	if err != nil {
		http.Error(w, "\nCan't add the borrow to database", http.StatusBadRequest)
		fmt.Print(err)
		fmt.Printf("\nCan't add the borrow to database %v", err)
		return
	}

	err = api.db.IncreaseStudentRented(request.UserId)
	if err != nil {
		http.Error(w, "\nCan't increase number of rented books in database", http.StatusBadRequest)
		fmt.Print(err)
		fmt.Printf("\nCan't increase number of rented books in database, %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *api) DeclineRequestHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["id"]

	err := api.db.DeleteRequest(uuid)
	if err != nil {
		http.Error(w, "Can't delete the request from database", http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *api) ReturnBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["id"]

	err := api.db.DeleteBorrow(uuid)
	if err != nil {
		http.Error(w, "Can't delete the borrow from database", http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	request, err := api.db.GetRequest(uuid)
	if err != nil {
		http.Error(w, "Can't get request from db", http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	err = api.db.DeleteBorrow(request.UserId)
	if err != nil {
		http.Error(w, "Can't decrease number of rented books in database", http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
