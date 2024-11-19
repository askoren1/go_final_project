package repository

//import ("fmt")

func (r *Repository) GetList() error {
	query := `;`

	if _, err := r.db.Exec(query); err != nil {
		return err
	}
 return nil
}
