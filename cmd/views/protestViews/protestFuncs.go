package protestViews

import (
	utils "PPO_BMSTU/cmd/cmdUtils"
	"PPO_BMSTU/cmd/modelTables"
	"PPO_BMSTU/cmd/views"
	"PPO_BMSTU/cmd/views/crewViews"
	"PPO_BMSTU/cmd/views/stringConst"
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
	"fmt"
)

func DeleteProtest(service registry.Services, protest *models.Protest) error {
	err := service.ProtestService.DeleteProtestByID(protest.ID)

	if err != nil {
		return err
	}

	fmt.Printf("Протест успешно удалён\n\n\n")

	return nil
}

func UpdateProtest(service registry.Services, protest *models.Protest) error {
	race := models.Race{}
	err := views.GetRaceInRating(service, protest.RatingID, &race)
	judge := models.Judge{}
	err = views.GetJudge(service, &judge)
	ruleNum := utils.EndlessReadInt(stringConst.RuleNumRequest)
	reviewDate := utils.EndlessReadDateTime(stringConst.DateRequest)
	status := utils.EndlessReadInt(stringConst.ProtestStatusRequest)
	comment := utils.EndlessReadRow(stringConst.CommentRequest)

	updatedProtest, err := service.ProtestService.UpdateProtestByID(protest.ID, race.ID, judge.ID, ruleNum, reviewDate, status, comment)

	if err != nil {
		return err
	}

	fmt.Printf("%s %s %s %s Протест успешно обновлён\n\n\n", updatedProtest.RuleNum, updatedProtest.ReviewDate, updatedProtest.Status, updatedProtest.Comment)

	return nil
}

func CreateProtest(service registry.Services, race *models.Race) error {
	judge := models.Judge{}
	err := views.GetJudge(service, &judge)
	ruleNum := utils.EndlessReadInt(stringConst.RuleNumRequest)
	reviewDate := utils.EndlessReadDateTime(stringConst.DateRequest)
	comment := utils.EndlessReadRow(stringConst.CommentRequest)

	fmt.Printf("Выберите протестуещего\n")
	crew := models.Crew{}
	err = crewViews.GetCrewInRating(service, &crew, race.RatingID)
	if err != nil {
		return err
	}
	protesteeSailNum := crew.SailNum

	fmt.Printf("Выберите опротестованного\n")
	err = crewViews.GetCrewInRating(service, &crew, race.RatingID)
	if err != nil {
		return err
	}
	protestorSailNum := crew.SailNum

	witnessesSailNumsMap := utils.EndlessReadIntSerialMap(stringConst.WitnessesSailNumsRequest)
	witnessesSailNums := make([]int, len(witnessesSailNumsMap))
	i := 0
	for _, k := range witnessesSailNumsMap {
		witnessesSailNums[i] = k
		i++
	}

	createdProtest, err := service.ProtestService.AddNewProtest(race.ID, race.RatingID, judge.ID, ruleNum, reviewDate, comment, protesteeSailNum, protestorSailNum, witnessesSailNums)

	if err != nil {
		return err
	}

	fmt.Printf("%s %s %s %s Протест успешно создан\n\n\n", createdProtest.RuleNum, createdProtest.ReviewDate, createdProtest.Status, createdProtest.Comment)

	return nil
}

func CompleteProtestReview(service registry.Services, protest *models.Protest) error {
	protesteePoints := utils.EndlessReadInt(stringConst.ProtesteePointsRequest)
	comment := utils.EndlessReadRow(stringConst.CommentRequest)

	err := service.ProtestService.CompleteReview(protest.ID, protesteePoints, comment)

	if err != nil {
		return err
	}

	fmt.Printf("Рассмотрение протеста успешно завершено\n\n\n")

	return nil
}

func GetProtestJudgeMenu(service registry.Services, judge *models.Judge, race *models.Race) error {
	protest := models.Protest{}
	err := views.GetProtestInRace(service, &protest, race)

	err = protestJudgeMenu(service, &protest)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func GetAllProtestsInRace(services registry.Services, race *models.Race) error {
	protests, err := services.ProtestService.GetProtestsDataByRaceID(race.ID)
	if err != nil {
		return err
	}
	return modelTables.Protests(protests)
}

func GetProtestViewerMenu(service registry.Services, race *models.Race) error {
	protest := models.Protest{}
	err := views.GetProtestInRace(service, &protest, race)

	err = protestViewerMenu(service, &protest)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func GetProtestParticipants(service registry.Services, protest *models.Protest) error {
	crewIDs, err := service.ProtestService.GetProtestParticipantsIDByID(protest.ID)
	if err != nil {
		return err
	}
	fmt.Printf("\n-----------\n")

	for crewID, role := range crewIDs {
		crew, err := service.CrewService.GetCrewDataByID(crewID)
		if err != nil {
			return err
		}
		roleStr, err := modelTables.ProtestParticipantRoleToString(role)
		if err != nil {
			return err
		}
		fmt.Printf("\n\t %d - %s", crew.SailNum, roleStr)
	}

	fmt.Printf("\n-----------\n")

	return err
}
