package spentenergy

import (
	"errors"
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
	// TODO: реализовать функцию
	//проверка входных данных
	if steps <= 0 {
		return 0, errors.New("errors steps")
	}
	if weight <= 0 {
		return 0, errors.New("errors weight")
	}
	if height <= 0 {
		return 0, errors.New("errors height")
	}
	if duration <= 0 {
		return 0, errors.New("errors duration")
	}

	//средняя скорость
	speed := MeanSpeed(steps, height, duration)

	//рассчёт калорий
	minutes := duration.Minutes()

	calories := (weight * speed * minutes) / minInH
	calories *= walkingCaloriesCoefficient
	return calories, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	//проверка входных данных
	if steps <= 0 {
		return 0, errors.New("errors steps")
	}
	if weight <= 0 {
		return 0, errors.New("errors weight")
	}
	if height <= 0 {
		return 0, errors.New("errors height")
	}
	if duration <= 0 {
		return 0, errors.New("errors duration")
	}

	//средняя скорость
	meanSpeed := MeanSpeed(steps, height, duration)

	//рассчёт и возврат каллорий
	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	//проверка duration
	if steps < 0 {
		return 0
	}

	if height < 0 {
		return 0
	}

	if duration <= 0 {
		return 0
	}

	//вычисление дистанции
	dist := Distance(steps, height)
	speed := dist / duration.Hours()

	return speed
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	//длина шага
	stepLength := height * stepLengthCoefficient
	//пройденное кол-во шагов
	distanceMeters := float64(steps) * stepLength

	return distanceMeters / mInKm
}
