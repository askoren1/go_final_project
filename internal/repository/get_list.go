package repository

import (
	"fmt"
	"github.com/askoren1/go_final_project/internal/models"
)

func (r *Repository) GetList() ([]models.Task, error) {
	query := `SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date ASC;`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		task := models.Task{}

		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return tasks, err
}
