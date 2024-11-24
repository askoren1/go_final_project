package next_date

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func NextDate(now time.Time, date string, repeat string) (string, error) {
	// Преобразуем строку с датой в объект времени
	t, err := time.Parse("20060102", date)
	if err != nil {
		return "", err
	}

	// Определяем правило повторения
	switch repeat {
	case "":
		// Если правило не указано, возвращаем пустую строку (удаление задачи)
		return "", errors.New("no repeat rule specified")
	case "y":
		// Ежегодное повторение
		next := t
		for {
			next = t.AddDate(1, 0, 0)
			if next.After(now) {
				break
			}
			t = next
		}
		return next.Format("20060102"), nil
	default:
		// Проверяем, начинается ли правило с "d"
		if len(repeat) > 1 && repeat[0] == 'd' {
			//Повторение через указанное количеством дней
			days, err := strconv.Atoi(repeat[2:])
			if err != nil || days > 400 {
				return "", fmt.Errorf("invalid days: %s", repeat)
			}

			//for next := t.AddDate(0, 0, days); next < now;

			next := t
			for {
				next = t.AddDate(0, 0, days)
				if next.After(now) {
					break
				}
				t = next
			}
			return next.Format("20060102"), nil
		}
		return "", fmt.Errorf("invalid repeat rule: %s", repeat)
	}
}
