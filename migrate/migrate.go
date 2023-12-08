package main

import (
	"fmt"

	"github.com/t-shimpo/go-mysql-docker/app/models"
	"github.com/t-shimpo/go-mysql-docker/db"
)

func main() {
	dbConnect := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConnect)
	dbConnect.AutoMigrate(&models.User{}, &models.Board{})
}