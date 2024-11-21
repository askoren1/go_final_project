package handler

import (
	"fmt"
	"net/http"
	nextdate "github.com/askoren1/go_final_project/internal/next_date"
	"time"
)

const DateToday = "20240126"

func (h *Handler) NextDate(w http.ResponseWriter, r *http.Request) {
	layout := "20060102"
	nowTime, _ := time.Parse(layout, DateToday)
	dateStr := r.FormValue("date")
	repeatStr := r.FormValue("repeat")

	if dateStr == "" {
		dateStr = DateToday
	} 

	

	date, err := nextdate.NextDate(nowTime, dateStr, repeatStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, date)
}