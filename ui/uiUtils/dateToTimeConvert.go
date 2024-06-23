package uiUtils

import (
	"log"
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
