package spentcalories

import (
	"errors"
	"fmt"
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
	temp := strings.Split(data, ",")
	if len(temp) != 3 {
		return 0, "", time.Duration(0), fmt.Errorf("Training data is not valid")
	}
	step, err := strconv.Atoi(temp[0])
	if err != nil {
		return 0, "", time.Duration(0), fmt.Errorf("Step is not valid")
	}
	duration, err := time.ParseDuration(temp[2])
	if err != nil {
		return 0, "", time.Duration(0), fmt.Errorf("Duration is not valid")
	}
	return step, temp[1], duration, nil

}

func distance(steps int, height float64) float64 {
	return stepLengthCoefficient * height * float64(steps) / mInKm

}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration < 0 {
		return 0.0
	}
	dist := distance(steps, height)
	return dist / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, train, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}
	durationInHours := duration.Hours()
	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	calRun, _ := RunningSpentCalories(steps, weight, height, duration)
	calWalk, _ := WalkingSpentCalories(steps, weight, height, duration)
	result := ``
	switch train {
	case "Бег":
		result = fmt.Sprintf("Тип тренировки: Бег\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", durationInHours, dist, speed, calRun)
	case "Ходьба":
		result = fmt.Sprintf("Тип тренировки: Ходьба\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", durationInHours, dist, speed, calWalk)
	default:
		return result, fmt.Errorf("неизвестный тип тренировки")
	}
	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps < 1 {
		return 0.0, errors.New("Steps must be greater than zero")
	}
	if weight < 0 {
		return 0.0, errors.New("Weight must be greater than zero")
	}
	if height < 0 {
		return 0.0, errors.New("Height must be greater than zero")
	}
	if duration < 0 {
		return 0.0, errors.New("Duration must be greater than zero")
	}
	averageSpeed := meanSpeed(steps, height, duration)

	return weight * averageSpeed * duration.Minutes() / minInH, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps < 1 {
		return 0.0, errors.New("Steps must be greater than zero")
	}
	if weight < 0 {
		return 0.0, errors.New("Weight must be greater than zero")
	}
	if height < 0 {
		return 0.0, errors.New("Height must be greater than zero")
	}
	if duration < 0 {
		return 0.0, errors.New("Duration must be greater than zero")
	}
	averageSpeed := meanSpeed(steps, height, duration)

	return weight * averageSpeed * duration.Minutes() / minInH * walkingCaloriesCoefficient, nil
}
