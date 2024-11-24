package repository

import (
	"database/sql"
	"fmt"
)

func (r *Repository) DeleteTask(id string) error {

	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM scheduler WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		return fmt.Errorf("ошибка проверки существования задачи: %w", err)
	}

	if !exists {
		return fmt.Errorf("задача с id %s не найдена", id)
	}

	_, err = r.db.Exec("DELETE FROM scheduler WHERE Id = :Id",
		sql.Named("Id", id))

	return err
}
