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
	split := strings.Split(data, ",")
	if len(split) != 3 {
		return 0, "", 0, fmt.Errorf("Ошибка")
	}
	steps, err := strconv.Atoi(split[0])
	if err != nil {
		return 0, "", 0, err
	}
	duration, err := time.ParseDuration(split[2])
	if err != nil {
		return 0, "", 0, err
	}
	return steps, split[1], duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	x := float64(steps) * stepLength
	distance := x / mInKm
	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)
	mSpeed := dist / duration.Hours()
	return mSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activ, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	switch activ {
	case "Бег":
		calories, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`Тип тренировки: Бег
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f`, duration.Hours(), distance(steps, height), meanSpeed(steps, height, duration), calories), nil
	case "Ходьба":
		calories, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`Тип тренировки: Ходьба
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f`, duration.Hours(), distance(steps, height), meanSpeed(steps, height, duration), calories), nil
	default:
		fmt.Println("неизвестный тип тренировки")

		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Ошибка")
	}
	avgSpeed := meanSpeed(steps, height, duration)
	minDuration := duration.Minutes()
	x := (weight * avgSpeed * minDuration) / minInH
	return x, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Ошибка")
	}
	avgSpeed := meanSpeed(steps, height, duration)
	minDuration := duration.Minutes()
	x := (weight * avgSpeed * minDuration) / minInH
	walkCalories := x * walkingCaloriesCoefficient
	return walkCalories, nil
}
