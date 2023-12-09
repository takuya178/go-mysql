package controller

import (
	"encoding/json"
	"net/http"

	"github.com/t-shimpo/go-mysql-docker/app/model"
	"github.com/t-shimpo/go-mysql-docker/usecase"
)

type IUserController interface {
	SignUp(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
	}
	defer r.Body.Close()

	userRes, err := uc.uu.SignUp(user)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userRes)
}
