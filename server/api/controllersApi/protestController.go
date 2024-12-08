package controllersApi

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/server/api/apiUtils"
	"PPO_BMSTU/server/api/modelsViewApi"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

// @Summary Получить все протесты для указанного рейтинга и гонки
// @Tags Protests
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param raceID path string true "Уникальный идентификатор гонки" format(uuid)
// @Success 200 {array} modelsViewApi.ProtestFormData "Список протестов успешно получен"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Рейтинг или гонка не найдены"
// @Failure 500 {object} modelsViewApi.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/ratings/{ratingID}/races/{raceID}/protests [get]
func (s *ServicesAPI) getProtests(c *gin.Context) {
	raceIDParam := c.Param("raceID")

	raceID, err := uuid.Parse(raceIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid race ID",
			Message: "The provided race ID is not a valid UUID.",
		})
		return
	}

	protests, err := s.Services.ProtestService.GetProtestsDataByRaceID(raceID)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Rating or race not found",
				Message: "The specified rating ID or race ID does not exist.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, protests)
}

// @Summary Создать новый протест
// @Tags Protests
// @OperationId createProtest
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param raceID path string true "Уникальный идентификатор гонки" format(uuid)
// @Param protestInput body modelsViewApi.ProtestCreate true "Данные для создания протеста"
// @Success 201 {object} modelsViewApi.ProtestFormData "Протест успешно создан"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Рейтинг или гонка не найдены"
// @Router /api/ratings/{ratingID}/races/{raceID}/protests [post]
func (s *ServicesAPI) createProtest(c *gin.Context) {
	ratingIDParam := c.Param("ratingID")
	raceIDParam := c.Param("raceID")

	// Проверка корректности формата ratingID
	ratingID, err := uuid.Parse(ratingIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid rating ID",
			Message: "The provided rating ID is not a valid UUID.",
		})
		return
	}

	// Проверка корректности формата raceID
	raceID, err := uuid.Parse(raceIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid race ID",
			Message: "The provided race ID is not a valid UUID.",
		})
		return
	}

	// Получение списка доступных участников из параметров запроса
	//var availableParticipants []int
	//if queryParticipants, exists := c.GetQueryArray("availableParticipants"); exists {
	//	for _, participant := range queryParticipants {
	//		num, err := strconv.Atoi(participant)
	//		if err != nil {
	//			c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
	//				Error:   "Invalid participant number",
	//				Message: "One or more participant numbers are not valid integers.",
	//			})
	//			return
	//		}
	//		availableParticipants = append(availableParticipants, num)
	//	}
	//}

	var input modelsViewApi.ProtestCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid input",
			Message: "The request body is not valid.",
		})
		return
	}

	protest, err := s.Services.ProtestService.AddNewProtest(raceID, ratingID, input.JudgeID, input.RuleNum, apiUtils.ConvertStringToTime(input.ReviewDate), input.Comment, input.ProtesteeSailNum, input.ProtestorSailNum, input.WitnessesSailNum)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Rating or race not found",
				Message: "The specified rating ID or race ID does not exist.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, protest)
}

// @Summary Получить информацию о протесте
// @Tags Protests
// @OperationId getProtest
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param raceID path string true "Уникальный идентификатор гонки" format(uuid)
// @Param protestID path string true "Уникальный идентификатор протеста" format(uuid)
// @Success 200 {object} modelsViewApi.ProtestFormData "Информация о протесте"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Протест, рейтинг или гонка не найдены"
// @Router /api/ratings/{ratingID}/races/{raceID}/protests/{protestID} [get]
func (s *ServicesAPI) getProtest(c *gin.Context) {
	protestIDParam := c.Param("protestID")

	protestID, err := uuid.Parse(protestIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid protest ID",
			Message: "The provided protest ID is not a valid UUID.",
		})
		return
	}

	protest, err := s.Services.ProtestService.GetProtestDataByID(protestID)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Protest not found",
				Message: "The specified protest ID does not exist.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, protest)
}

// @Summary Завершить рассмотрение протеста
// @Tags Protests
// @OperationId completeProtest
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param raceID path string true "Уникальный идентификатор гонки" format(uuid)
// @Param protestID path string true "Уникальный идентификатор протеста" format(uuid)
// @Param protestComplete body modelsViewApi.ProtestComplete true "Данные для завершения рассмотрения протеста"
// @Success 200 {object} modelsViewApi.ProtestFormData "Рассмотрение протеста успешно завершено"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Протест, рейтинг или гонка не найдены"
// @Router /api/ratings/{ratingID}/races/{raceID}/protests/{protestID}/complete [patch]
func (s *ServicesAPI) completeProtest(c *gin.Context) {
	protestIDParam := c.Param("protestID")

	protestID, err := uuid.Parse(protestIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid protest ID",
			Message: "The provided protest ID is not a valid UUID.",
		})
		return
	}

	var protestComplete modelsViewApi.ProtestComplete
	if err := c.ShouldBindJSON(&protestComplete); err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid input",
			Message: "The request body is not valid.",
		})
		return
	}

	// Завершение рассмотрения протеста
	err = s.Services.ProtestService.CompleteReview(protestID, protestComplete.ResPoints, protestComplete.Comment)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Protest not found",
				Message: "The specified protest ID does not exist.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	updatedProtest, err := s.Services.ProtestService.GetProtestDataByID(protestID)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Protest not found",
				Message: "The specified protest ID does not exist.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedProtest)
}

// @Summary Удалить протест
// @Tags Protests
// @OperationId deleteProtest
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param raceID path string true "Уникальный идентификатор гонки" format(uuid)
// @Param protestID path string true "Уникальный идентификатор протеста" format(uuid)
// @Success 204 "Протест успешно удален"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Протест, рейтинг или гонка не найдены"
// @Router /api/ratings/{ratingID}/races/{raceID}/protests/{protestID} [delete]
func (s *ServicesAPI) deleteProtest(c *gin.Context) {
	protestIDParam := c.Param("protestID")

	protestID, err := uuid.Parse(protestIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid protest ID",
			Message: "The provided protest ID is not a valid UUID.",
		})
		return
	}

	// Удаление протеста
	err = s.Services.ProtestService.DeleteProtestByID(protestID)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Protest not found",
				Message: "The specified protest ID does not exist.",
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

// @Summary Обновить информацию о протесте
// @Tags Protests
// @OperationId updateProtest
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param raceID path string true "Уникальный идентификатор гонки" format(uuid)
// @Param protestID path string true "Уникальный идентификатор протеста" format(uuid)
// @Param protestInput body modelsViewApi.ProtestInput true "Данные для обновления протеста"
// @Success 200 {object} modelsViewApi.ProtestFormData "Протест успешно обновлен"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Протест, рейтинг или гонка не найдены"
// @Router /api/ratings/{ratingID}/races/{raceID}/protests/{protestID} [put]
func (s *ServicesAPI) updateProtest(c *gin.Context) {
	raceIDParam := c.Param("raceID")
	protestIDParam := c.Param("protestID")

	// Проверка корректности формата raceID
	raceID, err := uuid.Parse(raceIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid race ID",
			Message: "The provided race ID is not a valid UUID.",
		})
		return
	}

	// Проверка корректности формата protestID
	protestID, err := uuid.Parse(protestIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid protest ID",
			Message: "The provided protest ID is not a valid UUID.",
		})
		return
	}

	// Чтение данных тела запроса
	var protestInput modelsViewApi.ProtestInput
	if err := c.ShouldBindJSON(&protestInput); err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid input",
			Message: "The request body is not valid.",
		})
		return
	}

	// Обновление протеста
	updatedProtest, err := s.Services.ProtestService.UpdateProtestByID(protestID, raceID, protestInput.JudgeID, protestInput.RuleNum, apiUtils.ConvertStringToTime(protestInput.ReviewDate), protestInput.Status, protestInput.Comment)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Protest not found",
				Message: "The specified protest ID does not exist.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedProtest)
}

// @Summary Добавить команду-участника протеста
// @Tags Protest Members
// @OperationId attachProtestMember
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param raceID path string true "Уникальный идентификатор гонки" format(uuid)
// @Param protestID path string true "Уникальный идентификатор протеста" format(uuid)
// @Param protestParticipantAttachInput body modelsViewApi.ProtestParticipantAttachInput true "Данные для добавления команды-участника"
// @Success 201 {object} map[string]string "Команда-участник успешно добавлена"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Протест, гонка или команда не найдены"
// @Router /api/ratings/{ratingID}/races/{raceID}/protests/{protestID}/members [post]
func (s *ServicesAPI) attachProtestMember(c *gin.Context) {
	protestIDParam := c.Param("protestID")

	// Проверка корректности формата protestID
	protestID, err := uuid.Parse(protestIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid protest ID",
			Message: "The provided protest ID is not a valid UUID.",
		})
		return
	}

	// Чтение данных тела запроса
	var protestParticipantAttachInput modelsViewApi.ProtestParticipantAttachInput
	if err := c.ShouldBindJSON(&protestParticipantAttachInput); err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid input",
			Message: "The request body is not valid.",
		})
		return
	}

	// Добавление команды-участника
	err = s.Services.ProtestService.AttachCrewToProtest(protestID, protestParticipantAttachInput.SailNum, protestParticipantAttachInput.Role)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Protest not found",
				Message: "The specified protest ID does not exist.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Participant attached successfully"})
}

// @Summary Получить информацию о всех командах-участниках протеста
// @Tags Protest Members
// @OperationId getProtestMembers
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param raceID path string true "Уникальный идентификатор гонки" format(uuid)
// @Param protestID path string true "Уникальный идентификатор протеста" format(uuid)
// @Success 200 {array} modelsViewApi.ProtestCrewFormData "Список команд-участников протеста"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Протест, гонка или команды не найдены"
// @Router /api/ratings/{ratingID}/races/{raceID}/protests/{protestID}/members [get]
func (s *ServicesAPI) getProtestMembers(c *gin.Context) {
	protestIDParam := c.Param("protestID")

	protestID, err := uuid.Parse(protestIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid protest ID",
			Message: "The provided protest ID is not a valid UUID.",
		})
		return
	}

	memberIDs, err := s.Services.ProtestService.GetProtestParticipantsIDByID(protestID)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Protest not found",
				Message: "The specified protest ID does not exist.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}
	var roles []int
	var members []models.Crew
	for memberID, role := range memberIDs {
		member, err := s.Services.CrewService.GetCrewDataByID(memberID)
		if err != nil {
			if err == repository_errors.DoesNotExist {
				c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
					Error:   "Protest participant not found",
					Message: "The specified protest participant does not exist.",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
				Error:   "Internal error",
				Message: err.Error(),
			})
			return
		}
		members = append(members, *member)
		roles = append(roles, role)

	}
	protestMembers, _ := modelsViewApi.FromProtestParticipantModelsToStringData(members, roles)
	c.JSON(http.StatusOK, protestMembers)
}

// @Summary Удалить команду-участника из протеста
// @Tags Protest Members
// @OperationId detachProtestMember
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param raceID path string true "Уникальный идентификатор гонки" format(uuid)
// @Param protestID path string true "Уникальный идентификатор протеста" format(uuid)
// @Param crewSailNum path integer true "Номер паруса команды" format(int)
// @Success 204 "Команда-участник успешно удалена из протеста"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Протест, гонка или команда не найдены"
// @Router /api/ratings/{ratingID}/races/{raceID}/protests/{protestID}/members/{crewSailNum} [delete]
func (s *ServicesAPI) detachProtestMember(c *gin.Context) {
	protestIDParam := c.Param("protestID")
	sailNum, err := strconv.Atoi(c.Param("crewSailNum"))
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid sail number",
			Message: "The provided sail number is not integer.",
		})
		return
	}

	protestID, err := uuid.Parse(protestIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid protest ID",
			Message: "The provided protest ID is not a valid UUID.",
		})
		return
	}

	// Удаление команды-участника
	err = s.Services.ProtestService.DetachCrewFromProtest(protestID, sailNum)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Protest member not found",
				Message: "The specified team is not part of the protest.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
