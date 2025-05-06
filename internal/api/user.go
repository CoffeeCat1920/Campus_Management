package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	modals "what/internal/models"
)

func (api *api) AddUserHandler(w http.ResponseWriter, r *http.Request) {
	info := &struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		http.Error(w, "Wrong Format", http.StatusBadRequest)
		fmt.Print(err)
	}

	user := modals.NewUser(info.Name, info.Password)
	err = api.db.AddUser(user)
	if err != nil {
		http.Error(w, "Can't add user to db", http.StatusInternalServerError)
		fmt.Print(err)
	}
}

func (api *api) RequestBorrowHandler(w http.ResponseWriter, r *http.Request) {
	info := &struct {
		UserName string `json:"username"`
		BookISBN string `json:"isbn"`
		Days     int    `json:"days"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		http.Error(w, "Wrong Format", http.StatusBadRequest)
		fmt.Print(err)
	}

	userId, err := api.db.GetUserUUIDFromName(info.UserName)
	if err != nil {
		http.Error(w, "Can't Get User of name "+info.UserName, http.StatusBadRequest)
		fmt.Print(err)
	}

	bookId, err := api.db.GetBookUUIDFromISBN(info.BookISBN)
	if err != nil {
		http.Error(w, "Can't Get User of bookid "+info.BookISBN, http.StatusBadRequest)
		fmt.Print(err)
	}

	request := modals.NewRequest(userId, bookId, info.Days)

	err = api.db.AddRequest(request)
	if err != nil {
		http.Error(w, "Can't Add request of user "+info.UserName, http.StatusBadRequest)
		fmt.Print(err)
	}
}
