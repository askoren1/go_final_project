package repository

func (r *Repository) AddTask() error {
	query := `;`

	if _, err := r.db.Exec(query); err != nil {
		return err
	}
 return nil
}



