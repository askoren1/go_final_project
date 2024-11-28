package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"

	"github.com/askoren1/go_final_project/internal/db"
	"github.com/askoren1/go_final_project/internal/handler"
	"github.com/askoren1/go_final_project/internal/repository"
)

func main() {

	dbConn := db.New() //Инициализируем базу данных
	defer db.Close(dbConn)
	repo := repository.New(dbConn) //Создаем репозиторий, который отвечает за взаимодействие с базой данных
	db.Migration(repo)

	handler := handler.New(repo)

	r := chi.NewRouter()
	r.Post("/api/task", handler.AddTask)
	r.Post("/api/task/done", handler.MarkTaskDone)
	r.Get("/api/tasks", handler.GetList)
	r.Get("/api/nextdate", handler.NextDate)
	r.Get("/api/task", handler.GetTask)
	r.Put("/api/task", handler.UpdateTask)
	r.Delete("/api/task", handler.DeleteTask)
	r.Handle("/*", http.FileServer(http.Dir("./web")))

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "8080" // Порт по умолчанию
	}
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
