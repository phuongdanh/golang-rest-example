package controller

import (
	"encoding/json"
	"net/http"

	"github.com/phuongdanh/golang-rest-example/application/model"
)

type User struct {
	Name string
}

func (this *User) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// query in database
	listUsers := model.UserExamples()
	json.NewEncoder(w).Encode(listUsers)
}

func (this *User) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// query in database
	listUsers := model.UserExamples()
	json.NewEncoder(w).Encode(listUsers[0])
}
