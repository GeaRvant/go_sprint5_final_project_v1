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
		return fmt.Errorf("недопустимые пробелы в начале или в конце строки")
	}
	strParts := strings.Split(datastring, ",")
	if len(strParts) != 2 {
		return fmt.Errorf("некорректное преобразования строки: ожидалось 2 части, получено %d", len(strParts))
	}
	// Проверка на пустые поля
	for i, part := range strParts {
		if strings.TrimSpace(part) == "" {
			return fmt.Errorf("поле %d содержит пустое значение", i+1)
		}
		// Проверка, что в части нет лишних пробелов внутри
		if part != strings.TrimSpace(part) {
			return fmt.Errorf("в поле %d обнаружены лишние пробелы", i+1)
		}
	}
	elementToInt, err := strconv.Atoi(strings.TrimSpace(strParts[0]))
	if err != nil {
		return fmt.Errorf("неверное преобразование в int: %w", err)
	}
	if elementToInt <= 0 {
		return fmt.Errorf("неверное количество шагов")
	}
	ds.Steps = elementToInt
	duration, err := time.ParseDuration(strings.TrimSpace(strParts[1]))
	if err != nil {
		return fmt.Errorf("некорректный парсинг длительности: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("длительность должна быть больше нуля")
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	dist := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("не удалось получить данные: %w", err)
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, calories), nil
}
