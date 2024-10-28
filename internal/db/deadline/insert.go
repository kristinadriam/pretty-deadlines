package db

import (
	log "github.com/sirupsen/logrus"

	"pretty-deadlines/internal/models"
)

func (db *Database) Insert(deadline models.Deadline) error {
	query := "INSERT INTO deadlines (title, description, deadline) VALUES (?, ?, ?)"
	dueDateStr := deadline.DueDate.Format("2006-01-02 15:04:05")
	args := make([]interface{}, 0)
	args = append(args, deadline.Title, deadline.Description, dueDateStr)
	_, err := db.Db().Exec(query, args...)

	if err != nil {
		log.Error("Failed to insert deadline: ", err)	
		return err
	}

	log.Info("Deadline inserted")
	return nil
}
