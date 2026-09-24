package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/rar-kb/go-4-sprint-final/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	parse := strings.Split(data, ",")
	// проверка что получено верное количество данных
	if len(parse) != 2 {
		return 0, 0, fmt.Errorf("Ожидалось 2 элемента, получено %d", len(parse))
	}
	// кол-во шагов
	steps, err := strconv.Atoi(parse[0])
	if err != nil {
		return 0, 0, err
	}
	// проверяем что шагов больше 0
	if steps < 0 {
		return 0, 0, fmt.Errorf("Количество шагов не может быть отрицательным")
	}
	// время
	time, err := time.ParseDuration(parse[1])
	if err != nil {
		return 0, 0, err
	}
	return steps, time, nil

}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, time, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if steps < 0 {
		return ""
	}
	// дистанция в метрах
	distanceMetr := float64(steps) * stepLength
	// дистанция в километрах
	distanceKilometr := distanceMetr / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, time)
	if err != nil {
		return ""
	}
	return fmt.Sprintf(
		"Количество шагов: %d. \nДистанция составила %.2f км. \nВы сожгли %.2f ккал.",
		steps, distanceKilometr, calories,
	)

}
