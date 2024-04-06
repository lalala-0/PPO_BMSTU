package models

import (
	"github.com/google/uuid"
)

type Judge struct {
	ID       uuid.UUID
	FIO      string
	Login    string
	Password string
	Role     int
	Post     string
}
