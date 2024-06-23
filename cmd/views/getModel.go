package views

import (
	utils "PPO_BMSTU/cmd/cmdUtils"
	"PPO_BMSTU/cmd/modelTables"
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
	"fmt"
	"github.com/google/uuid"
)

func GetJudge(service registry.Services, judge *models.Judge) error {
	judges, err := service.JudgeService.GetAllJudges()

	if err != nil {
		return err
	}

	err = modelTables.Judges(judges)
	if err != nil {
		return err
	}

	fmt.Printf("\n-----------\n" +
		"Введите номер соответствующего судьи\n" +
		"Введите 0, чтобы выйти\n\n")

	for {
		var judgeNumber int
		_, err = fmt.Scanf("%d", &judgeNumber)
		if err != nil {
			return err
		}

		if judgeNumber == 0 {
			return nil
		}

		if !utils.ValidateNumber(judgeNumber, len(judges)) {
			fmt.Println("Неверный номер записи")
			continue
		}
		*judge = judges[judgeNumber-1]
		return nil
	}
}

func GetParticipant(service registry.Services, participant *models.Participant) error {

	participants, err := service.ParticipantService.GetAllParticipants()

	if err != nil {
		return err
	}

	err = modelTables.Participants(participants)
	if err != nil {
		return err
	}

	fmt.Printf("\n-----------\n" +
		"Введите номер участника\n" +
		"Введите 0, чтобы выйти\n\n")

	for {
		var participantNumber int
		_, err = fmt.Scanf("%d", &participantNumber)
		if err != nil {
			return err
		}

		if !utils.ValidateNumber(participantNumber, len(participants)) {
			fmt.Println("Неверный номер участника")
			continue
		}
		*participant = participants[participantNumber-1]
		return nil
	}
}

func GetProtestInRace(service registry.Services, protest *models.Protest, race *models.Race) error {
	protests, err := service.ProtestService.GetProtestsDataByRaceID(race.ID)

	if err != nil {
		return err
	}

	err = modelTables.Protests(protests)
	if err != nil {
		return err
	}

	fmt.Printf("\n-----------\n" +
		"Введите номер протеста\n" +
		"Введите 0, чтобы выйти\n\n")

	for {
		var protestNumber int
		_, err = fmt.Scanf("%d", &protestNumber)
		if err != nil {
			return err
		}

		if !utils.ValidateNumber(protestNumber, len(protests)) {
			fmt.Println("Неверный номер протеста")
			continue
		}
		*protest = protests[protestNumber-1]
		return nil
	}
}

func GetRaceInRating(service registry.Services, ratingID uuid.UUID, race *models.Race) error {
	races, err := service.RaceService.GetRacesDataByRatingID(ratingID)

	if err != nil {
		return err
	}

	err = modelTables.Races(races)
	if err != nil {
		return err
	}

	fmt.Printf("\n-----------\n" +
		"Введите номер гонки\n" +
		"Введите 0, чтобы выйти\n\n")

	for {
		var raceNumber int
		_, err = fmt.Scanf("%d", &raceNumber)
		if err != nil {
			return err
		}

		if raceNumber == 0 {
			return nil
		}

		if !utils.ValidateNumber(raceNumber, len(races)) {
			fmt.Println("Неверный номер гонки")
			continue
		}
		*race = races[raceNumber-1]
		return nil
	}
}

func GetRating(service registry.Services, rating *models.Rating) error {

	ratings, err := service.RatingService.GetAllRatings()

	if err != nil {
		return err
	}

	err = modelTables.Ratings(ratings)
	if err != nil {
		return err
	}

	fmt.Printf("\n-----------\n" +
		"Введите номер рейтинга\n" +
		"Введите 0, чтобы выйти\n\n")

	for {
		var ratingNumber int
		_, err = fmt.Scanf("%d", &ratingNumber)
		if err != nil {
			return err
		}

		if !utils.ValidateNumber(ratingNumber, len(ratings)) {
			fmt.Println("Неверный номер рейтинга")
			continue
		}
		*rating = ratings[ratingNumber-1]
		return nil
	}
}
