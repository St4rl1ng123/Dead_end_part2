package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	//разделение строки
	parts := strings.Split(datastring, ",")

	//проверка длины
	if len(parts) != 3 {
		return errors.New("errors datastring")
	}

	//парсинг шагов
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("error steps")
	}
	t.Steps = steps

	//сохраняем значение в поле Training
	t.TrainingType = parts[1]

	//парсинг строки
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("error duration")
	}
	t.Duration = duration

	//успех
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	//вычисляем дистанцию
	distance := spentenergy.Distance(t.Steps, t.Height)

	//вычисляем скорость
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	//проверка вида тренировки
	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	//вывод строк по условию
	result := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		speed,
		calories,
	)

	return result, nil
}
