package main

import (
	"fmt"
	"net/http"

	"github.com/t-shimpo/go-mysql-docker/controller"
	"github.com/t-shimpo/go-mysql-docker/db"
	"github.com/t-shimpo/go-mysql-docker/repository"
	"github.com/t-shimpo/go-mysql-docker/router"
	"github.com/t-shimpo/go-mysql-docker/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	boardRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	boardUsecase := usecase.NewBoardUsecase(boardRepository)
	userController := controller.NewUserController(userUsecase)
	boardController := controller.NewBoardController(boardUsecase)
	r := router.NewRouter(userController, boardController)

	serverAddr := ":8080"
	fmt.Printf("Server is listening on %s...\n", serverAddr)
	err := http.ListenAndServe(serverAddr, r)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}