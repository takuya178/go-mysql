package router

import (
	"github.com/gorilla/mux"
	"github.com/t-shimpo/go-mysql-docker/controller"
)

func NewRouter(userController controller.IUserController) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/signup", userController.SignUp).Methods("POST")
	r.HandleFunc("/login", userController.LogIn).Methods("POST")
	return r
}