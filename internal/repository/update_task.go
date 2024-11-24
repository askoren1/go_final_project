package repository

import (
	"database/sql"
	"fmt"
)

func (r *Repository) UpdateTask(id, date, title, comment, repeat string) error {

	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM scheduler WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		return fmt.Errorf("ошибка проверки существования задачи: %w", err)
	}

	if !exists {
		return fmt.Errorf("задача с id %s не найдена", id)
	}

	_, err = r.db.Exec("UPDATE scheduler SET Date = :Date, Title = :Title, Comment = :Comment, Repeat = :Repeat WHERE Id = :Id",
		sql.Named("Date", date),
		sql.Named("Title", title),
		sql.Named("Comment", comment),
		sql.Named("Repeat", repeat),
		sql.Named("Id", id))

	return err
}
