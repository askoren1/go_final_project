package handler

import (
	"fmt"
	"net/http"
	"time"

	nextdate "github.com/askoren1/go_final_project/internal/next_date"
)

const DateToday = "20240126" //эту строку нужно закомментировать для использования актуальной даты

// Функция NextDate для вычисления следующей даты выполнения задачи на основе заданной даты и правила повторения
func (h *Handler) NextDate(w http.ResponseWriter, r *http.Request) {

	nowTime, _ := time.Parse(Layout, DateToday) //эту строку нужно закомментировать для использования актуальной даты
	// nowTime := time.Now().Truncate(24 * time.Hour).UTC() //эту строку нужно раскомментировать для использования актуальной даты

	// Получение данных из запроса
	dateStr := r.FormValue("date")
	repeatStr := r.FormValue("repeat")

	if dateStr == "" {
		dateStr = time.Now().Format(Layout)
	}

	// Вычисление следующей даты
	date, err := nextdate.NextDate(nowTime, dateStr, repeatStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, date)
}
