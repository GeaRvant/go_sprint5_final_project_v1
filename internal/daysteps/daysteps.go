package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	if datastring != strings.TrimSpace(datastring) {
		return fmt.Errorf("ошибка: недопустимые пробелы в начале или в конце строки")
	}
	strParts := strings.Split(datastring, ",")
	if len(strParts) != 2 {
		return fmt.Errorf("ошибка преобразования строки: ожидалось 2 части, получено %d", len(strParts))
	}
	// Проверка на пустые поля
	for i, part := range strParts {
		if strings.TrimSpace(part) == "" {
			return fmt.Errorf("ошибка: поле %d содержит пустое значение", i+1)
		}
	}
	elementToInt, err := strconv.Atoi(strings.TrimSpace(strParts[0]))
	if err != nil {
		return fmt.Errorf("ошибка преобразования в int: %w", err)
	}
	if elementToInt <= 0 {
		return fmt.Errorf("ошибка: неверное количество шагов")
	}
	ds.Steps = elementToInt
	duration, err := time.ParseDuration(strings.TrimSpace(strParts[1]))
	if err != nil {
		return fmt.Errorf("ошибка парсинга длительности: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("ошибка: длительность должна быть больше нуля")
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	dist := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка при получении данных: %w", err)
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, calories), nil
}
