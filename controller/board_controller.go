package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/t-shimpo/go-mysql-docker/app/model"
	"github.com/t-shimpo/go-mysql-docker/usecase"
)

type IBoardController interface {
	GetAllBoards(w http.ResponseWriter, r *http.Request)
	GetBoardById(w http.ResponseWriter, r *http.Request)
	CreateBoard(w http.ResponseWriter, r *http.Request)
}

type boardController struct {
	tu usecase.IBoardUsecase
}

func NewBoardController(tu usecase.IBoardUsecase) IBoardController {
	return &boardController{tu}
}

func (tc *boardController) GetAllBoards(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	userId, err := extractUserIDFromToken(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	boardRes, err := tc.tu.GetAllBoards(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(boardRes)
}

func extractUserIDFromToken(tokenString string) (uint, error) {
	token := "Bearer " + tokenString
	userId, err := strconv.ParseUint(token[7:], 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(userId), nil
}
