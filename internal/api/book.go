package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	modals "what/internal/models"

	"github.com/gorilla/mux"
)

type BookInfo struct {
	ISBN      string `json:"isbn"`
	Name      string `json:"name"`
	Available bool   `json:"available"`
}

func (api *api) AddBookHandler(w http.ResponseWriter, r *http.Request) {
	var info BookInfo
	err := json.NewDecoder(r.Body).Decode(&info)

	if err != nil {
		http.Error(w, "Form Error", http.StatusBadRequest)
		fmt.Print(err.Error())
		return
	}

	book := modals.NewBook(info.ISBN, info.Name)

	db := api.db
	err = db.AddBook(book)

	if err != nil {
		http.Error(w, "Can't get book from db", http.StatusInternalServerError)
		fmt.Print(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (api *api) DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isbn := vars["id"]

	uuid, err := api.db.GetBookUUIDFromISBN(isbn)
	if err != nil {
		http.Error(w, "Can't find the book", http.StatusInternalServerError)
		fmt.Print(err.Error())
	}

	err = api.db.DeleteBook(uuid)
	if err != nil {
		http.Error(w, "Can't Delete the book", http.StatusInternalServerError)
		fmt.Print(err.Error())
	}
}

func (api *api) EditBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isbn := vars["id"]

	info := &struct {
		Name string `json:"name"`
	}{}
	err := json.NewDecoder(r.Body).Decode(info)

	if err != nil {
		http.Error(w, "Bad Data Format", http.StatusBadRequest)
		fmt.Print(err.Error())
		return
	}

	uuid, err := api.db.GetBookUUIDFromISBN(isbn)
	if err != nil {
		http.Error(w, "Can't find book", http.StatusBadRequest)
		fmt.Print(err.Error())
		return
	}

	err = api.db.UpdateBookName(uuid, info.Name)
	if err != nil {
		http.Error(w, "Can't update book", http.StatusBadRequest)
		fmt.Print(err.Error())
		return
	}

	fmt.Print(info.Name)
	w.WriteHeader(http.StatusOK)
}

func (api *api) ToggleBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isbn := vars["id"]

	uuid, err := api.db.GetBookUUIDFromISBN(isbn)
	if err != nil {
		http.Error(w, "Can't find book", http.StatusBadRequest)
		fmt.Print(err.Error())
		return
	}

	err = api.db.ToggleBookAvailiablity(uuid)
	if err != nil {
		http.Error(w, "Can't toggle book", http.StatusBadRequest)
		fmt.Print(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *api) GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := api.db.GetAllBooks()
	if err != nil {
		http.Error(w, "Can't get books", http.StatusInternalServerError)
		fmt.Printf("Can't get books cause, %v", err)
		return
	}

	jsonData, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "Can't Marshall json", http.StatusInternalServerError)
		fmt.Printf("Can't Marshall books cause, %v", err)
		return
	}

	fmt.Print(books)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
