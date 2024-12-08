package controllersApi

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/server/api/apiUtils"
	"PPO_BMSTU/server/api/modelsViewApi"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

// getAllParticipants godoc
// @Summary Получить список всех участников
// @Description Получает список всех участников с возможностью фильтрации
// @Tags Participant
// @Produce json
// @Param fio query string false "Фильтр по ФИО"
// @Param category query string false "Фильтр по категории"
// @Param gender query string false "Фильтр по полу"
// @Param birthday query string false "Фильтр по дате рождения"
// @Param coach query string false "Фильтр по тренеру"
// @Success 200 {array} modelsViewApi.ParticipantFormData "Список всех участников"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Router /api/participants [get]
func (s *ServicesAPI) getAllParticipants(c *gin.Context) {
	// Получаем фильтры из query параметров
	fio := c.Query("fio")
	category := c.Query("category")
	gender := c.Query("gender")
	birthday := c.Query("birthday")
	coach := c.Query("coach")

	// Получение всех участников
	participants, err := s.Services.ParticipantService.GetAllParticipants()
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Error retrieving participants",
			Message: err.Error(),
		})
		return
	}

	// Применяем фильтры
	filteredParticipants := s.filterParticipants(participants, fio, category, gender, birthday, coach)

	// Конвертируем данные участников в форму API
	participantFormData, err := modelsViewApi.FromParticipantModelsToStringData(filteredParticipants)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	// Возвращаем ответ
	c.JSON(http.StatusOK, participantFormData)
}

// Функция для фильтрации участников
func (s *ServicesAPI) filterParticipants(participants []models.Participant, fio, category, gender, birthday, coach string) []models.Participant {
	var filteredParticipants []models.Participant
	for _, participant := range participants {
		if s.matchParticipantFilters(participant, fio, category, gender, birthday, coach) {
			filteredParticipants = append(filteredParticipants, participant)
		}
	}
	return filteredParticipants
}

func (s *ServicesAPI) matchParticipantFilters(participant models.Participant, fio, category, gender, birthday, coach string) bool {
	if !s.matchFio(participant.FIO, fio) {
		return false
	}
	if !s.matchCategory(participant.Category, category) {
		return false
	}
	if !s.matchGender(participant.Gender, gender) {
		return false
	}
	if !s.matchBirthday(participant.Birthday, birthday) {
		return false
	}
	if !s.matchCoach(participant.Coach, coach) {
		return false
	}
	return true
}

func (s *ServicesAPI) matchFio(fio, filterFio string) bool {
	return filterFio == "" || strings.Contains(strings.ToLower(fio), strings.ToLower(filterFio))
}

func (s *ServicesAPI) matchCategory(category int, filterCategory string) bool {
	return filterCategory == "" || modelsViewApi.CategoryMap[category] == filterCategory
}

func (s *ServicesAPI) matchGender(gender int, filterGender string) bool {
	return filterGender == "" || modelsViewApi.GenderMap[gender] == filterGender
}

func (s *ServicesAPI) matchBirthday(birthday time.Time, filterBirthday string) bool {
	return filterBirthday == "" || birthday.Format("2006-01-02") == filterBirthday
}

func (s *ServicesAPI) matchCoach(coach, filterCoach string) bool {
	return filterCoach == "" || strings.Contains(strings.ToLower(coach), strings.ToLower(filterCoach))
}

// @Summary Создать нового участника
// @Tags Participant
// @Produce json
// @Param body body modelsViewApi.ParticipantInput true "Данные для создания участника"
// @Success 201 {object} modelsViewApi.ParticipantFormData "Участник успешно создан"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Router /api/participants [post]
func (s *ServicesAPI) createParticipant(c *gin.Context) {
	var input modelsViewApi.ParticipantInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid input data",
			Message: "Failed to parse request body.",
		})
		return
	}

	createdParticipant, err := s.Services.ParticipantService.AddNewParticipant(uuid.New(), input.FIO, input.Category, input.Gender, apiUtils.ConvertStringToTime(input.Birthday), input.Coach)

	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Error creating participant",
			Message: err.Error(),
		})
		return
	}

	participantFormData, err := modelsViewApi.FromParticipantModelToStringData(createdParticipant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, participantFormData)
}

// @Summary Получить информацию о конкретном участнике
// @Tags Participant
// @Produce json
// @Param participantID path string true "Уникальный идентификатор участника" format(uuid)
// @Success 200 {object} modelsViewApi.ParticipantFormData "Информация о конкретном участнике"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Участник не найден"
// @Router /api/participants/{participantID} [get]
func (s *ServicesAPI) getParticipantById(c *gin.Context) {
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
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Participant not found",
				Message: "The specified participant ID does not exist.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	participantFormData, err := modelsViewApi.FromParticipantModelToStringData(participant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, participantFormData)
}

// @Summary Удалить участника
// @Tags Participant
// @Param participantID path string true "Уникальный идентификатор участника" format(uuid)
// @Success 204 "Участник успешно удален"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Участник не найден"
// @Router /api/participants/{participantID} [delete]
func (s *ServicesAPI) deleteParticipant(c *gin.Context) {
	participantIDParam := c.Param("participantID")

	participantID, err := uuid.Parse(participantIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid participant ID",
			Message: "The provided participant ID is not a valid UUID.",
		})
		return
	}

	err = s.Services.ParticipantService.DeleteParticipantByID(participantID)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Participant not found",
				Message: "The specified participant ID does not exist.",
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

// @Summary Обновить информацию об участнике
// @Tags Participant
// @Param participantID path string true "Уникальный идентификатор участника" format(uuid)
// @Param body body modelsViewApi.ParticipantInput true "Данные для обновления участника"
// @Success 200 {object} modelsViewApi.ParticipantFormData "Информация об участнике успешно обновлена"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Участник не найден"
// @Router /api/participants/{participantID} [put]
func (s *ServicesAPI) updateParticipant(c *gin.Context) {
	participantIDParam := c.Param("participantID")

	participantID, err := uuid.Parse(participantIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid participant ID",
			Message: "The provided participant ID is not a valid UUID.",
		})
		return
	}

	var input modelsViewApi.ParticipantInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid input data",
			Message: "Failed to parse request body.",
		})
		return
	}

	// Обновляем информацию об участнике
	participant, err := s.Services.ParticipantService.UpdateParticipantByID(participantID, input.FIO, input.Category, apiUtils.ConvertStringToTime(input.Birthday), input.Coach)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Participant not found",
				Message: "The specified participant ID does not exist.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	// Преобразуем модель участника в формат для ответа
	participantFormData, err := modelsViewApi.FromParticipantModelToStringData(participant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, participantFormData)
}
