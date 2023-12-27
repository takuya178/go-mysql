package router

import (
	"github.com/gorilla/mux"
	"github.com/t-shimpo/go-mysql-docker/controller"
)

func NewRouter(userController controller.IUserController, boardController controller.IBoardController) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/signup", userController.SignUp).Methods("POST")
	r.HandleFunc("/login", userController.LogIn).Methods("POST")

	r.HandleFunc("/boards", boardController.GetAllBoards).Methods("GET")
	r.HandleFunc("/boards", boardController.CreateBoard).Methods("POST")
	return r
}