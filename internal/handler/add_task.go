package handler

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type Task struct {

	 Date string `json:"data"`
     Title string `json:"title"`     
	 Comment string `json:"comment"`
	 Repeat string `json:"repeat"`
}

func (h *Handler) AddTask(w http.ResponseWriter, r *http.Request) {
	var t Task 
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	id, err := h.repo.AddTask(t.Date, t.Title, t.Comment, t.Repeat)
	
	if err != nil {
 		http.Error(w, err.Error(), http.StatusBadRequest)
 		return
	}


     response := map[string]string{"id": fmt.Sprintf("%d", id)}
	 w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	 json.NewEncoder(w).Encode(response)

}
