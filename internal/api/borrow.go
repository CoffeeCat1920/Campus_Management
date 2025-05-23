package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	modals "what/internal/models"

	"github.com/gorilla/mux"
)

func (api *api) AddBorrowHandler(w http.ResponseWriter, r *http.Request) {
	info := &struct {
		BookId string `json:"bookid"`
		UserId string `json:"userid"`
		Days   int    `json:"days"`
	}{}

	err := json.NewDecoder(r.Body).Decode(info)
	if err != nil {
		http.Error(w, "Format Error", http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	bookId, err := api.db.GetBookUUIDFromISBN(info.BookId)
	if err != nil {
		http.Error(w, "Can't find book", http.StatusNotFound)
		fmt.Printf("Can't get book cause %v\n", err)
		return
	}

	userId, err := api.db.GetStudentUUIDFromName(info.UserId)
	if err != nil {
		http.Error(w, "Can't find user", http.StatusNotFound)
		fmt.Printf("Can't find user cause, %v\n", err)
		return
	}

	borrow := modals.NewBorrow(bookId, userId, info.Days)
	borrow.Returned = false

	err = api.db.AddBorrow(borrow)
	if err != nil {
		http.Error(w, "Can't add borrow in db", http.StatusInternalServerError)
		fmt.Printf("Can't add borrow in db %v\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *api) BorrowFineHandler(w http.ResponseWriter, r *http.Request) {
	info := &struct {
		BorrowId string `json:"borrowid"`
	}{}

	err := json.NewDecoder(r.Body).Decode(info)
	if err != nil {
		http.Error(w, "Format Error", http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	fine, err := api.db.BorrowFine(info.BorrowId, 10)
	if err != nil {
		http.Error(w, "Can't Calculate Fine", http.StatusInternalServerError)
		fmt.Print(err)
		return
	}

	jsonData, err := json.Marshal(fine)
	if err != nil {
		http.Error(w, "Can't Marshall json", http.StatusInternalServerError)
		fmt.Printf("Can't Marshall fine cause, %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (api *api) GetBorrowByUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId := vars["id"]

	borrows, err := api.db.GetBorrowsByUserWithBookName(studentId)
	if err != nil {
		http.Error(w, "Can't get borrow", http.StatusInternalServerError)
		fmt.Printf("Can't get borrow cause, %v", err)
		return
	}

	jsonData, err := json.Marshal(borrows)
	if err != nil {
		http.Error(w, "Can't Marshall json", http.StatusInternalServerError)
		fmt.Printf("Can't Marshall books cause, %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (api *api) GetFineByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId := vars["id"]

	fine, err := api.db.GetFineByUser(studentId)
	if err != nil {
		http.Error(w, "Can't Calculate Fine", http.StatusInternalServerError)
		fmt.Print(err)
		return
	}

	fmt.Printf("\nfine:%v\n", fine)

	jsonData, err := json.Marshal(fine)
	if err != nil {
		http.Error(w, "Can't Marshall json", http.StatusInternalServerError)
		fmt.Printf("Can't Marshall fine cause, %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
