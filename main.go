package main

import (
	"fmt"

	"github.com/askoren1/go_final_project/db"
	"github.com/askoren1/go_final_project/http"
)

func main() {
	fmt.Println("Пример приложения")
	http.Print1()
	db.Print2()

}
