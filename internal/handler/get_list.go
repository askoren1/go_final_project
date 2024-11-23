package handler

import (
	"net/http"
	"encoding/json"
	"github.com/askoren1/go_final_project/internal/models"
	"log"
)

func (h *Handler) GetList(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	tasks, err := h.repo.GetList()


	if err != nil {
		log.Println("Error getting tasks:", err)
		http.Error(w, `{"error": "Ошибка получения списка задач"}`, http.StatusInternalServerError)
		return
	}

	if tasks == nil {
		tasks = []models.Task2{}
	}

	resp := struct {
		Tasks []models.Task2 `json:"tasks"`
	}{
		Tasks: tasks,
	}

	err = json.NewEncoder(w).Encode(resp)

	if err != nil {
		log.Println("Error encoding JSON:", err)
		http.Error(w, `{"error": "Ошибка кодирования JSON"}`, http.StatusInternalServerError)
		return
	}

}


