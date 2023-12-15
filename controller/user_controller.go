package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/t-shimpo/go-mysql-docker/app/model"
	"github.com/t-shimpo/go-mysql-docker/usecase"
)

type IUserController interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	LogIn(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userRes)
}

func (uc *userController) LogIn(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteStrictMode

	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ログインに成功しました"))
}
