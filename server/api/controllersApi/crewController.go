package controllersApi

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/server/api/apiUtils"
	"PPO_BMSTU/server/api/modelsViewApi"
	"PPO_BMSTU/server/ui/modelsUI"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

// getCrewsByRatingID godoc
// @Summary Получить все команды в рейтинге
// @Tags Crew
// @Produce json
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Success 200 {array} modelsViewApi.CrewFormData "Список команд успешно получен"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Не найдены рейтинг или команды"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 500 {object} modelsViewApi.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/ratings/{ratingID}/crews [get]
func (s *ServicesAPI) getCrewsByRatingID(c *gin.Context) {
	ratingIDParam := c.Param("ratingID")

	ratingID, err := uuid.Parse(ratingIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid rating ID",
			Message: "The provided rating ID is not a valid UUID.",
		})
		return
	}

	crews, err := s.Services.CrewService.GetCrewsDataByRatingID(ratingID)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.BadRequestError{
				Error: "Internal error",
			})
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.BadRequestError{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}
	res, err := modelsViewApi.FromCrewModelsToStringData(crews)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modelsViewApi.BadRequestError{
			Error:   "Internal error",
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusOK, res)
}

// createCrew godoc
// @Summary Создать новую команду
// @Tags Crew
// @Produce json
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param data body modelsViewApi.CrewInput true "Данные для создания новой команды"
// @Success 201 {object} modelsViewApi.CrewFormData "Команда успешно создана"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Рейтинг не найден"
// @Failure 500 {object} modelsViewApi.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/ratings/{ratingID}/crews [post]
func (s *ServicesAPI) createCrew(c *gin.Context) {
	ratingIDParam := c.Param("ratingID")
	ratingID, err := uuid.Parse(ratingIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid rating ID",
			Message: "The provided rating ID is not a valid UUID.",
		})
		return
	}

	var input modelsViewApi.CrewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid request body",
			Message: "The request body contains invalid data.",
		})
		return
	}

	if input.SailNum <= 0 {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid sail number",
			Message: "Sail number must be a positive integer.",
		})
		return
	}

	// Проверка наличия рейтинга
	rating, err := s.Services.RatingService.GetRatingDataByID(ratingID)
	if err == repository_errors.DoesNotExist {
		c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
			Error:   "Rating not found",
			Message: "The specified rating ID does not exist.",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: "Failed to get rating.",
		})
	}

	createdCrew, err := s.Services.CrewService.AddNewCrew(uuid.New(), ratingID, rating.Class, input.SailNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: "Failed to create crew.",
		})
		return
	}

	c.JSON(http.StatusCreated, createdCrew)
}

// getCrewByID godoc
// @Summary Получить информацию о команде
// @Tags Crew
// @Produce json
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param crewID path string true "Уникальный идентификатор команды" format(uuid)
// @Success 200 {object} modelsViewApi.CrewFormData "Информация о команде успешно получена"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Рейтинг или команда не найдены"
// @Failure 500 {object} modelsViewApi.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/ratings/{ratingID}/crews/{crewID} [get]
func (s *ServicesAPI) getCrewByID(c *gin.Context) {
	crewIDParam := c.Param("crewID")

	crewID, err := uuid.Parse(crewIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid crew ID",
			Message: "The provided crew ID is not a valid UUID.",
		})
		return
	}

	crew, err := s.Services.CrewService.GetCrewDataByID(crewID)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Crew not found",
				Message: "The specified crew ID does not exist.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	crewData, err := modelsViewApi.FromCrewModelToStringData(crew)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: "Failed to convert crew data.",
		})
		return
	}

	c.JSON(http.StatusOK, crewData)
}

// updateCrewSailNumber godoc
// @Summary Обновить номер паруса команды
// @Tags Crew
// @Produce json
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param crewID path string true "Уникальный идентификатор команды" format(uuid)
// @Param data body modelsViewApi.CrewInput true "Данные для обновления команды"
// @Success 200 {object} modelsViewApi.CrewFormData "Команда успешно обновлена"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Рейтинг или команда не найдены"
// @Failure 500 {object} modelsViewApi.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/ratings/{ratingID}/crews/{crewID} [put]
func (s *ServicesAPI) updateCrewSailNumber(c *gin.Context) {
	ratingIDParam := c.Param("ratingID")
	crewIDParam := c.Param("crewID")

	ratingID, err := uuid.Parse(ratingIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid rating ID",
			Message: "The provided rating ID is not a valid UUID.",
		})
		return
	}

	crewID, err := uuid.Parse(crewIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid crew ID",
			Message: "The provided crew ID is not a valid UUID.",
		})
		return
	}

	var input modelsViewApi.CrewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid request body",
			Message: "The request body contains invalid data.",
		})
		return
	}

	if input.SailNum <= 0 {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid sail number",
			Message: "Sail number must be a positive integer.",
		})
		return
	}

	rating, err := s.Services.RatingService.GetRatingDataByID(ratingID) // Логика для проверки существования рейтинга
	if err == repository_errors.DoesNotExist {
		c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
			Error:   "Rating not found",
			Message: "The specified rating ID does not exist.",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
	}

	updatedCrewData, err := s.Services.CrewService.UpdateCrewByID(crewID, ratingID, rating.Class, input.SailNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}
	updCrew, err := modelsUI.FromCrewModelToStringData(updatedCrewData)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid sail number",
			Message: "Updated crew data are not valid.",
		})
		return
	}

	c.JSON(http.StatusOK, updCrew)
}

// deleteCrewByID godoc
// @Summary Удалить команду
// @Tags Crew
// @Produce json
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param crewID path string true "Уникальный идентификатор команды" format(uuid)
// @Success 204 "Команда успешно удалена"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Рейтинг или команда не найдены"
// @Failure 500 {object} modelsViewApi.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/ratings/{ratingID}/crews/{crewID} [delete]
func (s *ServicesAPI) deleteCrewByID(c *gin.Context) {
	crewIDParam := c.Param("crewID")

	crewID, err := uuid.Parse(crewIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid crew ID",
			Message: "The provided crew ID is not a valid UUID.",
		})
		return
	}

	// Удаление команды
	if err := s.Services.CrewService.DeleteCrewByID(crewID); err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Crew not found",
				Message: "The specified crew ID does not exist.",
			})
		} else {
			c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
				Error:   "Internal error",
				Message: "Failed to delete crew.",
			})
		}
		return
	}

	// Возвращаем успешный ответ с кодом 204 (No Content)
	c.Status(http.StatusNoContent)
}

// getCrewMembersByID godoc
// @Summary Получить список участников команды
// @Tags Crew Members
// @Produce json
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param crewID path string true "Уникальный идентификатор команды" format(uuid)
// @Param name query string false "Фильтр по имени участника"
// @Param helmsman query int false "Фильтр по статусу рулевого (1 – рулевой, 0 – не рулевой)"
// @Success 200 {array} modelsViewApi.ParticipantFormData "Список участников команды"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Рейтинг, команда или участники не найдены"
// @Failure 500 {object} modelsViewApi.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/ratings/{ratingID}/crews/{crewID}/members [get]
func (s *ServicesAPI) getCrewMembersByID(c *gin.Context) {
	crewIDParam := c.Param("crewID")

	crewID, err := uuid.Parse(crewIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid crew ID",
			Message: "The provided crew ID is not a valid UUID.",
		})
		return
	}

	// Получение фильтров из параметров запроса
	nameFilter := c.Query("name")

	// Получение списка участников команды
	members, err := s.Services.ParticipantService.GetParticipantsDataByCrewID(crewID)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Participants not found",
				Message: "No participants found for the specified crew.",
			})
		} else {
			c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
				Error:   "Internal error",
				Message: err.Error(),
			})
		}
		return
	}

	// Фильтрация участников по имени, если фильтр задан
	if nameFilter != "" {
		var filteredMembers []models.Participant
		for _, member := range members {
			if strings.Contains(strings.ToLower(member.FIO), strings.ToLower(nameFilter)) {
				filteredMembers = append(filteredMembers, member)
			}
		}
		members = filteredMembers
	}
	resMembers, err := modelsViewApi.FromParticipantModelsToStringData(members)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid members data",
			Message: "Invalid members data.",
		})
		return
	}

	c.JSON(http.StatusOK, resMembers)
}

// attachCrewMember godoc
// @Summary Добавить участника в команду
// @Tags Crew Members
// @Produce json
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param crewID path string true "Уникальный идентификатор команды" format(uuid)
// @Param body body modelsViewApi.CrewParticipantAttachInput true "Данные для добавления участника"
// @Success 201 {string} string "Участник успешно добавлен"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Рейтинг или команда не найдены"
// @Failure 500 {object} modelsViewApi.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/ratings/{ratingID}/crews/{crewID}/members [post]
func (s *ServicesAPI) attachCrewMember(c *gin.Context) {
	crewIDParam := c.Param("crewID")

	crewID, err := uuid.Parse(crewIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid crew ID",
			Message: "The provided crew ID is not a valid UUID.",
		})
		return
	}

	var input modelsViewApi.CrewParticipantAttachInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid input data",
			Message: "Failed to parse request body.",
		})
		return
	}

	participantID, err := uuid.Parse(input.ParticipantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid crew ID",
			Message: "The provided crew ID is not a valid UUID.",
		})
		return
	}

	err = s.Services.CrewService.AttachParticipantToCrew(participantID, crewID, input.Helmsman)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Crew or rating not found",
				Message: "The specified crew or rating ID does not exist.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, "Participant successfully attached")
}

// getCrewMember godoc
// @Summary Получить информацию об участнике команды
// @Tags Crew Members
// @Produce json
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param crewID path string true "Уникальный идентификатор команды" format(uuid)
// @Param participantID path string true "Уникальный идентификатор участника" format(uuid)
// @Success 200 {object} modelsViewApi.ParticipantFormData "Информация об участнике успешно получена"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Команда, рейтинг или участник не найдены"
// @Failure 500 {object} modelsViewApi.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/ratings/{ratingID}/crews/{crewID}/members/{participantID} [get]
func (s *ServicesAPI) getCrewMember(c *gin.Context) {
	participantIDParam := c.Param("participantID")

	participantID, err := uuid.Parse(participantIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid participant ID",
			Message: "The provided participant ID is not a valid UUID.",
		})
		return
	}

	participant, err := s.Services.ParticipantService.GetParticipantDataByID(participantID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Crew, rating or participant not found",
				Message: "The specified crew, rating or participant ID does not exist.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	resParticipant, err := modelsViewApi.FromParticipantModelToStringData(participant)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid participant data",
			Message: "Participant data are not valid.",
		})
	}

	c.JSON(http.StatusOK, resParticipant)
}

// updateCrewMember godoc
// @Summary Изменить информацию об участнике команды
// @Tags Crew Members
// @Produce json
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param crewID path string true "Уникальный идентификатор команды" format(uuid)
// @Param participantID path string true "Уникальный идентификатор участника" format(uuid)
// @Param body body modelsViewApi.ParticipantInput true "Данные для обновления участника"
// @Success 200 {object} modelsViewApi.ParticipantFormData "Информация об участнике успешно обновлена"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Команда, рейтинг или участник не найдены"
// @Failure 500 {object} modelsViewApi.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/ratings/{ratingID}/crews/{crewID}/members/{participantID} [put]
func (s *ServicesAPI) updateCrewMember(c *gin.Context) {
	participantIDParam := c.Param("participantID")

	participantID, err := uuid.Parse(participantIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid participant ID",
			Message: "The provided participant ID is not a valid UUID.",
		})
		return
	}

	// Валидация тела запроса
	var input modelsViewApi.ParticipantInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid input data",
			Message: "Failed to parse request body.",
		})
		return
	}
	birthday := apiUtils.ConvertStringToTime(input.Birthday)
	// Обновление информации об участнике
	updatedParticipant, err := s.Services.ParticipantService.UpdateParticipantByID(participantID, input.FIO, input.Category, birthday, input.Coach)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Crew, rating or participant not found",
				Message: "The specified crew, rating or participant ID does not exist.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	// Возврат успешного ответа
	c.JSON(http.StatusOK, updatedParticipant)
}

// detachCrewMember godoc
// @Summary Удалить участника из команды
// @Tags Crew Members
// @Produce json
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param crewID path string true "Уникальный идентификатор команды" format(uuid)
// @Param participantID path string true "Уникальный идентификатор участника" format(uuid)
// @Success 204 {string} string "Участник успешно удален"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Рейтинг или команда не найдены"
// @Failure 500 {object} modelsViewApi.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/ratings/{ratingID}/crews/{crewID}/members/{participantID} [delete]
func (s *ServicesAPI) detachCrewMember(c *gin.Context) {
	crewIDParam := c.Param("crewID")
	participantIDParam := c.Param("participantID")

	crewID, err := uuid.Parse(crewIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid crew ID",
			Message: "The provided crew ID is not a valid UUID.",
		})
		return
	}

	participantID, err := uuid.Parse(participantIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid participant ID",
			Message: "The provided participant ID is not a valid UUID.",
		})
		return
	}

	// Удаление участника из команды
	err = s.Services.CrewService.DetachParticipantFromCrew(participantID, crewID)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Crew or rating not found",
				Message: "The specified crew or rating ID does not exist.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent) // 204 No Content
}
