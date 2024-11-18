package repository

import ("fmt")

func (r *Repository) CreateScheduler() error {
	query := `CREATE TABLE scheduler (
	   id INTEGER PRIMARY KEY AUTOINCREMENT,
	   date CHAR(8) NOT NULL DEFAULT "",
	   title VARCHAR(256) NOT NULL DEFAULT "",
	   comment TEXT NOT NULL DEFAULT "",
	   repeat VARCHAR(128) NOT NULL DEFAULT "");`

	if _, err := r.db.Exec(query); err != nil {
		return err
	}
	fmt.Println("База данных и таблицы созданы.")

	return nil
}
