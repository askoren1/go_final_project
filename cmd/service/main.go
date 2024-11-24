package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/askoren1/go_final_project/internal/db"
	"github.com/askoren1/go_final_project/internal/handler"
	"github.com/askoren1/go_final_project/internal/repository"
)

func main() {

	db := db.New()             //Инициализируем базу данных
	repo := repository.New(db) //Создаем репозиторий, который отвечает за взаимодействие с базой данных
	migration(repo)

	handler := handler.New(repo)

	r := chi.NewRouter()
	r.Post("/api/task", handler.AddTask)
	r.Get("/api/tasks", handler.GetList)
	r.Get("/api/nextdate", handler.NextDate)
	r.Get("/api/task", handler.GetTask)
	r.Put("/api/task", handler.UpdateTask)

	r.Handle("/*", http.FileServer(http.Dir("./web")))
	if err := http.ListenAndServe(":7540", r); err != nil {
		log.Fatal(err)
	}
}

func migration(repo *repository.Repository) { //функция для создания таблицы в базе данных, если она еще не существует
	appPath, err := os.Executable() //Получаем путь к исполняемому файлу приложения
	if err != nil {
		log.Fatal(err)
	}
	dbFile := filepath.Join(filepath.Dir(appPath), "scheduler.db") // Конструируем полный путь к файлу БД scheduler.db

	_, err = os.Stat(dbFile) //Проверяем, существует ли файл базы данных по указанному пути
	var install bool
	if err != nil {
		install = true
	}

	if install {
		if err := repo.CreateScheduler(); err != nil { //Вызываем метод CreateScheduler() у репозитория для создания таблицы
			log.Fatal(err)
		}
	} else {
		fmt.Println("База данных уже существует.")
	}
}
