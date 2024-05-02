package models

import (
	"github.com/google/uuid"
)

type CrewResInRace struct {
	CrewID           uuid.UUID
	RaceID           uuid.UUID
	Points           int
	SpecCircumstance int
}

//
//// SpecCircumstance vars
//const dns = 1  // не стартовала (не подпадает под DNC и OCS),
//const dnf = 2  // не финишировала,
//const dnc = 3  // не стартовала; не прибыла в район старта,
//const ocs = 4  //  не стартовала; находилась на стороне дистанции от стартовой линии в момент сигнала "Старт" для нее и не стартовала или нарушила правило 30.1,
//const zfp = 5  // 20% наказание по правилу 30.2,
//const ufd = 6  // дисквалификация по правилу 30.3,
//const bfd = 7  // дисквалификация по правилу 30.4,
//const scp = 8  // применено "Наказание штрафными очками",
//const ret = 9  // вышла из гонки,
//const dsq = 10 // дисквалификация,
//const dne = 11 // не исключаемая дисквалификация,
//const RDG = 12 // исправлен результат,
//const DPI = 13 // наказание по усмотрению протестового комитета.
