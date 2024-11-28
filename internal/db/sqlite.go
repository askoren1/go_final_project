package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func New() *sql.DB { //объявляем функцию для создания нового подключения к базе данных SQLite

	db, err := sql.Open("sqlite3", "scheduler.db") //Открываем соединение с базой данных SQLite, подготавливаем драйвер
	if err != nil {
		log.Fatal("init db", err)
	}

	if err := db.Ping(); err != nil { //Проверяем наличие ошибок при открытии соединения
		log.Fatal("ping db", err)
	}

	return db //Возвращаем указатель на открытое и проверенное соединение с базой данных
}

// Close закрывает соединение с базой данных.
func Close(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Printf("Error closing database: %v", err) // Логируем ошибку, но не паникуем
	}
}
