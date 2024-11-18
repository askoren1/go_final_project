package main

import (
	"fmt"
	"os"
	"path/filepath"
	//"database/sql"
	"github.com/go-chi/chi/v5"
	"log"
	//_ "modernc.org/sqlite"
	"net/http"

	"github.com/askoren1/go_final_project/db"
	"github.com/askoren1/go_final_project/internal/handler/add_task"
	"github.com/askoren1/go_final_project/repository"
	)

func main() {

	db := db.New()
	rep := repository.New(db)
	migration(rep)

	r := chi.NewRouter()
	r.Post("/api/task", add_task.AddTask)

	r.Handle("/*", http.FileServer(http.Dir("./web")))
	if err := http.ListenAndServe(":7540", r); err != nil {
		log.Fatal(err)
	}
}

func migration(rep *repository.Repository) {
	appPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dbFile := filepath.Join(filepath.Dir(appPath), "scheduler.db")
	_, err = os.Stat(dbFile)
	fmt.Println(dbFile)

	var install bool
	if err != nil {
		install = true
	}
	fmt.Println(install)

	if install {
		if err := rep.CreateScheduler(); err != nil {
			log.Fatal(err)
		}
	} else {fmt.Println("База данных уже существует.")}
}
