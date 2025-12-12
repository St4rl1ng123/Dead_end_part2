package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	//разделение строки
	parts := strings.Split(datastring, ",")

	//проверка длины
	if len(parts) != 2 {
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
	ds.Steps = steps

	//парсинг длительности
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("error duration")
	}
	ds.Duration = duration

	//успех
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	//вычисляем дистанцию
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	//вычисляем калории
	calories, err := spentenergy.WalkingSpentCalories(
		ds.Steps,
		ds.Weight,
		ds.Height,
		ds.Duration,
	)
	if err != nil {
		return "", err
	}

	//вывод строки по условию
	result := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	)

	//Успех!
	return result, nil
}
