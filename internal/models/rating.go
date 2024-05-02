package models

import (
	"github.com/google/uuid"
)

type Rating struct {
	ID         uuid.UUID
	Name       string
	Class      int
	BlowoutCnt int
}
