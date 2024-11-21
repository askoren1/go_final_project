package repository

//import ("fmt")

func (r *Repository) GetList() error {
	query := `SELECT * FROM scheduler;`

	if _, err := r.db.Exec(query); err != nil {
		return err
	}
 return nil
}
