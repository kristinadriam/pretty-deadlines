package db

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func DeleteDeadlineById(db *sql.DB, id int) error {
	query := `DELETE FROM deadlines WHERE id = ?`

	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Failed to delete deadline: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Failed to retrieve affected rows: %v", err)
	}

	if rowsAffected == 0 {
		log.Info("No deadline found with the provided id.")
	} else {
		log.Info("Deadline with id \"%d\" successfully deleted!\n", id)
	}

	return nil
}
