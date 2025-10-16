package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	if datastring != strings.TrimSpace(datastring) {
		return fmt.Errorf("недопустимые пробелы в начале или в конце строки")
	}
	strParts := strings.Split(datastring, ",")
	if len(strParts) != 3 {
		return fmt.Errorf("некорректное преобразование строки: ожидалось 3 части, получено %d", len(strParts))
	}
	// Проверка на пустые поля
	for i, part := range strParts {
		if strings.TrimSpace(part) == "" {
			return fmt.Errorf("поле %d содержит пустое значение", i+1)
		}
	}
	elementToInt, err := strconv.Atoi(strings.TrimSpace(strParts[0]))
	if err != nil {
		return fmt.Errorf("преобразования в int: %w", err)
	}
	if elementToInt <= 0 {
		return fmt.Errorf("неверное количество шагов")
	}
	t.Steps = elementToInt
	t.TrainingType = strings.TrimSpace(strParts[1])
	duration, err := time.ParseDuration(strings.TrimSpace(strParts[2]))
	if err != nil {
		return fmt.Errorf("некорректный парсинг длительности: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("длительность должна быть больше нуля")
	}
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	dist := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	durationHour := t.Duration.Hours()

	var calories float64
	var err error

	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}
	if err != nil {
		return "", fmt.Errorf("не удалось получить данные: %w", err)
	}
	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, durationHour, dist, speed, calories)

	return result, nil
}
