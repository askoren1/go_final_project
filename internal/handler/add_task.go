package handler

import (
	"encoding/json"
	"fmt"
	nextdate "github.com/askoren1/go_final_project/internal/next_date"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

type Task struct {
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func (h *Handler) AddTask(w http.ResponseWriter, r *http.Request) {
	var t Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Проверка правила повторения
	switch t.Repeat {
	case "":
		// Правило "" - корректное, ничего не делаем

	case "y":
		// Правило "y" - корректное, ничего не делаем

	default:
		// Проверяем формат "d <число>"
		match, _ := regexp.MatchString(`^d \d+$`, t.Repeat) // Проверка на соответствие регулярному выражению
		if !match {
			response := map[string]string{"error": "Некорректный формат правила повторения"}
			json.NewEncoder(w).Encode(response)
			return
		}

		daysStr := t.Repeat[2:] // Извлекаем число дней
		days, err := strconv.Atoi(daysStr)
		if err != nil || days > 400 {
			response := map[string]string{"error": "Некорректное значение дней в правиле повторения"}
			json.NewEncoder(w).Encode(response)
			return
		}

	}

	if t.Title == "" { // проверка title на пустоту
		response := map[string]string{"error": "Не указан заголовок задачи"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// DateToday2 - сегодня, string
	// nowTime - сегодня, time.Time
	// t.Date - входное значение даты, string
	// date2 - входное значение даты, time.Time
	// dateInTable - дата для передачи в таблицу, string

	DateToday2 := time.Now().Format("20060102")
	nowTime := time.Now()

	//DateToday2 := now.Format("20060102") // Используем переданное значение now
	//nowTime := now                       // Используем переданное значение now

	var dateInTable string
	layout := "20060102"

	if t.Date == "" || t.Date == "today" { // проверка date на пустоту
		dateInTable = DateToday2
	} else { // проверка даты на корректность
		date2, err := time.Parse(layout, t.Date)
		if err != nil {
			response := map[string]string{"error": "Некорректное значение даты"}
			json.NewEncoder(w).Encode(response)
			return
		}

		if nowTime.Truncate(24 * time.Hour).UTC().After(date2) { //замена даты, если дата задачи меньше сегодняшней

			if t.Repeat == "" {
				dateInTable = DateToday2
			} else {
				dateInTable, _ = nextdate.NextDate(nowTime, t.Date, t.Repeat)
			}

		} else {
			dateInTable = t.Date
		}
	}

	id, err := h.repo.AddTask(dateInTable, t.Title, t.Comment, t.Repeat)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]string{"id": fmt.Sprintf("%d", id)}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(response)

}
