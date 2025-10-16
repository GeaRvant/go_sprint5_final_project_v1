package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("некорректное количество шагов (%d)", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("некорректное значение веса (%.2f)", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("некорректное значение роста (%.2f)", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("некорректное значение длительности (%v)", duration)
	}
	speed := MeanSpeed(steps, height, duration)
	durationMin := duration.Minutes()
	return ((weight * speed * durationMin) / minInH) * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("некорректное количество шагов (%d)", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("некорректное значение веса (%.2f)", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("некорректное значение роста (%.2f)", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("некорректное значение длительности (%v)", duration)
	}
	speed := MeanSpeed(steps, height, duration)
	durationMin := duration.Minutes()
	return (weight * speed * durationMin) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || height <= 0 { //добавлена дополнительная проверка на корректный ввод роста
		return 0
	}
	if duration <= 0 {
		return 0
	}
	dist := Distance(steps, height)
	durationHour := duration.Hours()
	return dist / durationHour
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return (float64(steps) * stepLength) / mInKm
}
