package controllers

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/ui/modelsUI"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (s *Services) getRaceMenu(c *gin.Context) {
	ridStr := c.Param("ratingID")
	ratingID, err := uuid.Parse(ridStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат UUID")
		return
	}
	rating, err := s.Services.RatingService.GetRatingDataByID(ratingID)
	if rating == nil {
		c.String(http.StatusNotFound, "Rating not found")
		return
	}

	idStr := c.Param("raceID")
	raceID, err := uuid.Parse(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат UUID")
		return
	}
	race, err := s.Services.RaceService.GetRaceDataByID(raceID)
	if race == nil {
		c.String(http.StatusNotFound, "Race not found")
		return
	}

	judge := s.authenticatedJudge(c)
	c.HTML(http.StatusOK, "race", s.raceMenu(race, rating, judge))
}

func (s *Services) raceMenu(race *models.Race, rating *models.Rating, judge *models.Judge) gin.H {
	protests, err := s.Services.ProtestService.GetProtestsDataByRaceID(race.ID)
	if err != nil {
		log.Printf("Error getting races: %v", err)
		protests = []models.Protest{}
	}
	protestsData, err := modelsUI.FromProtestModelsToStringData(protests)
	if err != nil {
		log.Printf("Error class to string convert: %v", err)
	}

	raceView, _ := modelsUI.FromRaceModelToStringData(race)
	ratingView, _ := modelsUI.FromRatingModelToStringData(rating)

	var result = gin.H{
		"title":      "",
		"race":       race,
		"raceView":   raceView,
		"rating":     rating,
		"ratingView": ratingView,
		"protests":   protestsData,
		"judge":      judge,
	}

	return result
}

// CREATE

func (s *Services) createRaceGet(c *gin.Context) {
	c.HTML(200, "createRace", gin.H{
		"title":    "Создать гонку",
		"raceView": modelsUI.RaceInput{},
		"formData": modelsUI.RaceInput{},
	})
}

func (s *Services) createRacePost(c *gin.Context) {
	var input modelsUI.RaceInput
	err := c.Bind(&input)
	if err != nil {
		log.Printf("Error binding data: %v", err)
		c.HTML(400, "createRace", gin.H{
			"title":    "Создать гонку",
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	raceDate, err := parseDateTime(input.Date)
	if err != nil {
		c.HTML(http.StatusBadRequest, "createRace", gin.H{
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	ridStr := c.Param("ratingID")
	ratingID, err := uuid.Parse(ridStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат UUID")
		return
	}
	rating, err := s.Services.RatingService.GetRatingDataByID(ratingID)
	if rating == nil {
		c.String(http.StatusNotFound, "Rating not found")
		return
	}

	raceID := uuid.New()
	_, err = s.Services.RaceService.AddNewRace(raceID, rating.ID, input.Number, raceDate, rating.Class)
	if err != nil {
		log.Printf("Error creating race: %v", err)
		c.HTML(400, "createRace", gin.H{
			"title":    "Создать гонку",
			"error":    err.Error(),
			"formData": input,
			"classMap": modelsUI.ClassMap,
		})
		return
	}

	referer := c.Request.Referer()
	if referer == "" {
		referer = "/"
	} else {
		u, _ := url.Parse(referer)
		u.Path = u.Path[:strings.LastIndex(u.Path, "/")]
		u.Path = u.Path[:strings.LastIndex(u.Path, "/")]
		referer = u.String()
	}

	c.Redirect(http.StatusFound, referer)
}

// UPDATE
func (s *Services) updateRaceGet(c *gin.Context) {
	raceID, err := uuid.Parse(c.Param("raceID"))
	if err != nil {
		c.HTML(400, "updateRace", gin.H{
			"title": "Редактировать гонку",
			"error": "Неверный идентификатор гонки",
		})
		return
	}

	race, err := s.Services.RaceService.GetRaceDataByID(raceID)
	if err != nil {
		c.HTML(400, "updateRace", gin.H{
			"title": "Редактировать гонку",
			"error": "Гонка не найден",
		})
		return
	}

	editedRaceView, _ := modelsUI.FromRaceModelToInputData(race)

	c.HTML(200, "updateRace", gin.H{
		"title":    "Редактировать гонку",
		"raceView": race,
		"formData": editedRaceView,
	})
}

func (s *Services) updateRacePost(c *gin.Context) {
	var input modelsUI.RaceInput
	err := c.Bind(&input)
	if err != nil {
		log.Printf("Error binding data: %v", err)
		c.HTML(400, "updateRace", gin.H{
			"title":    "Редактировать гонку",
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	raceID, err := uuid.Parse(c.Param("raceID"))
	if err != nil {
		c.HTML(400, "updateRace", gin.H{
			"title": "Редактировать профиль",
			"error": "Неверный идентификатор гонки",
		})
		return
	}
	race, err := s.Services.RaceService.GetRaceDataByID(raceID)
	if err != nil {
		c.HTML(400, "updateRace", gin.H{
			"title":    "Редактировать гонку",
			"error":    "Гонка не найден",
			"formData": input,
		})
		return
	}

	raceDate, err := parseDateTime(input.Date)
	if err != nil {
		c.HTML(http.StatusBadRequest, "createRace", gin.H{
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	_, err = s.Services.RaceService.UpdateRaceByID(race.ID, race.RatingID, input.Number, raceDate, race.Class)
	if err != nil {
		log.Printf("Error updating race: %v", err)
		c.HTML(400, "updateRace", gin.H{
			"title":    "Редактировать гонку",
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	referer := c.Request.Referer()
	if referer == "" {
		referer = "/"
	} else {
		u, _ := url.Parse(referer)
		u.Path = u.Path[:strings.LastIndex(u.Path, "/")]
		referer = u.String()
	}

	c.Redirect(http.StatusFound, referer)
}

// DELETE
func (s *Services) deleteRace(c *gin.Context) {
	raceID, err := uuid.Parse(c.Param("raceID"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error-page", gin.H{
			"error": err.Error(),
		})
		return
	}

	err = s.Services.RaceService.DeleteRaceByID(raceID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error-page", gin.H{
			"error": err.Error(),
		})
		return
	}

	referer := c.Request.Referer()
	if referer == "" {
		referer = "/"
	} else {
		u, _ := url.Parse(referer)
		u.Path = u.Path[:strings.LastIndex(u.Path, "/")]
		u.Path = u.Path[:strings.LastIndex(u.Path, "/")]
		referer = u.String()
	}

	c.Redirect(http.StatusFound, referer)
}

// START

func (s *Services) startRaceGet(c *gin.Context) {
	ridStr := c.Param("ratingID")
	ratingID, err := uuid.Parse(ridStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат UUID")
		return
	}
	rating, err := s.Services.RatingService.GetRatingDataByID(ratingID)
	if rating == nil {
		c.String(http.StatusNotFound, "Rating not found")
		return
	}

	idStr := c.Param("raceID")
	raceID, err := uuid.Parse(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат UUID")
		return
	}
	race, err := s.Services.RaceService.GetRaceDataByID(raceID)
	if race == nil {
		c.String(http.StatusNotFound, "Race not found")
		return
	}
	crews, err := s.Services.CrewService.GetCrewsDataByRatingID(ratingID)
	if err != nil {
		c.String(http.StatusBadRequest, "Crews not found")
		return
	}
	crewsView, _ := modelsUI.FromCrewModelsToStringData(crews)
	judge := s.authenticatedJudge(c)
	raceView, _ := modelsUI.FromRaceModelToStringData(race)
	ratingView, _ := modelsUI.FromRatingModelToStringData(rating)
	var result = gin.H{
		"title":         "",
		"race":          race,
		"raceView":      raceView,
		"rating":        rating,
		"ratingView":    ratingView,
		"judge":         judge,
		"circumstances": modelsUI.SpecCircumstanceMap,
		"crewsView":     crewsView,
		"formData":      []modelsUI.StartInput{},
	}
	c.HTML(http.StatusOK, "start", result)
}

func (s *Services) startRacePost(c *gin.Context) {
	var formData []modelsUI.StartInput

	// Привязка данных запроса к структуре formData
	if err := c.ShouldBind(&formData); err != nil {
		c.HTML(http.StatusBadRequest, "start", gin.H{
			"title":    "Процедура старта",
			"error":    "Ошибка привязки данных: " + err.Error(),
			"formData": formData,
		})
		return
	}

	falseStartYachtList := make(map[int]int)
	for _, el := range formData {
		if el.SpecCirc != 0 {
			falseStartYachtList[el.SailNum] = el.SpecCirc
		}
	}

	raceID, err := uuid.Parse(c.Param("raceID"))
	if err != nil {
		c.HTML(400, "makeStartProcedure", gin.H{
			"title":    "Процедура старта",
			"error":    "Неверный идентификатор гонки",
			"formData": nil,
		})
		return
	}
	ridStr := c.Param("ratingID")
	ratingID, err := uuid.Parse(ridStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат UUID")
		return
	}

	err = s.Services.RaceService.MakeStartProcedure(raceID, falseStartYachtList)
	if err != nil {
		c.HTML(500, "start", gin.H{
			"title":    "Процедура старта",
			"error":    err.Error(),
			"formData": formData,
		})
		return
	}

	c.Redirect(302, "/ratings/"+ratingID.String()+"/races/"+raceID.String())
}
