package db

import (
	log "github.com/sirupsen/logrus"
	"pretty-deadlines/internal/models"
)

func (db *Database) Insert(deadline models.Deadline) error {
	query := "INSERT INTO deadlines (title, description, due_date) VALUES (?, ?, ?)"
	_, err := db.Exec(query, deadline.Title, deadline.Description, deadline.DueDate)

	if err != nil {
		log.Info("Deadline inserted")
	}
	return err
}
