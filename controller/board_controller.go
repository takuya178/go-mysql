package controller

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
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
	userId := getToken(w, r)
	boardRes, err := tc.tu.GetAllBoards(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseData, err := json.Marshal(boardRes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(responseData)
}

func (tc *boardController) GetBoardById(w http.ResponseWriter, r *http.Request) {
	userId := getToken(w, r)
	boardId, err := strconv.ParseUint(r.URL.Query().Get("boardId"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid Board ID format", http.StatusBadRequest)
		return
	}
	boardRes, err := tc.tu.GetBoardById(userId, uint(boardId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseData, err := json.Marshal(boardRes)
	if err != nil {
		return
	}
	w.Write(responseData)
}

func (tc *boardController) CreateBoard(w http.ResponseWriter, r *http.Request) {
	userId := getToken(w, r)
	var board model.Board
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&board); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	board.UserID = userId
	boardRes, err := tc.tu.CreateBoard(board)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	responseData, err := json.Marshal(boardRes)
	if err != nil {
		return
	}
	w.Write(responseData)
}

func getToken(w http.ResponseWriter, r *http.Request) uint  {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
	}

	user_id, err := extractUserIDFromToken(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
	return user_id
}

func extractUserIDFromToken(tokenString string) (uint, error) {
	splitToken := strings.Split(tokenString, "Bearer ")
	token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid token claims")
	}

	userId, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("User ID not found in token")
	}

	return uint(userId), nil
}
