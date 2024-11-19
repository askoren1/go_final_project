package repository

import (
		"database/sql"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {   //создаем конструктор для структуры Repository
	return &Repository {
		db: db,                   //возвращаем указатель на созданную структуру
	}
}



