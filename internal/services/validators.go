package services

import (
	"PPO_BMSTU/internal/models"
	"regexp"
	"time"
)

func validFIO(fio string) bool {
	return len(fio) > 0
}

func validLogin(login string) bool {
	return len(login) > 0
}

func validRole(role int) bool {
	return role == 1 || role == 2
}
func validProtestRole(role int) bool {
	return role == 1 || role == 2 || role == 3
}
func validFlag(role int) bool {
	return role == 1 || role == 0
}
func validPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	reLetter := regexp.MustCompile(`[a-zA-Z]`)
	reNumber := regexp.MustCompile(`[0-9]`)

	return reLetter.MatchString(password) && reNumber.MatchString(password)
}

func validNumber(num int) bool {
	return num > 0
}

func validSailNum(num int) bool {
	return num > 0
}

func validBlowoutCnt(num int) bool {
	return num >= 0
}
func validClass(class int) bool {
	return class > 0 && class < 12
}

func validRuleNum(num int) bool {
	return num == 31 || num == 42 || num >= 10 && num < 14 || num > 14 && num < 24
}

func validStatus(status int) bool {
	return status == models.PendingReview || status == models.Reviewed
}

func validCategory(category int) bool {
	return category >= 1 && category <= 8
}

func validSpecCircumstance(category int) bool {
	return category >= 1 && category <= 13
}

func validGender(gender int) bool {
	return gender == 1 || gender == 2
}

func validBirthDay(birthdate time.Time) bool {
	year := birthdate.Year()
	nyear := time.Now().Year()
	return nyear-year > 10 && nyear-year < 150
}
