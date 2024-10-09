package uiUtils

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func ConvertStringToTime(dateStr string) time.Time {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func FormatDate(t time.Time) string {
	return t.Format("02-01-2006") // DD-MM-YYYY
}

func ParseHtmlToggle(rawBool string) bool {
	// convert "on"/"off" to true/false
	return rawBool == "on"
}

func ParseString(input string) ([]int, error) {
	var witnesses []int
	entries := strings.Split(input, ",")
	for _, entry := range entries {
		num, err := strconv.Atoi(strings.TrimSpace(entry))
		if err != nil {
			return nil, err
		}
		witnesses = append(witnesses, num)
	}
	return witnesses, nil
}

func SliceToMapSerial(arr []int) map[int]int {
	result := make(map[int]int)

	for i, val := range arr {
		result[val] = i + 1
	}

	return result
}

func SliceToMapConst(arr []int) map[int]int {
	result := make(map[int]int)

	for _, val := range arr {
		result[val] = 3
	}

	return result
}

// Пример функции обработки даты и времени
func ParseDateTime(dateTimeStr string) (time.Time, error) {
	layout := "2006-01-02T15:04"
	return time.Parse(layout, dateTimeStr)
}
