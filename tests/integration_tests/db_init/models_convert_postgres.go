package db_init

import (
	"PPO_BMSTU/internal/models"
	postgres_rep "PPO_BMSTU/internal/repository/postgres"
)

func CrewFromDb(crewDB *postgres_rep.CrewDB) *models.Crew {
	return &models.Crew{
		ID:       crewDB.ID,
		RatingID: crewDB.RatingID,
		SailNum:  crewDB.SailNum,
		Class:    crewDB.Class,
	}
}

func CrewToDb(crew *models.Crew) *postgres_rep.CrewDB {
	return &postgres_rep.CrewDB{
		ID:       crew.ID,
		RatingID: crew.RatingID,
		SailNum:  crew.SailNum,
		Class:    crew.Class,
	}
}
