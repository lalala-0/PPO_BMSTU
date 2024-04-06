package models

import (
	"github.com/google/uuid"
)

type Rating struct {
	ID         uuid.UUID
	Class      string
	BlowoutCnt int
}
