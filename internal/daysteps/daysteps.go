package daysteps

import (
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
	"log"
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
	if duration <= 0 {
		return 0, time.Duration(0), fmt.Errorf("invalid duration")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	if steps < 1 {
		return ""
	}
	distance := float64(steps) * stepLength
	distanceKm := distance / float64(mInKm)
	ccal, err2 := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err2 != nil {
		log.Println(err2)
		return ""
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKm, ccal)
	return result
}
