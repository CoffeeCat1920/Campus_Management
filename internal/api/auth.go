package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	modals "what/internal/models"
)

func (api *api) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	info := &struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}{}

	err := json.NewDecoder(r.Body).Decode(info)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	user, err := api.db.GetStudentFromName(info.Name)
	if err != nil {
		http.Error(w, "Invalid User Request", http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	if !user.CheckPassword(info.Password) {
		http.Error(w, "Wrong Password", http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	err = api.auth.SetUserToken(w, user)
	if err != nil {
		http.Error(w, "Can't set cookie", http.StatusInternalServerError)
		fmt.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *api) LoginAdminDataHandle(w http.ResponseWriter, r *http.Request) {
	err := api.auth.AdminAuthData(r)

	if err != nil {
		http.Error(w, "Not LoggedIn", http.StatusUnauthorized)
		fmt.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *api) LoginAdminHandle(w http.ResponseWriter, r *http.Request) {
	info := &struct {
		Password string `json:"password"`
	}{}

	err := json.NewDecoder(r.Body).Decode(info)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	admin, err := modals.NewAdmin()
	if err != nil {
		http.Error(w, "Admin can be initialized", http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	err = admin.CheckPassword(info.Password)

	if err != nil {
		http.Error(w, "Invalid Password", http.StatusUnauthorized)
		fmt.Print(err)
		return
	}

	err = api.auth.SetAdminToken(w)
	if err != nil {
		http.Error(w, "Can't set cookie", http.StatusInternalServerError)
		fmt.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *api) LogoutAdminHandler(w http.ResponseWriter, r *http.Request) {
	api.auth.ClearAdminToken(w)
}
