package service_errors

import "errors"

var (
	InvalidFIO         = errors.New("invalid fio")
	InvalidCategory    = errors.New("invalid category")
	InvalidLogin       = errors.New("invalid login")
	InvalidPassword    = errors.New("invalid password")
	InvalidRole        = errors.New("invalid role")
	NotUnique          = errors.New("such row already exists")
	MismatchedPassword = errors.New("passwords do not match")
	InvalidBirthDay    = errors.New("invalid birthdate")
	InvalidGender      = errors.New("invalid gender")
	InvalidRuleNum     = errors.New("invalid rule number")
	InvalidReviewDate  = errors.New("invalid protest review date")
	InvalidStatus      = errors.New("invalid protest status")
	InvalidNumber      = errors.New("invalid race number")
	InvalidDate        = errors.New("invalid date")
	InvalidClass       = errors.New("invalid yacht class")
	InvalidBlowoutCnt  = errors.New("invalid blowout count")
)
