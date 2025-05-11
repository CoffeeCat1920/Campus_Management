package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	modals "what/internal/models"

	"github.com/gorilla/mux"
)

func (api *api) AddStudentHandler(w http.ResponseWriter, r *http.Request) {
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

	student := modals.NewUser(info.Name, info.Password)
	err = api.db.AddUser(student)
	if err != nil {
		http.Error(w, "Can't add student to db", http.StatusInternalServerError)
		fmt.Print(err)
		return
	}
}

func (api *api) RequestBorrowHandler(w http.ResponseWriter, r *http.Request) {
	info := &struct {
		StudentName string `json:"name"`
		BookISBN    string `json:"isbn"`
		Days        int    `json:"days"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		http.Error(w, "Wrong Format", http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	studentId, err := api.db.GetStudentUUIDFromName(info.StudentName)
	if err != nil {
		http.Error(w, "Can't get student of name "+info.StudentName, http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	bookId, err := api.db.GetBookUUIDFromISBN(info.BookISBN)
	if err != nil {
		http.Error(w, "Can't get book of ISBN "+info.BookISBN, http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	request := modals.NewRequest(studentId, bookId, info.Days)

	err = api.db.AddRequest(request)
	if err != nil {
		http.Error(w, "Can't add borrow request for "+info.StudentName, http.StatusBadRequest)
		fmt.Print(err)
		return
	}
}

func (api *api) DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["id"]

	err := api.db.DeleteStudentsFromUUID(uuid)
	if err != nil {
		http.Error(w, "Can't find user", http.StatusNotFound)
		fmt.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *api) GetAllStudentsHandler(w http.ResponseWriter, r *http.Request) {
	students, err := api.db.GetAllStudents()
	if err != nil {
		http.Error(w, "Can't get students", http.StatusInternalServerError)
		fmt.Printf("Can't get students cause, %v", err)
		return
	}

	fmt.Print(students)

	jsonData, err := json.Marshal(students)
	if err != nil {
		http.Error(w, "Can't Marshall json", http.StatusInternalServerError)
		fmt.Printf("Can't Marshall students cause, %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
