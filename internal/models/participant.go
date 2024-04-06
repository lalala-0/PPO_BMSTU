package models

import (
	"github.com/google/uuid"
	"time"
)

type Participant struct {
	ID       uuid.UUID
	FIO      string
	Category string
	Gender   string
	Birthday time.Time
	Coach    string
}
