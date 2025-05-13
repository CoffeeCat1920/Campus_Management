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

	w.WriteHeader(http.StatusOK)
}

func (api *api) EditStudentHandler(w http.ResponseWriter, r *http.Request) {
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

	_, err = api.db.GetStudentFromUUID(uuid)
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

func (api *api) GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["id"]

	student, err := api.db.GetStudentFromUUID(uuid)
	if err != nil {
		http.Error(w, "Can't Find User", http.StatusNotFound)
		fmt.Print(err)
		return
	}

	jsonData, err := json.Marshal(student)
	if err != nil {
		http.Error(w, "Can't Marshall json", http.StatusInternalServerError)
		fmt.Printf("Can't Marshall students cause, %v", err)
		return
	}

	fmt.Print(student)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (api *api) RequestBorrowHandler(w http.ResponseWriter, r *http.Request) {
	info := &struct {
		BookISBN string `json:"isbn"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		http.Error(w, "Wrong Format", http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	bookId, err := api.db.GetBookUUIDFromISBN(info.BookISBN)
	if err != nil {
		http.Error(w, "Can't get book of ISBN "+info.BookISBN, http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	studentId := r.Context().Value("uuid").(string)
	request := modals.NewRequest(studentId, bookId)

	err = api.db.AddRequest(request)
	if err != nil {
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

func (api *api) BooksExcept(w http.ResponseWriter, r *http.Request) {
	studentId := r.Context().Value("uuid").(string)

	books, err := api.db.GetAllBooksExcept(studentId)
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

func (api *api) NumberOfBorrowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId := vars["id"]

	count, err := api.db.NumberOfRentedBooks(studentId)
	fmt.Printf("\nCount: %v\n", count)
	if err != nil {
		http.Error(w, "Can't get number of borrowed books", http.StatusInternalServerError)
		fmt.Printf("Can't get books cause, %v", err)
		return
	}

	jsonData, err := json.Marshal(count)
	if err != nil {
		http.Error(w, "Can't Marshall json", http.StatusInternalServerError)
		fmt.Printf("Can't Marshall nob cause, %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
