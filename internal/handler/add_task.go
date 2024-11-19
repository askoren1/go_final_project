package handler

import (
	"net/http"
	"encoding/json"
)

type Task struct {
     Title string `json:"title"`
     Date string `json:"data"`

}

func (h *Handler) AddTask(w http.ResponseWriter, r *http.Request) {
	var t Task 
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	//Провеверка на входящие данные
	// // Если ошибка, возвращаем 404 ошибку
	// if err := h.repo.AddTask(); err != nil {
	// 	// Возвращаем 500 ошибку

	// }
    // // Возвращаем OK
	result, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(result)



}
