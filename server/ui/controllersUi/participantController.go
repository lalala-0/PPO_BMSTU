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
	"net/url"
	"strings"
)

func (s *ServicesUI) getParticipantMenu(c *gin.Context) {
	participant, err := s.getParticipantByID(c)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	judge := s.authenticatedJudge(c)
	rating, crew, err := s.getRatingAndCrew(c)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.HTML(http.StatusOK, "participant", s.participantMenu(participant, crew, rating, judge))
}

// Вспомогательная функция для получения данных участника
func (s *ServicesUI) getParticipantByID(c *gin.Context) (*models.Participant, error) {
	idStr := c.Param("participantID")
	participantID, err := uuid.Parse(idStr)
	if err != nil {
		return nil, fmt.Errorf("Неверный формат UUID")
	}
	participant, err := s.Services.ParticipantService.GetParticipantDataByID(participantID)
	if participant == nil || err != nil {
		return nil, fmt.Errorf("Crew not found")
	}
	return participant, nil
}

// Вспомогательная функция для получения данных рейтинга и экипажа
func (s *ServicesUI) getRatingAndCrew(c *gin.Context) (*models.Rating, *models.Crew, error) {
	if c.Param("ratingID") != "" {
		ratingID, err := uuid.Parse(c.Param("ratingID"))
		if err != nil {
			return nil, nil, fmt.Errorf("Неверный формат UUID")
		}
		rating, err := s.Services.RatingService.GetRatingDataByID(ratingID)
		if rating == nil || err != nil {
			return nil, nil, fmt.Errorf("Rating not found")
		}

		crewID, err := uuid.Parse(c.Param("crewID"))
		if err != nil {
			return nil, nil, fmt.Errorf("Неверный формат UUID")
		}
		crew, err := s.Services.CrewService.GetCrewDataByID(crewID)
		if crew == nil || err != nil {
			return nil, nil, fmt.Errorf("Crew not found")
		}
		return rating, crew, nil
	}
	return nil, nil, nil
}

func (s *ServicesUI) participantMenu(participant *models.Participant, crew *models.Crew, rating *models.Rating, judge *models.Judge) gin.H {
	participantView, _ := modelsUI2.FromParticipantModelToStringData(participant)
	var result gin.H

	if rating == nil {
		result = gin.H{
			"title":           "",
			"judge":           judge,
			"participant":     participant,
			"participantView": participantView,
		}
	} else {
		ratingView, _ := modelsUI2.FromRatingModelToStringData(rating)
		crewView, _ := modelsUI2.FromCrewModelToStringData(crew)

		result = gin.H{
			"title":           "",
			"judge":           judge,
			"participant":     participant,
			"participantView": participantView,
			"crew":            crew,
			"crewView":        crewView,
			"rating":          rating,
			"ratingView":      ratingView,
		}
	}
	return result
}

// DELETE
func (s *ServicesUI) deleteParticipant(c *gin.Context) {
	participantID, err := uuid.Parse(c.Param("participantID"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error-page", gin.H{
			"error": err.Error(),
		})
		return
	}

	err = s.Services.ParticipantService.DeleteParticipantByID(participantID)
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

// UPDATE
func (s *ServicesUI) updateParticipantGet(c *gin.Context) {
	participantID, err := uuid.Parse(c.Param("participantID"))
	if err != nil {
		c.HTML(400, "updateParticipant", gin.H{
			"title": "Редактировать участника",
			"error": "Неверный идентификатор участника",
		})
		return
	}

	participant, err := s.Services.ParticipantService.GetParticipantDataByID(participantID)
	if err != nil {
		c.HTML(400, "updateParticipant", gin.H{
			"title": "Редактировать участника",
			"error": "Участник не найден",
		})
		return
	}

	editedParticipantView, _ := modelsUI2.FromParticipantModelToInputData(participant)

	c.HTML(200, "updateParticipant", gin.H{
		"title":           "Редактировать участника",
		"participantView": participant,
		"formData":        editedParticipantView,
		"categoryMap":     modelsUI2.CategoryMap,
	})
}

func (s *ServicesUI) updateParticipantPost(c *gin.Context) {
	var input modelsUI2.ParticipantInput
	err := c.Bind(&input)
	if err != nil {
		log.Printf("Error binding data: %v", err)
		c.HTML(400, "updateParticipant", gin.H{
			"title":    "Редактировать участника",
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	participantID, err := uuid.Parse(c.Param("participantID"))
	if err != nil {
		c.HTML(400, "updateParticipant", gin.H{
			"title": "Редактировать участника",
			"error": "Неверный идентификатор участника",
		})
		return
	}

	birthday, err := uiUtils.ParseDateTime(input.Birthday)
	if err != nil {
		c.HTML(400, "updateParticipant", gin.H{
			"title": "Редактировать участника",
			"error": "Неверный идентификатор участника",
		})
		return
	}

	_, err = s.Services.ParticipantService.UpdateParticipantByID(participantID, input.FIO, input.Category, birthday, input.Coach)
	if err != nil {
		log.Printf("Error updating participant: %v", err)
		c.HTML(400, "updateParticipant", gin.H{
			"title":    "Редактировать участника",
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

// CREATE
func (s *ServicesUI) createParticipantGet(c *gin.Context) {
	c.HTML(200, "createParticipant", gin.H{
		"title":       "Создать участника",
		"formData":    modelsUI2.ParticipantInput{},
		"categoryMap": modelsUI2.CategoryMap,
		"genderMap":   modelsUI2.GenderMap,
	})
}

func (s *ServicesUI) createParticipantPost(c *gin.Context) {
	var input modelsUI2.ParticipantInput
	err := c.Bind(&input)
	if err != nil {
		log.Printf("Error binding data: %v", err)
		c.HTML(400, "createParticipant", gin.H{
			"title":    "Создать участника",
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	birthday, err := uiUtils.ParseDateTime(input.Birthday)
	if err != nil {
		c.HTML(400, "createParticipant", gin.H{
			"title": "Создать участника",
			"error": "Неверный формат даты",
		})
		return
	}

	_, err = s.Services.ParticipantService.AddNewParticipant(uuid.New(), input.FIO, input.Category, input.Gender, birthday, input.Coach)
	if err != nil {
		log.Printf("Error updating participant: %v", err)
		c.HTML(400, "createParticipant", gin.H{
			"title":    "Создать участника",
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
		u.Path = u.Path[:strings.LastIndex(u.Path, "/")]
		referer = u.String()
	}
	c.Redirect(http.StatusFound, referer)
}
