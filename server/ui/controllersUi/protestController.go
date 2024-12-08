package controllersUi

import (
	"PPO_BMSTU/internal/models"
	modelsUI2 "PPO_BMSTU/server/ui/modelsUI"
	"PPO_BMSTU/server/ui/uiUtils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func (s *ServicesUI) getProtestMenu(c *gin.Context) {
	protest, err := s.getProtestDataFromParam(c, "protestID")
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	protestCrews, err := s.getProtestParticipants(protest.ID)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	rating, err := s.getRatingByID(c)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	race, err := s.getRaceByID(c)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	judge := s.authenticatedJudge(c)

	c.HTML(http.StatusOK, "protest", s.protestMenu(protest, protestCrews, race, rating, judge))
}

// Вспомогательная функция для парсинга UUID из параметра
func (s *ServicesUI) parseUUIDParam(c *gin.Context, paramName string) (uuid.UUID, error) {
	idStr := c.Param(paramName)
	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("Неверный формат UUID для %s", paramName)
	}
	return id, nil
}

// Вспомогательная функция для получения данных о протесте с параметром protestID
func (s *ServicesUI) getProtestDataFromParam(c *gin.Context, paramName string) (*models.Protest, error) {
	protestID, err := s.parseUUIDParam(c, paramName)
	if err != nil {
		return nil, err
	}

	protest, err := s.getProtestByID(protestID)
	if err != nil {
		return nil, err
	}

	return protest, nil
}

// Вспомогательная функция для получения данных о протесте
func (s *ServicesUI) getProtestByID(protestID uuid.UUID) (*models.Protest, error) {
	protest, err := s.Services.ProtestService.GetProtestDataByID(protestID)
	if protest == nil || err != nil {
		return nil, fmt.Errorf("Protest not found")
	}
	return protest, nil
}

// Вспомогательная функция для получения участников протеста
func (s *ServicesUI) getProtestParticipants(protestID uuid.UUID) ([]modelsUI2.ProtestCrewFormData, error) {
	crewIDs, err := s.Services.ProtestService.GetProtestParticipantsIDByID(protestID)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при получении участников протеста")
	}
	var protestCrews []modelsUI2.ProtestCrewFormData
	for crewID, role := range crewIDs {
		crew, err := s.Services.CrewService.GetCrewDataByID(crewID)
		if err != nil {
			return nil, fmt.Errorf("Команда-участник протеста не найдена")
		}
		protestCrewView, _ := modelsUI2.FromProtestParticipantModelToStringData(crew, role)
		protestCrews = append(protestCrews, protestCrewView)
	}
	return protestCrews, nil
}

func (s *ServicesUI) protestMenu(protest *models.Protest, protestCrews []modelsUI2.ProtestCrewFormData, race *models.Race, rating *models.Rating, judge *models.Judge) gin.H {
	raceView, _ := modelsUI2.FromRaceModelToStringData(race)
	ratingView, _ := modelsUI2.FromRatingModelToStringData(rating)
	protestView, _ := modelsUI2.FromProtestModelToStringData(protest)
	var result = gin.H{
		"title":               "",
		"protest":             protest,
		"protestView":         protestView,
		"protestParticipants": protestCrews,

		"race":       race,
		"raceView":   raceView,
		"rating":     rating,
		"ratingView": ratingView,
		"judge":      judge,
	}

	return result
}

// CREATE
func (s *ServicesUI) createProtestGet(c *gin.Context) {
	ridStr := c.Param("ratingID")
	ratingID, err := uuid.Parse(ridStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат UUID")
		return
	}

	crews, err := s.Services.CrewService.GetCrewsDataByRatingID(ratingID)
	if err != nil {
		c.String(http.StatusBadRequest, "Crews not found")
		return
	}

	c.HTML(200, "createProtest", gin.H{
		"title":    "Создать протест",
		"crews":    crews,
		"formData": modelsUI2.ProtestCreate{},
	})
}

func (s *ServicesUI) createProtestPost(c *gin.Context) {
	var input modelsUI2.ProtestCreate
	err := c.ShouldBind(&input)
	if err != nil {
		s.renderProtestFormWithError(c, err.Error(), input)
		return
	}

	witnesses, err := uiUtils.ParseString(c.PostForm("witnesses"))
	if err != nil || len(witnesses) < 1 || len(witnesses) > 5 {
		s.renderProtestFormWithError(c, "Вы должны выбрать от одного до пяти яхт-свидетелей.", input)
		return
	}

	reviewDate, err := uiUtils.ParseDateTime(input.ReviewDate)
	if err != nil {
		s.renderProtestFormWithError(c, err.Error(), input)
		return
	}

	rating, err := s.getRatingByID(c)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	race, err := s.getRaceByID(c)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	judge := s.authenticatedJudge(c)

	_, err = s.Services.ProtestService.AddNewProtest(race.ID, rating.ID, judge.ID, input.RuleNum, reviewDate, input.Comment, input.ProtesteeSailNum, input.ProtestorSailNum, witnesses)
	if err != nil {
		log.Printf("Error creating protest: %v", err)
		s.renderProtestFormWithError(c, err.Error(), input)
		return
	}

	c.Redirect(302, "/ratings/"+rating.ID.String()+"/races/"+race.ID.String())
}

// Вспомогательная функция для рендеринга формы с ошибкой
func (s *ServicesUI) renderProtestFormWithError(c *gin.Context, errorMessage string, input modelsUI2.ProtestCreate) {
	c.HTML(400, "createProtest", gin.H{
		"title":    "Создать протест",
		"error":    errorMessage,
		"formData": input,
	})
}

// Вспомогательная функция для получения данных рейтинга
func (s *ServicesUI) getRatingByID(c *gin.Context) (*models.Rating, error) {
	ridStr := c.Param("ratingID")
	ratingID, err := uuid.Parse(ridStr)
	if err != nil {
		return nil, fmt.Errorf("Неверный формат UUID для рейтинга")
	}
	rating, _ := s.Services.RatingService.GetRatingDataByID(ratingID)
	if rating == nil {
		return nil, fmt.Errorf("Rating not found")
	}
	return rating, nil
}

// Вспомогательная функция для получения данных гонки
func (s *ServicesUI) getRaceByID(c *gin.Context) (*models.Race, error) {
	ridStr := c.Param("raceID")
	raceID, err := uuid.Parse(ridStr)
	if err != nil {
		return nil, fmt.Errorf("Неверный формат UUID для гонки")
	}
	race, _ := s.Services.RaceService.GetRaceDataByID(raceID)
	if race == nil {
		return nil, fmt.Errorf("Race not found")
	}
	return race, nil
}

// UPDATE

func (s *ServicesUI) updateProtestGet(c *gin.Context) {
	protestID, err := uuid.Parse(c.Param("protestID"))
	if err != nil {
		c.HTML(400, "updateProtest", gin.H{
			"title": "Редактировать протест",
			"error": "Неверный идентификатор протеста",
		})
		return
	}

	protest, err := s.Services.ProtestService.GetProtestDataByID(protestID)
	if err != nil {
		c.HTML(400, "updateProtest", gin.H{
			"title": "Редактировать протест",
			"error": "Протест не найден",
		})
		return
	}

	editedProtestView, _ := modelsUI2.FromProtestModelToInputData(protest)

	c.HTML(200, "updateProtest", gin.H{
		"title":       "Редактировать протест",
		"protestView": protest,
		"formData":    editedProtestView,
	})
}

func (s *ServicesUI) updateProtestPost(c *gin.Context) {
	var input modelsUI2.ProtestInput
	err := c.Bind(&input)
	if err != nil {
		log.Printf("Error binding data: %v", err)
		c.HTML(400, "updateProtest", gin.H{
			"title":    "Редактировать протест",
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	protestID, err := uuid.Parse(c.Param("protestID"))
	if err != nil {
		c.HTML(400, "updateProtest", gin.H{
			"title": "Редактировать протест",
			"error": "Неверный идентификатор протеста",
		})
		return
	}
	protest, err := s.Services.ProtestService.GetProtestDataByID(protestID)
	if err != nil {
		c.HTML(400, "updateProtest", gin.H{
			"title":    "Редактировать протест",
			"error":    "Протест не найден",
			"formData": input,
		})
		return
	}

	protestDate, err := uiUtils.ParseDateTime(input.ReviewDate)
	if err != nil {
		c.HTML(http.StatusBadRequest, "updateProtest", gin.H{
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	_, err = s.Services.ProtestService.UpdateProtestByID(protest.ID, protest.RaceID, protest.JudgeID, input.RuleNum, protestDate, protest.Status, input.Comment)
	if err != nil {
		log.Printf("Error updating protest: %v", err)
		c.HTML(400, "updateProtest", gin.H{
			"title":    "Редактировать протест",
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	c.Redirect(302, "/ratings/"+protest.RatingID.String()+"/races/"+protest.RaceID.String()+"/protests/"+protestID.String())
}

// COMPLETE

func (s *ServicesUI) completeProtestGet(c *gin.Context) {
	protestID, err := uuid.Parse(c.Param("protestID"))
	if err != nil {
		c.HTML(400, "completeProtest", gin.H{
			"title": "Завершить рассмотрение протеста",
			"error": "Неверный идентификатор протеста",
		})
		return
	}

	protest, err := s.Services.ProtestService.GetProtestDataByID(protestID)
	if err != nil {
		c.HTML(400, "completeProtest", gin.H{
			"title": "Завершить рассмотрение протеста",
			"error": "Протест не найден",
		})
		return
	}

	c.HTML(200, "completeProtest", gin.H{
		"title":       "Завершить рассмотрение протеста",
		"protestView": protest,
		"formData":    modelsUI2.ProtestComplete{},
	})
}

func (s *ServicesUI) completeProtestPost(c *gin.Context) {
	var input modelsUI2.ProtestComplete
	err := c.Bind(&input)
	if err != nil {
		log.Printf("Error binding data: %v", err)
		c.HTML(400, "completeProtest", gin.H{
			"title":    "Завершить рассмотрение протеста",
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	protestID, err := uuid.Parse(c.Param("protestID"))
	if err != nil {
		c.HTML(400, "completeProtest", gin.H{
			"title":    "Завершить рассмотрение протеста",
			"error":    "Неверный идентификатор протеста",
			"formData": input,
		})
		return
	}

	err = s.Services.ProtestService.CompleteReview(protestID, input.ResPoints, input.Comment)
	if err != nil {
		log.Printf("Error updating protest: %v", err)
		c.HTML(400, "completeProtest", gin.H{
			"title":    "Завершить рассмотрение протеста",
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	c.Redirect(302, "/ratings/"+c.Param("ratingID")+"/races/"+c.Param("raceID")+"/protests/"+protestID.String())
}

// DELETE
func (s *ServicesUI) deleteProtest(c *gin.Context) {
	protestID, err := uuid.Parse(c.Param("protestID"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error-page", gin.H{
			"error": err.Error(),
		})
		return
	}

	err = s.Services.ProtestService.DeleteProtestByID(protestID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error-page", gin.H{
			"error": err.Error(),
		})
		return
	}

	ratingID, err := uuid.Parse(c.Param("ratingID"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error-page", gin.H{
			"error": err.Error(),
		})
		return
	}

	raceID, err := uuid.Parse(c.Param("raceID"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error-page", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, "/ratings/"+ratingID.String()+"/races/"+raceID.String())
}

// ATTACH PROTEST PARTICIPANT

func (s *ServicesUI) attachProtestParticipantGet(c *gin.Context) {
	protestID, err := uuid.Parse(c.Param("protestID"))
	if err != nil {
		c.HTML(400, "attachProtestParticipant", gin.H{
			"title": "Добавить участника протеста",
			"error": "Неверный идентификатор протеста",
		})
		return
	}

	protest, err := s.Services.ProtestService.GetProtestDataByID(protestID)
	if err != nil {
		c.HTML(400, "attachProtestParticipant", gin.H{
			"title": "Добавить участника протеста",
			"error": "Протест не найден",
		})
		return
	}

	ProtestView, _ := modelsUI2.FromProtestModelToStringData(protest)

	crews, err := s.Services.CrewService.GetCrewsDataByRatingID(protest.RatingID)
	if err != nil {
		c.HTML(400, "attachProtestParticipant", gin.H{
			"title": "Добавить участника протеста",
			"error": "Команды-участники рейтинга не найдены",
		})
		return
	}

	c.HTML(200, "attachProtestParticipant", gin.H{
		"title":       "Добавить участника протеста",
		"protestView": ProtestView,
		"crews":       crews,
		"formData":    modelsUI2.ProtestParticipantAttachInput{},
		"roleMap":     modelsUI2.RoleMap,
	})
}

func (s *ServicesUI) attachProtestParticipantPost(c *gin.Context) {
	var input modelsUI2.ProtestParticipantAttachInput
	err := c.Bind(&input)
	if err != nil {
		log.Printf("Error binding data: %v", err)
		c.HTML(400, "attachProtestParticipant", gin.H{
			"title":    "Добавить участника протеста",
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	protestID, err := uuid.Parse(c.Param("protestID"))
	if err != nil {
		c.HTML(400, "attachProtestParticipant", gin.H{
			"title": "Добавить участника протеста",
			"error": "Неверный идентификатор протеста",
		})
		return
	}

	err = s.Services.ProtestService.AttachCrewToProtest(protestID, input.SailNum, input.Role)
	if err != nil {
		log.Printf("Error updating protest: %v", err)
		c.HTML(400, "attachProtestParticipant", gin.H{
			"title":    "Добавить участника протеста",
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	c.Redirect(302, "/ratings/"+c.Param("ratingID")+"/races/"+c.Param("raceID")+"/protests/"+protestID.String())
}

// DETACH PROTEST PARTICIPANT
func (s *ServicesUI) detachProtestParticipantGet(c *gin.Context) {
	protestID, err := uuid.Parse(c.Param("protestID"))
	if err != nil {
		c.HTML(400, "detachProtestParticipant", gin.H{
			"title": "Удалить участника протеста",
			"error": "Неверный идентификатор протеста",
		})
		return
	}

	protest, err := s.Services.ProtestService.GetProtestDataByID(protestID)
	if err != nil {
		c.HTML(400, "detachProtestParticipant", gin.H{
			"title": "Удалить участника протеста",
			"error": "Протест не найден",
		})
		return
	}

	ProtestView, _ := modelsUI2.FromProtestModelToStringData(protest)

	crewIDs, err := s.Services.ProtestService.GetProtestParticipantsIDByID(protest.ID)
	if err != nil {
		c.HTML(400, "detachProtestParticipant", gin.H{
			"title": "Удалить участника протеста",
			"error": "Команды-участники протеста не найдены",
		})
		return
	}
	var protestCrews = []modelsUI2.ProtestCrewFormData{}
	for crewID, role := range crewIDs {
		crew, err := s.Services.CrewService.GetCrewDataByID(crewID)
		if err != nil {
			c.String(http.StatusBadRequest, "Команда-участник протеста не найдена")
			return
		}
		protestCrewView, _ := modelsUI2.FromProtestParticipantModelToStringData(crew, role)
		protestCrews = append(protestCrews, protestCrewView)
	}

	c.HTML(200, "detachProtestParticipant", gin.H{
		"title":       "Удалить участника протеста",
		"protestView": ProtestView,
		"crews":       protestCrews,
		"formData":    modelsUI2.ProtestParticipantDetachInput{},
	})
}

func (s *ServicesUI) detachProtestParticipantPost(c *gin.Context) {
	var input modelsUI2.ProtestParticipantDetachInput
	err := c.Bind(&input)
	if err != nil {
		log.Printf("Error binding data: %v", err)
		c.HTML(400, "detachProtestParticipant", gin.H{
			"title":    "Удалить участника протеста",
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	protestID, err := uuid.Parse(c.Param("protestID"))
	if err != nil {
		c.HTML(400, "detachProtestParticipant", gin.H{
			"title": "Удалить участника протеста",
			"error": "Неверный идентификатор протеста",
		})
		return
	}

	err = s.Services.ProtestService.DetachCrewFromProtest(protestID, input.SailNum)
	if err != nil {
		log.Printf("Error detachProtestParticipant: %v", err)
		c.HTML(400, "detachProtestParticipant", gin.H{
			"title":    "Удалить участника протеста",
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	c.Redirect(302, "/ratings/"+c.Param("ratingID")+"/races/"+c.Param("raceID")+"/protests/"+protestID.String())
}
