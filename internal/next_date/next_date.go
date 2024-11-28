package next_date

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Функция для вычисления следующей даты выполнения задачи, учитывая ее исходную дату и правило повторения
func NextDate(now time.Time, date string, repeat string) (string, error) {
	t, err := time.Parse("20060102", date) //Парсинг исходной даты
	if err != nil {
		return "", err
	}

	//Обработка правила повторения
	switch repeat {
	case "": //Пустое правило
		return "", errors.New("no repeat rule specified")
	case "y": //Ежегодное повторение
		t = t.AddDate(1, 0, 0)
		for t.Before(now) {
			t = t.AddDate(1, 0, 0)
		}
		return t.Format("20060102"), nil
	default: //Повторное через N дней
		if len(repeat) > 1 && repeat[0] == 'd' {
			days, err := strconv.Atoi(repeat[2:])
			if err != nil || days > 400 {
				return "", fmt.Errorf("invalid days: %s", repeat)
			}
			t = t.AddDate(0, 0, days)
			for t.Before(now) {
				t = t.AddDate(0, 0, days)
			}
			return t.Format("20060102"), nil
		}
		return "", fmt.Errorf("invalid repeat rule: %s", repeat)
	}
}
