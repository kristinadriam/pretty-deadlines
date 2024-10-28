package main

import (
	"os"
	"pretty-deadlines/internal/db/deadline"
	"pretty-deadlines/internal/models"
	"time"
)

func main() {
	// test
	database, err := db.InitDb()
	if err != nil {
		os.Exit(1)
	}

	database.Insert(models.Deadline{Title: "cool", Description: "coooool", DueDate: time.Now().AddDate(0, 0, 7)})
}
