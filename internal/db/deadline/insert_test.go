package db

import (
	"database/sql"
	"pretty-deadlines/internal/models"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	defer db.Close()

	_, err = db.Exec("CREATE TABLE deadlines (title TEXT, description TEXT, deadline DATETIME)")
	assert.NoError(t, err)

	testDb := &Database{db: db}

	deadline := models.Deadline{
		Title:       "Test Title",
		Description: "Test Description",
		DueDate:     time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC),
	}

	err = testDb.Insert(deadline)
	assert.NoError(t, err)

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM deadlines WHERE title = ?", deadline.Title).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)
}
