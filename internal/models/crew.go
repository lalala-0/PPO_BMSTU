package models

import (
	"github.com/google/uuid"
)

type Crew struct {
	ID       uuid.UUID
	RatingID uuid.UUID
	SailNum  int
	Class    int
}

//
//// Class yacht vars
//const Laser = 1
//const LaserRadial = 2
//const Optimist = 3
//const Zoom8 = 4
//const Finn = 5
//const SB20 = 6
//const J70 = 6
//const Nacra17 = 6
//const C49er = 6
//const RS_X = 6
//const Cadet = 6
