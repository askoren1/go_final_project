package repository

import (
	"database/sql"
)

func (r *Repository) DeleteTask(id string) error {

	_, err := r.db.Exec("DELETE FROM scheduler WHERE Id = :Id",
		sql.Named("Id", id))

	return err
}
