package controllersUi

import (
	"PPO_BMSTU/internal/models"
	modelsUI2 "PPO_BMSTU/server/ui/modelsUI"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (s *ServicesUI) getCrewMenu(c *gin.Context) {
	idStr := c.Param("crewID")
	crewID, err := uuid.Parse(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат UUID")
		return
	}
	crew, err := s.Services.CrewService.GetCrewDataByID(crewID)
	if crew == nil {
		c.String(http.StatusNotFound, "Crew not found")
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

	participants, err := s.Services.ParticipantService.GetParticipantsDataByCrewID(crewID)
	//if participants == nil {
	//	c.String(http.StatusNotFound, "Crew participants not found")
	//	return
	//}
	judge := s.authenticatedJudge(c)

	c.HTML(http.StatusOK, "crew", s.crewMenu(crew, rating, judge, participants))
}

func (s *ServicesUI) crewMenu(crew *models.Crew, rating *models.Rating, judge *models.Judge, participants []models.Participant) gin.H {
	ratingView, _ := modelsUI2.FromRatingModelToStringData(rating)
	participantsView, _ := modelsUI2.FromParticipantModelsToStringData(participants)
	crewView, _ := modelsUI2.FromCrewModelToStringData(crew)

	var result = gin.H{
		"title":            "",
		"judge":            judge,
		"crew":             crew,
		"crewView":         crewView,
		"rating":           rating,
		"ratingView":       ratingView,
		"participants":     participants,
		"participantsView": participantsView,
	}

	return result
}

// CREATE
func (s *ServicesUI) createCrewGet(c *gin.Context) {
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

	c.HTML(200, "createCrew", gin.H{
		"title":    "Создать команду",
		"crews":    crews,
		"formData": modelsUI2.CrewInput{},
		"classMap": modelsUI2.ClassMap,
	})
}

func (s *ServicesUI) createCrewPost(c *gin.Context) {
	var input modelsUI2.CrewInput
	err := c.ShouldBind(&input)
	if err != nil {
		log.Printf("Error binding data: %v", err)
		c.HTML(400, "createCrew", gin.H{
			"title":    "Создать команду",
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
	_, err = s.Services.CrewService.AddNewCrew(uuid.New(), ratingID, input.Class, input.SailNum)
	if err != nil {
		log.Printf("Error creating crew: %v", err)
		c.HTML(400, "createCrew", gin.H{
			"title":    "Создать команду",
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

// UPDATE
func (s *ServicesUI) updateCrewGet(c *gin.Context) {
	crewID, err := uuid.Parse(c.Param("crewID"))
	if err != nil {
		c.HTML(400, "updateCrew", gin.H{
			"title": "Редактировать команду",
			"error": "Неверный идентификатор команды",
		})
		return
	}

	crew, err := s.Services.CrewService.GetCrewDataByID(crewID)
	if err != nil {
		c.HTML(400, "updateCrew", gin.H{
			"title": "Редактировать команду",
			"error": "Команда не найдена",
		})
		return
	}

	editedCrewView, _ := modelsUI2.FromCrewModelToInputData(crew)

	c.HTML(200, "updateCrew", gin.H{
		"title":    "Редактировать команду",
		"crewView": crew,
		"formData": editedCrewView,
		"classMap": modelsUI2.ClassMap,
	})
}

func (s *ServicesUI) updateCrewPost(c *gin.Context) {
	var input modelsUI2.CrewInput
	err := c.Bind(&input)
	if err != nil {
		log.Printf("Error binding data: %v", err)
		c.HTML(400, "updateCrew", gin.H{
			"title":    "Редактировать команду",
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	crewID, err := uuid.Parse(c.Param("crewID"))
	if err != nil {
		c.HTML(400, "updateCrew", gin.H{
			"title": "Редактировать команду",
			"error": "Неверный идентификатор команды",
		})
		return
	}

	ratingID, err := uuid.Parse(c.Param("ratingID"))
	if err != nil {
		c.HTML(400, "updateCrew", gin.H{
			"title": "Редактировать команду",
			"error": "Неверный идентификатор рейтинга",
		})
		return
	}

	_, err = s.Services.CrewService.UpdateCrewByID(crewID, ratingID, input.Class, input.SailNum)
	if err != nil {
		log.Printf("Error updating crew: %v", err)
		c.HTML(400, "updateCrew", gin.H{
			"title":    "Редактировать команду",
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
func (s *ServicesUI) deleteCrew(c *gin.Context) {
	crewID, err := uuid.Parse(c.Param("crewID"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error-page", gin.H{
			"error": err.Error(),
		})
		return
	}

	err = s.Services.CrewService.DeleteCrewByID(crewID)
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

// ATTACH CREW PARTICIPANT
func (s *ServicesUI) attachCrewParticipantGet(c *gin.Context) {
	crewID, err := uuid.Parse(c.Param("crewID"))
	if err != nil {
		c.HTML(400, "attachCrewParticipant", gin.H{
			"title": "Добавить участника команды",
			"error": "Неверный идентификатор команды",
		})
		return
	}

	crew, err := s.Services.CrewService.GetCrewDataByID(crewID)
	if err != nil {
		c.HTML(400, "attachCrewParticipant", gin.H{
			"title": "Добавить участника команды",
			"error": "Команда не найдена",
		})
		return
	}

	CrewView, _ := modelsUI2.FromCrewModelToStringData(crew)

	participantModels, err := s.Services.ParticipantService.GetAllParticipants()
	if err != nil {
		c.HTML(400, "attachCrewParticipant", gin.H{
			"title": "Добавить участника команды",
			"error": "Участники не найдены",
		})
		return
	}
	participants, _ := modelsUI2.FromParticipantModelsToInputData(participantModels)

	c.HTML(200, "attachCrewParticipant", gin.H{
		"title":        "Добавить участника команды",
		"crewView":     CrewView,
		"participants": participants,
		"formData":     modelsUI2.CrewParticipantAttachInput{Helmsman: 0},
		"helmsmanMap":  modelsUI2.HelmsmanMap,
	})
}

func (s *ServicesUI) attachCrewParticipantPost(c *gin.Context) {
	var input modelsUI2.CrewParticipantAttachInput
	err := c.Bind(&input)
	if err != nil {
		log.Printf("Error binding data: %v", err)
		c.HTML(400, "attachCrewParticipant", gin.H{
			"title":    "Добавить участника команды",
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	crewID, err := uuid.Parse(c.Param("crewID"))
	if err != nil {
		c.HTML(400, "attachCrewParticipant", gin.H{
			"title":    "Добавить участника команды",
			"error":    "Неверный идентификатор команды",
			"formData": input,
		})
		return
	}

	participantID, err := uuid.Parse(input.ParticipantID)
	if err != nil {
		c.HTML(400, "attachCrewParticipant", gin.H{
			"title":    "Добавить участника команды",
			"error":    "Неверный идентификатор участника",
			"formData": input,
		})
		return
	}

	err = s.Services.CrewService.AttachParticipantToCrew(participantID, crewID, input.Helmsman)
	if err != nil {
		log.Printf("Error updating crew: %v", err)
		c.HTML(400, "attachCrewParticipant", gin.H{
			"title":    "Добавить участника команды",
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

// DETACH CREW PARTICIPANT
func (s *ServicesUI) detachCrewParticipantGet(c *gin.Context) {
	crewID, err := uuid.Parse(c.Param("crewID"))
	if err != nil {
		c.HTML(400, "detachCrewParticipant", gin.H{
			"title": "Удалить участника команды",
			"error": "Неверный идентификатор команды",
		})
		return
	}

	crew, err := s.Services.CrewService.GetCrewDataByID(crewID)
	if err != nil {
		c.HTML(400, "detachCrewParticipant", gin.H{
			"title": "Удалить участника команды",
			"error": "Команда не найдена",
		})
		return
	}

	CrewView, _ := modelsUI2.FromCrewModelToStringData(crew)

	participantModels, err := s.Services.ParticipantService.GetAllParticipants()
	participants, err := modelsUI2.FromParticipantModelsToInputData(participantModels)

	c.HTML(200, "detachCrewParticipant", gin.H{
		"title":        "Удалить участника команды",
		"crewView":     CrewView,
		"participants": participants,
		"formData":     modelsUI2.CrewParticipantDetachInput{},
	})
}

func (s *ServicesUI) detachCrewParticipantPost(c *gin.Context) {
	var input modelsUI2.CrewParticipantDetachInput
	err := c.Bind(&input)
	if err != nil {
		log.Printf("Error binding data: %v", err)
		c.HTML(400, "detachCrewParticipant", gin.H{
			"title":    "Удалить участника команды",
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	crewID, err := uuid.Parse(c.Param("crewID"))
	if err != nil {
		c.HTML(400, "detachCrewParticipant", gin.H{
			"title": "Удалить участника команды",
			"error": "Неверный идентификатор команды",
		})
		return
	}

	participantID, err := uuid.Parse(input.ParticipantID)
	if err != nil {
		c.HTML(400, "attachCrewParticipant", gin.H{
			"title":    "Добавить участника команды",
			"error":    "Неверный идентификатор участника",
			"formData": input,
		})
		return
	}

	err = s.Services.CrewService.DetachParticipantFromCrew(participantID, crewID)
	if err != nil {
		log.Printf("Error detachCrewParticipant: %v", err)
		c.HTML(400, "detachCrewParticipant", gin.H{
			"title":    "Удалить участника команды",
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
