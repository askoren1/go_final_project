package repository

import (
	"database/sql"
	"fmt"
)

func (r *Repository) DeleteTask(id string) error {
	res, err := r.db.Exec("DELETE FROM scheduler WHERE Id = :Id", sql.Named("Id", id))
	if err != nil {
		return fmt.Errorf("ошибка удаления задачи: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения количества затронутых строк: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("задача с id %s не найдена", id)
	}

	return nil
}
