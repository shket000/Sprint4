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
	// TODO: реализовать функцию
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}
