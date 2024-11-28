package handler

import (
	"fmt"
	nextdate "github.com/askoren1/go_final_project/internal/next_date"
	"net/http"
	"time"
)

const DateToday = "20240126" //эту строку нужно закомментировать для использования актуальной даты

// Функция NextDate для вычисления следующей даты выполнения задачи на основе заданной даты и правила повторения
func (h *Handler) NextDate(w http.ResponseWriter, r *http.Request) {
	layout := "20060102"                        //эту строку нужно закомментировать для использования актуальной даты
	nowTime, _ := time.Parse(layout, DateToday) //эту строку нужно закомментировать для использования актуальной даты
	// nowTime := time.Now().Truncate(24 * time.Hour).UTC() //эту строку нужно раскомментировать для использования актуальной даты

	// Получение данных из запроса
	dateStr := r.FormValue("date")
	repeatStr := r.FormValue("repeat")

	if dateStr == "" {
		dateStr = time.Now().Format("20060102")
	}

	// Вычисление следующей даты
	date, err := nextdate.NextDate(nowTime, dateStr, repeatStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, date)
}
