package repository

import ("fmt")

type Task2 struct {
	ID    string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func (r *Repository) GetList() ([]Task2, error) {
	query := `SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date ASC;`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

var tasks []Task2

for rows.Next() {
	task := Task2{}

	err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	tasks = append(tasks, task)
}

if err := rows.Err(); err != nil {
	fmt.Println(err)
	return nil, err
}

 return tasks, err
}
