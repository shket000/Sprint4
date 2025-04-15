package daysteps

import (
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	temp := strings.Split(data, ",")
	if len(temp) != 2 {
		return 0, time.Duration(0), fmt.Errorf("invalid data")
	}
	steps, err := strconv.Atoi(temp[0])
	if err != nil {
		return 0, time.Duration(0), fmt.Errorf("invalid step")
	}
	if steps < 1 {
		return 0, time.Duration(0), fmt.Errorf("invalid step")
	}
	duration, err := time.ParseDuration(temp[1])
	if err != nil {
		return 0, time.Duration(0), fmt.Errorf("invalid hours")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
}
