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

	user, err := api.db.GetUserFromName(info.Name)
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

func (api *api) LoginStudentDataHandler(w http.ResponseWriter, r *http.Request) {
	claims, err := api.auth.StudentAuthData(r)

	if err != nil {
		http.Error(w, "Not LoggedIn", http.StatusUnauthorized)
		fmt.Print(err)
		return
	}

	data := &struct {
		UUID string
		Name string
		Type int
	}{
		UUID: claims.UUID,
		Name: claims.Name,
		Type: claims.Type,
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		http.Error(w, "Can't marshall data", http.StatusInternalServerError)
		fmt.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (api *api) LoginLibrarianDataHandler(w http.ResponseWriter, r *http.Request) {
	claims, err := api.auth.LibrarianAuthData(r)

	if err != nil {
		http.Error(w, "Not LoggedIn", http.StatusUnauthorized)
		fmt.Print(err)
		return
	}

	data := &struct {
		UUID string
		Name string
		Type int
	}{
		UUID: claims.UUID,
		Name: claims.Name,
		Type: claims.Type,
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		http.Error(w, "Can't marshall data", http.StatusInternalServerError)
		fmt.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (api *api) LoginAdminDataHandler(w http.ResponseWriter, r *http.Request) {
	err := api.auth.AdminAuthData(r)

	if err != nil {
		http.Error(w, "Not LoggedIn", http.StatusUnauthorized)
		fmt.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *api) LoginAdminHandler(w http.ResponseWriter, r *http.Request) {
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

func (api *api) LogoutUserHandler(w http.ResponseWriter, r *http.Request) {
	api.auth.ClearUserToken(w)
}
