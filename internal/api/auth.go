package api

import (
	"encoding/json"
	"net/http"
	"what/internal/auth"
)

func (api *api) LoginHandler(w http.ResponseWriter, r *http.Request) {
	info := &struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}{}

	err := json.NewDecoder(r.Body).Decode(info)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
	}

	user, err := api.db.GetStudentFromName(info.Name)
	if err != nil {
		http.Error(w, "Invalid User Request", http.StatusBadRequest)
		print(err)
		return
	}

	if !user.CheckPassword(info.Password) {
		http.Error(w, "Wrong Password", http.StatusBadRequest)
		print(err)
		return
	}

	err = auth.SetToken(w, user)
	if !user.CheckPassword(info.Password) {
		http.Error(w, "Can't set cookie", http.StatusInternalServerError)
		print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
