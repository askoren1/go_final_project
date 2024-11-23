package handler

import (
	"net/http"
	"log"
     "encoding/json"
)

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, `{"error": "Не указан идентификатор"}`, http.StatusBadRequest)
        return
    }

	task, err := h.repo.GetTaskByID(id)
	if err != nil {
        log.Println("Error getting task:", err)
        http.Error(w, `{"error": "Ошибка получения задачи"}`, http.StatusInternalServerError)
        return
    }

	if task == nil {
        http.Error(w, `{"error": "Задача не найдена"}`, http.StatusNotFound)
        return
    }

	err = json.NewEncoder(w).Encode(task)
	if err != nil {
        log.Println("Error encoding JSON:", err)
        http.Error(w, `{"error": "Ошибка кодирования JSON"}`, http.StatusInternalServerError)
        return
    }
}


