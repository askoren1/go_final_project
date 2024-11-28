package handler

import (
	"encoding/json"
	"github.com/askoren1/go_final_project/internal/models"
	nextdate "github.com/askoren1/go_final_project/internal/next_date"
	"net/http"
	"regexp"
	"strconv"
	"time"
	)

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var t models.Task2

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if t.Title == "" {
		response := map[string]string{"error": "Не указан заголовок задачи"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Проверка правила повторения
	switch t.Repeat {
	case "", "y":
		// Правила "", "y" - корректные, ничего не делаем
	default:
		match, _ := regexp.MatchString(`^d \d+$`, t.Repeat)
		if !match {
			response := map[string]string{"error": "Некорректный формат правила повторения"}
			json.NewEncoder(w).Encode(response)
			return
		}

		daysStr := t.Repeat[2:]
		days, err := strconv.Atoi(daysStr)
		if err != nil || days > 400 {
			response := map[string]string{"error": "Некорректное значение дней в правиле повторения"}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	DateToday2 := time.Now().Format("20060102")
	nowTime := time.Now()
	var dateInTable string
	layout := "20060102"

	if t.Date == "" || t.Date == "today" {
		dateInTable = DateToday2
	} else {
		date2, err := time.Parse(layout, t.Date)
		if err != nil {
			response := map[string]string{"error": "Некорректное значение даты"}
			json.NewEncoder(w).Encode(response)
			return
		}

		if nowTime.Truncate(24 * time.Hour).UTC().After(date2) {
			if t.Repeat == "" {
				dateInTable = DateToday2
			} else {
				dateInTable, err = nextdate.NextDate(nowTime, t.Date, t.Repeat)
				if err != nil {
					response := map[string]string{"error": "Ошибка вычисления следующей даты: " + err.Error()}
					json.NewEncoder(w).Encode(response)
					return
				}
			}

		} else {
			dateInTable = t.Date
		}
	}

	err := h.repo.UpdateTask(t.ID, dateInTable, t.Title, t.Comment, t.Repeat)
	if err != nil {
		// Возвращаем JSON с ошибкой
		w.WriteHeader(http.StatusNotFound) // Устанавливаем код ответа 404
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Возвращаем JSON с сообщением об успехе
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{})
}
