package ratingViews

import (
	utils "PPO_BMSTU/cmd/cmdUtils"
	"PPO_BMSTU/cmd/modelTables"
	"PPO_BMSTU/cmd/views"
	"PPO_BMSTU/cmd/views/stringConst"
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
	"fmt"
	"github.com/google/uuid"
)

func DeleteRating(service registry.Services, rating *models.Rating) error {
	err := service.RatingService.DeleteRatingByID(rating.ID)

	if err != nil {
		return err
	}

	fmt.Printf("%s %d %d Рейтинг успешно удалён\n\n\n", rating.Name, rating.Class, rating.BlowoutCnt)

	return nil
}

func UpdateRating(service registry.Services, rating *models.Rating) error {
	name := utils.EndlessReadWord(stringConst.RatingNameRequest)
	class := utils.EndlessReadInt(stringConst.ClassRequest)
	blowCnt := utils.EndlessReadInt(stringConst.BlowCntRequest)

	updatedRating, err := service.RatingService.UpdateRatingByID(rating.ID, name, class, blowCnt)

	if err != nil {
		return err
	}

	fmt.Printf("%s %d %d Рейтинг успешно обновлён\n\n\n", updatedRating.Name, updatedRating.Class, updatedRating.BlowoutCnt)

	return nil
}

func GetAllRatings(services registry.Services) error {
	ratings, err := services.RatingService.GetAllRatings()
	if err != nil {
		return err
	}
	return modelTables.Ratings(ratings)
}

func CreateRating(service registry.Services) error {
	name := utils.EndlessReadWord(stringConst.RatingNameRequest)
	class := utils.EndlessReadInt(stringConst.ClassRequest)
	blowCnt := utils.EndlessReadInt(stringConst.BlowCntRequest)

	rating, err := service.RatingService.AddNewRating(uuid.New(), name, class, blowCnt)

	if err != nil {
		return err
	}

	fmt.Printf("%s %d %d Рейтинг успешно создан\n\n\n", rating.Name, rating.Class, rating.BlowoutCnt)

	return nil
}

func GetRatingJudgeMenu(service registry.Services, judge *models.Judge) error {
	rating := models.Rating{}
	err := views.GetRating(service, &rating)
	if err != nil {
		return err
	}
	err = ratingJudgeMenu(service, &rating, judge)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func GetRatingViewerMenu(service registry.Services) error {
	rating := models.Rating{}
	err := views.GetRating(service, &rating)
	if err != nil {
		return err
	}
	err = ratingViewerMenu(service, &rating)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func GetRatingTable(service registry.Services, rating *models.Rating) error {
	ratingTableLines, err := service.RatingService.GetRatingTable(rating.ID)
	if err != nil {
		return err
	}
	return modelTables.RatingTableLines(ratingTableLines)

}
