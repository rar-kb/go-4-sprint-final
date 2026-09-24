package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parse := strings.Split(data, ",")
	if len(parse) != 3 {
		return 0, "", 0, fmt.Errorf("Ожидалось 3 элемента, получено %d", len(parse))
	}
	steps, err := strconv.Atoi(parse[0])
	if err != nil {
		return 0, "", 0, err
	}
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("количество шагов должно быть больше нуля")
	}
	time, err := time.ParseDuration(parse[2])
	if err != nil {
		return 0, "", 0, err
	}
	if time <= 0 { // ← И ЭТОГО
		return 0, "", 0, fmt.Errorf("продолжительность должна быть больше нуля")
	}
	// вид активности
	activ := parse[1]
	return steps, activ, time, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	//Рассчитываем длину шага умножая рост пользователя на коэффициент длины шага
	distanceStep := height * stepLengthCoefficient
	// умножаем пройденное кол-во шагов н адлину шага и делим полученное значение на число метров в километре
	distanceKm := (float64(steps) * distanceStep) / mInKm
	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	// находим дистанцию
	dist := distance(steps, height)
	// вычисляем среднюю скорость : делим дистанцию на время в часах
	speed := dist / duration.Hours()
	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activ, time, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	// объявляем переменные дистанции, средней скорости и калорий
	var dist, speed, calories float64

	switch activ {
	case "Бег":
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, time)
		calories, err = RunningSpentCalories(steps, weight, height, time)
	case "Ходьба":
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, time)
		calories, err = WalkingSpentCalories(steps, weight, height, time)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activ, time.Hours(), dist, speed, calories,
	), nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// проверка что шагов не отрицательное кол-во
	if steps <= 0 {
		return 0, fmt.Errorf("Не может быть отрицательное число шагов")
	}
	// такие же проверки проводим с весом и ростом и временем
	if weight <= 0 {
		return 0, fmt.Errorf("Не может быть отрицательное число веса")
	}
	if height <= 0 {
		return 0, fmt.Errorf("Не может быть отрицательное число роста")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("Не может быть отрицательное число времени")
	}
	// рассчитываем среднюю скорость
	meanSpeed := meanSpeed(steps, height, duration)
	// переводим продолжительность в минуты
	durationInMinutes := duration.Minutes()
	result := (weight * meanSpeed * durationInMinutes) / minInH
	return result, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("Не может быть отрицательное число шагов")
	}
	// такие же проверки проводим с весом и ростом и временем
	if weight <= 0 {
		return 0, fmt.Errorf("Не может быть отрицательное число веса")
	}
	if height <= 0 {
		return 0, fmt.Errorf("Не может быть отрицательное число роста")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("Не может быть отрицательное число времени")
	}
	// рассчитываем среднюю скорость
	meanSpeed := meanSpeed(steps, height, duration)
	// переводим продолжительность в минуты
	durationInMinutes := duration.Minutes()
	// записываем результат
	result := ((weight * meanSpeed * durationInMinutes) / minInH) * walkingCaloriesCoefficient
	return result, nil
}
