package repository

import (
	"database/sql"
	"fmt"
)

func (r *Repository) UpdateTask(id, date, title, comment, repeat string) error {
	res, err := r.db.Exec("UPDATE scheduler SET Date = :Date, Title = :Title, Comment = :Comment, Repeat = :Repeat WHERE Id = :Id",
		sql.Named("Date", date),
		sql.Named("Title", title),
		sql.Named("Comment", comment),
		sql.Named("Repeat", repeat),
		sql.Named("Id", id))
	if err != nil {
		return fmt.Errorf("ошибка обновления задачи: %w", err)
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
