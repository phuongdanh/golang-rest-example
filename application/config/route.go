package config

import (
	"net/http"

	"github.com/phuongdanh/golang-rest-example/application/controller"
)

var userController controller.User

func InitRoute() {
	http.HandleFunc("/users/list", userController.List)
	http.HandleFunc("/users/get", userController.Get)
	http.HandleFunc("/users/create", userController.List)
	http.HandleFunc("/users/update", userController.List)
	http.HandleFunc("/users/delete", userController.List)
}
