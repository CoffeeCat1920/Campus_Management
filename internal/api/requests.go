package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (api *api) GetRequestByUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId := vars["id"]

	recipes, err := api.db.GetRequestsByUse(studentId)
	if err != nil {
		fmt.Print(err)
		http.Error(w, "Can't find any requests for the user", http.StatusNotFound)
		return
	}

	jsonData, err := json.Marshal(recipes)
	if err != nil {
		fmt.Print(err)
		http.Error(w, "Can't Marshall the requests", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
