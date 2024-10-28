package db

import (
	"fmt"
	"pretty-deadlines/internal/models"
	"time"

	log "github.com/sirupsen/logrus"
)

func (db *Database) GetAllDeadlines() ([]models.Deadline, error) {
	query := "SELECT title, description, deadline FROM deadlines"
	rows, err := db.Db().Query(query)
	if err != nil {
		log.Error("Failed to get all deadlines", err)
		return nil, fmt.Errorf("Failed to execute query: %v", err)
	}
	defer rows.Close()

	var deadlines []models.Deadline
	for rows.Next() {
		var dl models.Deadline
		var dueDateStr string

		if err := rows.Scan(&dl.Title, &dl.Description, &dueDateStr); err != nil {
			log.Error("Failed to scan row: ", err)
			return nil, fmt.Errorf("Failed to scan row: %v", err)
		}

		dl.DueDate, err = time.Parse("2006-01-02T15:04:05Z", dueDateStr)
		if err != nil {
			log.Error("Failed to parse due date: ", err)
			return nil, fmt.Errorf("Failed to parse due date: %v", err)
		}

		deadlines = append(deadlines, dl)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over rows: %v", err)
	}

	log.Info("Get all deadlines")
	return deadlines, nil
}
