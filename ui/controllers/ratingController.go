package controllers

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/ui/modelsUI"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"sort"
)

// GetMenu

func (s *Services) getRatingMenu(c *gin.Context) {
	idStr := c.Param("ratingID")
	ratingID, err := uuid.Parse(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат UUID")
		return
	}

	rating, err := s.Services.RatingService.GetRatingDataByID(ratingID)
	if rating == nil {
		c.String(http.StatusNotFound, "Rating not found")
		return
	}
	authJudge := s.authenticatedJudge(c)

	c.HTML(http.StatusOK, "rating", s.ratingMenu(rating, authJudge))
}

func (s *Services) ratingMenu(rating *models.Rating, authJudge *models.Judge) gin.H {
	races, err := s.Services.RaceService.GetRacesDataByRatingID(rating.ID)
	if err != nil {
		log.Printf("Error getting races: %v", err)
		races = []models.Race{}
	}
	racesData, err := modelsUI.FromRaceModelsToStringData(races)
	if err != nil {
		log.Printf("Error class to string convert: %v", err)
	}

	crews, err := s.Services.CrewService.GetCrewsDataByRatingID(rating.ID)
	if err != nil {
		log.Printf("Error getting crews: %v", err)
		crews = []models.Crew{}
	}
	crewsData, err := modelsUI.FromCrewModelsToStringData(crews)
	if err != nil {
		log.Printf("Error class to string convert: %v", err)
	}

	ratingView, _ := modelsUI.FromRatingModelToStringData(rating)
	var result = gin.H{
		"judge":      authJudge,
		"title":      "",
		"rating":     rating,
		"ratingView": ratingView,
		"races":      racesData,
		"crews":      crewsData,
	}

	return result
}

// GetRatingTable
func (s *Services) getRatingTable(c *gin.Context) {
	idStr := c.Param("ratingID")
	ratingID, err := uuid.Parse(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат UUID")
		return
	}

	rating, err := s.Services.RatingService.GetRatingDataByID(ratingID)
	if err != nil {
		c.String(http.StatusNotFound, "Rating not found")
		return
	}

	ratingTableLines, err := s.Services.RatingService.GetRatingTable(rating.ID)
	if err != nil {
		log.Printf("Error getting ratingTableLines: %v", err)
		c.String(http.StatusInternalServerError, "Internal server error")
		return
	}

	var raceNums []int
	for raceNum, _ := range ratingTableLines[1].ResInRace {
		raceNums = append(raceNums, raceNum)
	}
	sort.Ints(raceNums)

	// Подготовка данных для шаблона
	var lines []gin.H
	for _, line := range ratingTableLines {
		var resInRace []int

		for _, raceNum := range raceNums {
			resInRace = append(resInRace, line.ResInRace[raceNum])
		}

		lineData := gin.H{
			"SailNum":          line.SailNum,
			"ParticipantNames": line.ParticipantNames[0],
			"ResInRace":        resInRace,
			"PointsSum":        line.PointsSum,
			"Rank":             line.Rank,
		}
		lines = append(lines, lineData)
	}

	// Отправка данных в шаблон
	c.HTML(http.StatusOK, "rating-table", gin.H{
		"title":    "Таблица рейтинга",
		"lines":    lines,
		"raceNums": raceNums,
	})
}

func (s *Services) ratingTable(rating *models.Rating, lines []models.RatingTableLine) gin.H {
	var result = gin.H{
		"title": rating.Name, // Здесь можно использовать любое поле рейтинга для заголовка
		"lines": lines,
	}
	return result
}

// UPDATE
func (s *Services) updateRatingGet(c *gin.Context) {
	ratingID, err := uuid.Parse(c.Param("ratingID"))
	if err != nil {
		c.HTML(400, "updateRating", gin.H{
			"title": "Редактировать рейтинг",
			"error": "Неверный идентификатор рейтинга",
		})
		return
	}

	rating, err := s.Services.RatingService.GetRatingDataByID(ratingID)
	if err != nil {
		c.HTML(400, "updateRating", gin.H{
			"title": "Редактировать рейтинг",
			"error": "Рейтинг не найден",
		})
		return
	}

	editedRatingView, _ := modelsUI.FromRatingModelToInputData(rating)

	c.HTML(200, "updateRating", gin.H{
		"title":      "Редактировать рейтинг",
		"ratingView": rating,
		"formData":   editedRatingView,
		"classMap":   modelsUI.ClassMap,
	})
}

func (s *Services) updateRatingPost(c *gin.Context) {
	var input modelsUI.RatingInput
	err := c.Bind(&input)
	if err != nil {
		log.Printf("Error binding data: %v", err)
		c.HTML(400, "updateRating", gin.H{
			"title":    "Редактировать рейтинг",
			"error":    err.Error(),
			"formData": input,
			"classMap": modelsUI.ClassMap,
		})
		return
	}

	ratingID, err := uuid.Parse(c.Param("ratingID"))
	if err != nil {
		c.HTML(400, "updateRating", gin.H{
			"title": "Редактировать профиль",
			"error": "Неверный идентификатор рейтинга",
		})
		return
	}
	rating, err := s.Services.RatingService.GetRatingDataByID(ratingID)
	if err != nil {
		c.HTML(400, "updateRating", gin.H{
			"title":    "Редактировать рейтинг",
			"error":    "Рейтинг не найден",
			"formData": input,
			"classMap": modelsUI.ClassMap,
		})
		return
	}
	_, err = s.Services.RatingService.UpdateRatingByID(rating.ID, input.Name, input.Class, input.BlowoutCnt)
	if err != nil {
		log.Printf("Error updating rating: %v", err)
		c.HTML(400, "updateRating", gin.H{
			"title":    "Редактировать рейтинг",
			"error":    err.Error(),
			"formData": input,
			"classMap": modelsUI.ClassMap,
		})
		return
	}

	c.Redirect(302, "/ratings/"+ratingID.String())
}

func (s *Services) deleteRating(c *gin.Context) {
	ratingID, err := uuid.Parse(c.Param("ratingID"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error-page", gin.H{
			"error": err.Error(),
		})
		return
	}

	err = s.Services.RatingService.DeleteRatingByID(ratingID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error-page", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, "/")
}

// CREATE
func (s *Services) createRatingGet(c *gin.Context) {
	c.HTML(200, "createRating", gin.H{
		"title":    "Создать рейтинг",
		"formData": modelsUI.RatingInput{},
		"classMap": modelsUI.ClassMap,
	})
}

func (s *Services) createRatingPost(c *gin.Context) {
	var input modelsUI.RatingInput
	err := c.Bind(&input)
	if err != nil {
		log.Printf("Error binding data: %v", err)
		c.HTML(400, "createRating", gin.H{
			"title":    "Создать рейтинг",
			"error":    err.Error(),
			"formData": input,
			"classMap": modelsUI.ClassMap,
		})
		return
	}

	_, err = s.Services.RatingService.AddNewRating(uuid.New(), input.Name, input.Class, input.BlowoutCnt)
	if err != nil {
		log.Printf("Error updating rating: %v", err)
		c.HTML(400, "createRating", gin.H{
			"title":    "Создать рейтинг",
			"error":    err.Error(),
			"formData": input,
			"classMap": modelsUI.ClassMap,
		})
		return
	}

	c.Redirect(302, "/")
}
