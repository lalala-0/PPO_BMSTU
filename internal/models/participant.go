package models

import (
	"github.com/google/uuid"
	"time"
)

type Participant struct {
	ID       uuid.UUID
	FIO      string
	Category int
	Gender   int
	Birthday time.Time
	Coach    string
}

//// Gender vars
//const Male = true
//const Female = false
//
//// Category vars
//const MasterInternational = 1 //"Master of Sports of Russia of international class"
//const MasterRussia = 2        // "Master of Sports of Russia"
//const Candidate = 3           // "Candidate for Master of Sports"
//const sport1category = 4      // "1 sports category"
//const sport2category = 5      // "2 sports category"
//const sport3category = 6      // "3 sports category"
//const junior1category = 7     // "1 junior category"
//const junior2category = 8     // "2 junior category"
