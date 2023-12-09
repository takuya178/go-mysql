package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/t-shimpo/go-mysql-docker/app/model"
	"github.com/t-shimpo/go-mysql-docker/db"
)

func main() {
	dbConnect := db.NewDB()
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Failed to load .env: %v\n", err)
	}
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConnect)
	dbConnect.AutoMigrate(&model.User{}, &model.Board{})
}