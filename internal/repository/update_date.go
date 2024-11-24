package repository

import (
	"database/sql"
)

func (r *Repository) UpdateDate(id, date string) error {

	_, err := r.db.Exec("UPDATE scheduler SET Date = :Date WHERE Id = :Id",
		sql.Named("Date", date),
		sql.Named("Id", id))

	return err
}
