package handler

import (
	"fmt"
	"net/http"
	nextdate "github.com/askoren1/go_final_project/internal/next_date"
	"time"
)

func (h *Handler) NextDate(w http.ResponseWriter, r *http.Request) {
	date, err := nextdate.NextDate(time.Now(), "dateStr", "repeat")
	if err != nil {
		//TODO вернуть ошибку
	}
	fmt.Fprintln(w, date)
	w.Write([]byte("date"))

}