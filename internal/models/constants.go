package models

const (
	// Class yacht vars
	Laser       = 1
	LaserRadial = 2
	Optimist    = 3
	Zoom8       = 4
	Finn        = 5
	SB20        = 6
	J70         = 7
	Nacra17     = 8
	C49er       = 9
	RS_X        = 10
	Cadet       = 11

	// SpecCircumstance vars
	DNS = 1  // не стартовала (не подпадает под DNC и OCS),
	DNF = 2  // не финишировала,
	DNC = 3  // не стартовала; не прибыла в район старта,
	OCS = 4  //  не стартовала; находилась на стороне дистанции от стартовой линии в момент сигнала "Старт" для нее и не стартовала или нарушила правило 30.1,
	ZFP = 5  // 20% наказание по правилу 30.2,
	UFD = 6  // дисквалификация по правилу 30.3,
	BFD = 7  // дисквалификация по правилу 30.4,
	SCP = 8  // применено "Наказание штрафными очками",
	RET = 9  // вышла из гонки,
	DSQ = 10 // дисквалификация,
	DNE = 11 // не исключаемая дисквалификация,
	RDG = 12 // исправлен результат,
	DPI = 13 // наказание по усмотрению протестового комитета.

	// Role vars
	MainJudge    = 1
	NotMainJudge = 2

	// Gender vars
	Male   = 0
	Female = 1

	// Category vars
	MasterInternational = 1 //"Master of Sports of Russia of international class"
	MasterRussia        = 2 // "Master of Sports of Russia"
	Candidate           = 3 // "Candidate for Master of Sports"
	Sport1category      = 4 // "1 sports category"
	Sport2category      = 5 // "2 sports category"
	Sport3category      = 6 // "3 sports category"
	Junior1category     = 7 // "1 junior category"
	Junior2category     = 8 // "2 junior category"

	// Protest status vars
	PendingReview = 1
	Reviewed      = 2

	// Protest participants role vars
	Protestor = 1
	Protestee = 2
	Witness   = 3
)
