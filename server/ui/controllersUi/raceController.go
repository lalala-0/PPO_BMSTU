package controllersUi

import (
	"PPO_BMSTU/internal/models"
	modelsUI2 "PPO_BMSTU/server/ui/modelsUI"
	"PPO_BMSTU/server/ui/uiUtils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (s *ServicesUI) getRaceMenu(c *gin.Context) {
	ridStr := c.Param("ratingID")
	ratingID, err := uuid.Parse(ridStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат UUID")
		return
	}
	rating, err := s.Services.RatingService.GetRatingDataByID(ratingID)
	if rating == nil || err != nil {
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
	if race == nil || err != nil {
		c.String(http.StatusNotFound, "Race not found")
		return
	}

	judge := s.authenticatedJudge(c)
	c.HTML(http.StatusOK, "race", s.raceMenu(race, rating, judge))
}

func (s *ServicesUI) raceMenu(race *models.Race, rating *models.Rating, judge *models.Judge) gin.H {
	protests, err := s.Services.ProtestService.GetProtestsDataByRaceID(race.ID)
	if err != nil {
		log.Printf("Error getting races: %v", err)
		protests = []models.Protest{}
	}
	protestsData, err := modelsUI2.FromProtestModelsToStringData(protests)
	if err != nil {
		log.Printf("Error class to string convert: %v", err)
	}

	raceView, _ := modelsUI2.FromRaceModelToStringData(race)
	ratingView, _ := modelsUI2.FromRatingModelToStringData(rating)

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

func (s *ServicesUI) createRaceGet(c *gin.Context) {
	c.HTML(200, "createRace", gin.H{
		"title":    "Создать гонку",
		"raceView": modelsUI2.RaceInput{},
		"formData": modelsUI2.RaceInput{},
	})
}

func (s *ServicesUI) createRacePost(c *gin.Context) {
	var input modelsUI2.RaceInput
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

	raceDate, err := uiUtils.ParseDateTime(input.Date)
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
	rating, _ := s.Services.RatingService.GetRatingDataByID(ratingID)
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
			"classMap": modelsUI2.ClassMap,
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
func (s *ServicesUI) updateRaceGet(c *gin.Context) {
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

	editedRaceView, _ := modelsUI2.FromRaceModelToInputData(race)

	c.HTML(200, "updateRace", gin.H{
		"title":    "Редактировать гонку",
		"raceView": race,
		"formData": editedRaceView,
	})
}

func (s *ServicesUI) updateRacePost(c *gin.Context) {
	var input modelsUI2.RaceInput
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

	raceDate, err := uiUtils.ParseDateTime(input.Date)
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
func (s *ServicesUI) deleteRace(c *gin.Context) {
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

func (s *ServicesUI) startRaceGet(c *gin.Context) {
	ridStr := c.Param("ratingID")
	ratingID, err := uuid.Parse(ridStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат UUID")
		return
	}
	rating, err := s.Services.RatingService.GetRatingDataByID(ratingID)
	if err != nil {
		c.HTML(http.StatusBadRequest, "start", gin.H{
			"title": "Процедура старта",
			"error": "Гонка не найден",
		})
		return
	}

	crews, err := s.Services.CrewService.GetCrewsDataByRatingID(ratingID)
	if err != nil {
		c.String(http.StatusBadRequest, "Crews not found")
		return
	}

	raceID, err := uuid.Parse(c.Param("raceID"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "start", gin.H{
			"title": "Процедура старта",

			"error": "Неверный идентификатор гонки",
		})
		return
	}

	race, err := s.Services.RaceService.GetRaceDataByID(raceID)
	if err != nil {
		c.HTML(http.StatusBadRequest, "start", gin.H{
			"title": "Процедура старта",
			"error": "Гонка не найден",
		})
		return
	}

	c.HTML(200, "start", gin.H{
		"title":      "Старт",
		"crews":      crews,
		"raceView":   race,
		"ratingView": rating,
	})
}

func (s *ServicesUI) startRacePost(c *gin.Context) {
	raceID, err := uuid.Parse(c.Param("raceID"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "start", gin.H{
			"title": "Процедура старта",

			"error": "Неверный идентификатор гонки",
		})
		return
	}
	falseStartYachtList, err := uiUtils.ParseString(c.PostForm("falshStartSailNums"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "start", gin.H{
			"title": "Процедура старта",
			"error": "Некорректный формат номеров парусов фальш-стартовавших яхт.",
		})
		return
	}

	err = s.Services.RaceService.MakeStartProcedure(raceID, uiUtils.SliceToMapConst(falseStartYachtList))
	if err != nil {
		c.HTML(500, "start", gin.H{
			"title": "Процедура старта",
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
		referer = u.String()
	}

	c.Redirect(http.StatusFound, referer)
}

// FINISH

func (s *ServicesUI) finishRaceGet(c *gin.Context) {
	ridStr := c.Param("ratingID")
	ratingID, err := uuid.Parse(ridStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат UUID")
		return
	}
	rating, err := s.Services.RatingService.GetRatingDataByID(ratingID)
	if err != nil {
		c.HTML(http.StatusBadRequest, "start", gin.H{
			"title": "Процедура старта",
			"error": "Гонка не найден",
		})
		return
	}

	crews, err := s.Services.CrewService.GetCrewsDataByRatingID(ratingID)
	if err != nil {
		c.String(http.StatusBadRequest, "Crews not found")
		return
	}

	raceID, err := uuid.Parse(c.Param("raceID"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "start", gin.H{
			"title": "Процедура старта",

			"error": "Неверный идентификатор гонки",
		})
		return
	}

	race, err := s.Services.RaceService.GetRaceDataByID(raceID)
	if err != nil {
		c.HTML(http.StatusBadRequest, "start", gin.H{
			"title": "Процедура старта",
			"error": "Гонка не найден",
		})
		return
	}

	c.HTML(200, "finish", gin.H{
		"title":      "Финиш",
		"crews":      crews,
		"ratingView": rating,
		"raceView":   race,
	})
}

func (s *ServicesUI) finishRacePost(c *gin.Context) {
	rID := c.Param("raceID")
	raceID, err := uuid.Parse(rID)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат UUID")
		return
	}
	finisherList, err := uiUtils.ParseString(c.PostForm("finishSailNums"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "finish", gin.H{
			"title": "Процедура финииша",
			"error": "Некорректный формат номеров парусов финишировавших яхт.",
		})
		return
	}

	err = s.Services.RaceService.MakeFinishProcedure(raceID, uiUtils.SliceToMapSerial(finisherList), make(map[int]int))
	if err != nil {
		c.HTML(500, "finish", gin.H{
			"title": "Процедура финиша",
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
		referer = u.String()
	}

	c.Redirect(http.StatusFound, referer)
}
