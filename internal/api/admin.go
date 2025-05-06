package api

import (
	"fmt"
	"net/http"
	modals "what/internal/models"

	"github.com/gorilla/mux"
)

func (api *api) AcceptRequestHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["id"]

	request, err := api.db.GetRequest(uuid)
	if err != nil {
		http.Error(w, "Can't get the request from database", http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	err = api.db.DeleteRequest(uuid)
	if err != nil {
		http.Error(w, "Can't delete the request from database", http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	rentedBooks, err := api.db.NumberOfRentedBooks(request.UserId)
	if err != nil {
		http.Error(w, "Can't get number of rented books from database", http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	if rentedBooks >= 3 {
		http.Error(w, ErrMaxRentLimit.Error(), http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	borrow := modals.NewBorrow(request.BookId, request.UserId, request.Days)

	err = api.db.AddBorrow(borrow)
	if err != nil {
		http.Error(w, "Can't add the borrow to database", http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	err = api.db.IncreaseRented(request.UserId)
	if err != nil {
		http.Error(w, "Can't increase number of rented books in database", http.StatusBadRequest)
		fmt.Print(err)
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
