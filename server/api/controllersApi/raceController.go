package controllersApi

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/server/api/apiUtils"
	"PPO_BMSTU/server/api/modelsViewApi"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

// getRacesByRatingID godoc
// @Summary Получить гонки по ID рейтинга
// @Tags Race
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Success 200 {array} modelsViewApi.RaceFormData "Список гонок успешно получен"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Рейтинг или гонки не найдены"
// @Failure 500 {object} modelsViewApi.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/ratings/{ratingID}/races [get]
func (s *ServicesAPI) getRacesByRatingID(c *gin.Context) {
	// Получаем ratingID из параметров запроса
	ratingID, err := s.parseUUIDParam(c, "ratingID")
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid rating ID",
			Message: "The provided rating ID is not a valid UUID.",
		})
		return
	}

	// Получаем данные о гонках
	races, err := s.Services.RaceService.GetRacesDataByRatingID(ratingID)
	if err != nil {
		s.handleRaceServiceError(c, err)
		return
	}

	// Получаем фильтры из query параметров
	date := c.Query("date")
	class := c.Query("class")
	number := c.Query("number")

	// Применяем фильтры
	filteredRaces := s.filterRaces(races, date, class, number)

	if len(filteredRaces) == 0 {
		c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
			Error:   "No races found",
			Message: "There are no races associated with the specified rating.",
		})
		return
	}

	c.JSON(http.StatusOK, filteredRaces)
}

// Функция для парсинга UUID из параметра
func (s *ServicesAPI) parseUUIDParam(c *gin.Context, paramName string) (uuid.UUID, error) {
	param := c.Param(paramName)
	return uuid.Parse(param)
}

// Функция для обработки ошибок от RaceService
func (s *ServicesAPI) handleRaceServiceError(c *gin.Context, err error) {
	if err == repository_errors.DoesNotExist {
		c.JSON(http.StatusNotFound, modelsViewApi.BadRequestError{
			Error:   "Race or rating not found",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
		Error:   "Internal server error",
		Message: "Rating or races not found.",
	})
}

// Функция для фильтрации гонок
func (s *ServicesAPI) filterRaces(races []models.Race, date, class, number string) []*modelsViewApi.RaceFormData {
	var filteredRaces []*modelsViewApi.RaceFormData
	for _, race := range races {
		if (date == "" || race.Date.Format("2006-01-02") == date) &&
			(class == "" || modelsViewApi.ClassMap[race.Class] == class) &&
			(number == "" || strconv.Itoa(race.Number) == number) {
			raceStr, _ := modelsViewApi.FromRaceModelToStringData(&race)
			filteredRaces = append(filteredRaces, &raceStr)
		}
	}
	return filteredRaces
}

// createRace godoc
// @Summary Создать новую гонку для рейтинга
// @Tags Race
// @Accept json
// @Produce json
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param input body modelsViewApi.RaceInput true "Входные данные для создания гонки"
// @Success 201 {object} modelsViewApi.RaceInput "Гонка успешно создана"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Рейтинг не найден"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 500 {object} modelsViewApi.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/ratings/{ratingID}/races [post]
func (s *ServicesAPI) createRace(c *gin.Context) {
	var input modelsViewApi.RaceInput
	ratingIDParam := c.Param("ratingID")

	// Попытка привязки входных данных к структуре RaceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid input",
			Message: "The provided input data is invalid.",
		})
		return
	}

	// Преобразование ratingID из строки в uuid
	ratingID, err := uuid.Parse(ratingIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid rating ID",
			Message: "The provided rating ID is not a valid UUID.",
		})
		return
	}

	date := apiUtils.ConvertStringToTime(input.Date)

	race, err := s.Services.RaceService.AddNewRace(uuid.New(), ratingID, input.Number, date, input.Class)
	if err != nil {
		{
			c.JSON(http.StatusInternalServerError, modelsViewApi.BadRequestError{
				Error:   "Internal error",
				Message: err.Error(),
			})
		}
		return
	}

	// Возвращаем успешный результат
	c.JSON(http.StatusCreated, modelsViewApi.RaceInput{
		Date:   race.Date.String(),
		Number: race.Number,
		Class:  race.Class,
	})
}

// getRaceByID godoc
// @Summary Получить информацию о гонке
// @Tags Race
// @Produce json
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param raceID path string true "Уникальный идентификатор гонки" format(uuid)
// @Success 200 {object} modelsViewApi.RaceFormData "Информация о гонке успешно получена"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Гонка не найдена"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 500 {object} modelsViewApi.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/ratings/{ratingID}/races/{raceID} [get]
func (s *ServicesAPI) getRaceByID(c *gin.Context) {
	raceIDParam := c.Param("raceID")

	// Преобразование raceID из строки в uuid
	raceID, err := uuid.Parse(raceIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid race ID",
			Message: "The provided race ID is not a valid UUID.",
		})
		return
	}

	// Получение информации о гонке через сервисный уровень
	race, err := s.Services.RaceService.GetRaceDataByID(raceID)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.BadRequestError{
				Error:   "Race or rating not found",
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.BadRequestError{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	class, exists := modelsViewApi.ClassMap[race.Class]
	if !exists {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid class",
			Message: fmt.Sprintf("The class '%d' is not a valid class in modelsView ClassMap.", race.Class),
		})
		return
	}

	c.JSON(http.StatusOK, modelsViewApi.RaceFormData{
		ID:     race.ID,
		Date:   race.Date.String(),
		Number: race.Number,
		Class:  class,
	})
}

// updateRace godoc
// @Summary Обновить информацию о гонке
// @Tags Race
// @Accept json
// @Produce json
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param raceID path string true "Уникальный идентификатор гонки" format(uuid)
// @Param input body modelsViewApi.RaceInput true "Входные данные для обновления гонки"
// @Success 200 {object} modelsViewApi.RaceFormData "Гонка успешно обновлена"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Гонка не найдена"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 500 {object} modelsViewApi.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/ratings/{ratingID}/races/{raceID} [put]
func (s *ServicesAPI) updateRace(c *gin.Context) {
	ratingIDParam := c.Param("ratingID")
	raceIDParam := c.Param("raceID")

	raceID, err := uuid.Parse(raceIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid race ID",
			Message: "The provided race ID is not a valid UUID.",
		})
		return
	}

	ratingID, err := uuid.Parse(ratingIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid rating ID",
			Message: "The provided rating ID is not a valid UUID.",
		})
		return
	}

	var input modelsViewApi.RaceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid input",
			Message: "The provided input data is invalid.",
		})
		return
	}
	date := apiUtils.ConvertStringToTime(input.Date)

	updatedRace, err := s.Services.RaceService.UpdateRaceByID(raceID, ratingID, input.Number, date, input.Class)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.BadRequestError{
				Error:   "Race or rating not found",
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.BadRequestError{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	class, exists := modelsViewApi.ClassMap[updatedRace.Class]
	if !exists {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid class",
			Message: fmt.Sprintf("The class '%d' is not a valid class in modelsView ClassMap.", updatedRace.Class),
		})
		return
	}

	c.JSON(http.StatusOK, modelsViewApi.RaceFormData{
		ID:     updatedRace.ID,
		Date:   updatedRace.Date.String(),
		Number: updatedRace.Number,
		Class:  class,
	})
}

// deleteRace godoc
// @Summary Удалить гонку
// @Tags Race
// @Produce json
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param raceID path string true "Уникальный идентификатор гонки, которую нужно удалить" format(uuid)
// @Success 204 "Гонка успешно удалена"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Гонка не найдена"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 500 {object} modelsViewApi.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/ratings/{ratingID}/races/{raceID} [delete]
func (s *ServicesAPI) deleteRace(c *gin.Context) {
	raceIDParam := c.Param("raceID")

	// Преобразование raceID из строки в uuid
	raceID, err := uuid.Parse(raceIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid race ID",
			Message: "The provided race ID is not a valid UUID.",
		})
		return
	}

	err = s.Services.RaceService.DeleteRaceByID(raceID)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.BadRequestError{
				Error:   "Race or rating not found",
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, modelsViewApi.BadRequestError{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// startProcedure godoc
// @Summary Завершить стартовую процедуру гонки
// @Tags Race
// @Produce json
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param raceID path string true "Уникальный идентификатор гонки" format(uuid)
// @Param input body modelsViewApi.StartInput true "Входные данные для завершения стартовой процедуры"
// @Success 200 {string} string "Процедура старта успешно выполнена"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Гонка не найдена"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 500 {object} modelsViewApi.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/ratings/{ratingID}/races/{raceID}/start [post]
func (s *ServicesAPI) startProcedure(c *gin.Context) {
	raceIDParam := c.Param("raceID")

	raceID, err := uuid.Parse(raceIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid race ID",
			Message: "The provided race ID is not a valid UUID.",
		})
		return
	}

	var input modelsViewApi.StartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid request body",
			Message: "The request body must be a valid JSON object.",
		})
		return
	}

	falseStartYachtMap := modelsViewApi.FromStartInputViewToStartInput(input.FalseStartList, input.SpecCircumstance)
	err = s.Services.RaceService.MakeStartProcedure(raceID, falseStartYachtMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modelsViewApi.BadRequestError{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "Процедура старта успешно выполнена")
}

// finishProcedure godoc
// @Summary Завершить финишную процедуру гонки
// @Tags Race
// @Produce json
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param raceID path string true "Уникальный идентификатор гонки" format(uuid)
// @Param input body modelsViewApi.FinishInput true "Входные данные для завершения финишной процедуры"
// @Success 200 {string} string "Процедура финиша успешно выполнена"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Гонка не найдена"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 500 {object} modelsViewApi.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/ratings/{ratingID}/races/{raceID}/finish [post]
func (s *ServicesAPI) finishProcedure(c *gin.Context) {
	ratingIDParam := c.Param("ratingID")
	raceIDParam := c.Param("raceID")

	ratingID, err := uuid.Parse(ratingIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid race ID",
			Message: "The provided race ID is not a valid UUID.",
		})
		return
	}

	raceID, err := uuid.Parse(raceIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid race ID",
			Message: "The provided race ID is not a valid UUID.",
		})
		return
	}

	var input modelsViewApi.FinishInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid request body",
			Message: "The request body must be a valid JSON object.",
		})
		return
	}

	crews, err := s.Services.CrewService.GetCrewsDataByRatingID(ratingID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modelsViewApi.BadRequestError{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	finisherMap, nonFinisherMap := modelsViewApi.FromFinishInputViewToFinishInput(input.FinisherList, crews)

	err = s.Services.RaceService.MakeFinishProcedure(raceID, finisherMap, nonFinisherMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modelsViewApi.BadRequestError{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "Процедура финиша успешно выполнена")
}
