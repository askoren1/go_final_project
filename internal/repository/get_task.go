package repository

import (
	"database/sql"
	"github.com/askoren1/go_final_project/internal/models"
	"strconv"
)

func (r *Repository) GetTaskByID(id string) (*models.Task2, error) {

	var task models.Task2
	var taskID int64

	row := r.db.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?", id)
	err := row.Scan(&taskID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	task.ID = strconv.FormatInt(taskID, 10)

	return &task, nil
}
