package controllersApi

import (
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/server/api/modelsViewApi"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// getAllRatings godoc
// @Summary Получить список всех рейтингов
// @Description Получает список всех рейтингов с возможностью фильтрации
// @Tags Rating
// @Produce json
// @Param name query string false "Фильтр по названию рейтинга"
// @Param class query string false "Фильтр по классу лодки"
// @Param blowoutCnt query int false "Фильтр по количеству гонок, не учитываемых в результате"
// @Success 200 {array} modelsViewApi.RatingFormData "Список рейтингов"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Router /api/ratings [get]
func (s *ServicesAPI) getAllRatings(c *gin.Context) {
	// Получаем фильтры из query параметров
	name := c.Query("name")
	class := c.Query("class")
	blowoutCnt := c.Query("blowoutCnt")

	// Получение моделей уровня сервиса
	ratingModels, err := s.Services.RatingService.GetAllRatings()
	if err != nil {
		s.handleInternalError(c, "Не удалось получить рейтинги.")
		return
	}

	// Конвертация моделей уровня сервиса в модели API
	ratings, err := modelsViewApi.FromRatingModelsToStringData(ratingModels)
	if err != nil {
		s.handleConversionError(c, err)
		return
	}

	// Применение фильтров
	filteredRatings := s.filterRatings(ratings, name, class, blowoutCnt)

	c.JSON(http.StatusOK, filteredRatings)
}

// Функция для обработки ошибок конвертации
func (s *ServicesAPI) handleConversionError(c *gin.Context, err error) {
	if err == repository_errors.DoesNotExist {
		c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
			Error:   "Rating not found",
			Message: "The specified rating ID does not exist.",
		})
		return
	}
	c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
		Error:   "Internal error",
		Message: err.Error(),
	})
}

// Функция для фильтрации рейтингов
func (s *ServicesAPI) filterRatings(ratings []modelsViewApi.RatingFormData, name, class, blowoutCnt string) []modelsViewApi.RatingFormData {
	var filteredRatings []modelsViewApi.RatingFormData
	for _, rating := range ratings {
		if (name == "" || rating.Name == name) &&
			(class == "" || rating.Class == class) &&
			(blowoutCnt == "" || fmt.Sprintf("%d", rating.BlowoutCnt) == blowoutCnt) {
			filteredRatings = append(filteredRatings, rating)
		}
	}
	return filteredRatings
}

// Функция для обработки внутренних ошибок
func (s *ServicesAPI) handleInternalError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error":   "Internal Server Error",
		"message": message,
	})
}

// createRating godoc
// @Summary Создать новый рейтинг
// @Tags Rating
// @Accept json
// @Produce json
// @Param input body modelsViewApi.RatingInput true "Входные данные для создания рейтинга"
// @Success 201 {object} modelsViewApi.RatingInput "Рейтинг успешно создан"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Router /api/ratings [post]
func (s *ServicesAPI) createRating(c *gin.Context) {
	var input modelsViewApi.RatingInput

	// Попытка привязки входных данных к структуре RatingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid input",
			Message: "The provided input data is invalid.",
		})
		return
	}

	// Логика создания нового рейтинга через сервисный уровень
	rating, err := s.Services.RatingService.AddNewRating(uuid.New(), input.Name, input.Class, input.BlowoutCnt)
	if err != nil {
		// Если возникла ошибка на уровне сервиса, обрабатываем ее
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	// Возвращаем успешный результат
	c.JSON(http.StatusCreated, rating)
}

// getRating godoc
// @Summary Получить информацию о рейтинге
// @Tags Rating
// @Accept json
// @Produce json
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Success 200 {object} modelsViewApi.RatingFormData "Информация о рейтинге успешно получена"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Рейтинг не найден"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Router /api/ratings/{ratingID} [get]
func (s *ServicesAPI) getRating(c *gin.Context) {
	ratingIDParam := c.Param("ratingID")

	ratingID, err := uuid.Parse(ratingIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid rating ID",
			Message: "The provided rating ID is not a valid UUID.",
		})
		return
	}

	rating, err := s.Services.RatingService.GetRatingDataByID(ratingID)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Rating not found",
				Message: "The rating with the provided ID does not exist.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	class, exists := modelsViewApi.ClassMap[rating.Class]
	if !exists {
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Invalid class",
			Message: fmt.Sprintf("The class '%d' is not a valid class in modelsView ClassMap.", rating.Class),
		})
		return
	}

	c.JSON(http.StatusOK, modelsViewApi.RatingFormData{
		ID:         rating.ID,
		Name:       rating.Name,
		Class:      class,
		BlowoutCnt: rating.BlowoutCnt,
	})
}

// updateRating godoc
// @Summary Обновить информацию о рейтинге
// @Tags Rating
// @Accept json
// @Produce json
// @Param ratingID path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param input body modelsViewApi.RatingInput true "Входные данные для обновления рейтинга"
// @Success 200 {object} modelsViewApi.RatingFormData "Рейтинг успешно обновлен"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 500 {object} modelsViewApi.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/ratings/{ratingID} [put]
func (s *ServicesAPI) updateRating(c *gin.Context) {
	var input modelsViewApi.RatingInput
	ratingIDParam := c.Param("ratingID")

	ratingID, err := uuid.Parse(ratingIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid rating ID",
			Message: "The provided rating ID is not a valid UUID.",
		})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid input",
			Message: "The provided input data is invalid.",
		})
		return
	}

	updatedRating, err := s.Services.RatingService.UpdateRatingByID(ratingID, input.Name, input.Class, input.BlowoutCnt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modelsViewApi.BadRequestError{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	class, exists := modelsViewApi.ClassMap[updatedRating.Class]
	if !exists {
		c.JSON(http.StatusNotFound, modelsViewApi.BadRequestError{
			Error:   "Invalid class",
			Message: fmt.Sprintf("The class '%d' is not a valid class in modelsView ClassMap.", input.Class),
		})
		return
	}

	c.JSON(http.StatusOK, modelsViewApi.RatingFormData{
		ID:         updatedRating.ID,
		Name:       updatedRating.Name,
		Class:      class,
		BlowoutCnt: updatedRating.BlowoutCnt,
	})
}

// deleteRating godoc
// @Summary Удалить рейтинг
// @Tags Rating
// @Param ratingID path string true "Уникальный идентификатор рейтинга, который нужно удалить" format(uuid)
// @Success 204 "Рейтинг успешно удалён"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Рейтинг не найден"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Router /api/ratings/{ratingID} [delete]
func (s *ServicesAPI) deleteRating(c *gin.Context) {
	ratingIDParam := c.Param("ratingID")

	ratingID, err := uuid.Parse(ratingIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid rating ID",
			Message: "The provided rating ID is not a valid UUID.",
		})
		return
	}

	err = s.Services.RatingService.DeleteRatingByID(ratingID)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Rating not found",
				Message: "The rating with the provided ID does not exist.",
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

// @Summary Получить рейтинговую таблицу
// @Tags Ranking
// @Param ratingID path string true "ID рейтинга" format(uuid)
// @Success 200 {object} modelsViewApi.RankingResponse "Рейтинговая таблица успешно получена"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Рейтинг не найден"
// @Router /api/ratings/{ratingID}/rankings [get]
func (s *ServicesAPI) getRankingTable(c *gin.Context) {
	ratingIDParam := c.Param("ratingID")

	ratingID, err := uuid.Parse(ratingIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid rating ID",
			Message: "The provided rating ID is not a valid UUID.",
		})
		return
	}

	rankingTable, err := s.Services.RatingService.GetRatingTable(ratingID)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Rating not found",
				Message: "The specified rating ID does not exist.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	races, err := s.Services.RaceService.GetRacesDataByRatingID(ratingID)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Race not found",
				Message: "The specified race in rating does not exist.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	resRankingTable := modelsViewApi.FromRatingTableLinesModelTiStringData(rankingTable)

	var resRaces []modelsViewApi.RaceInfo
	for _, race := range races {
		resRaces = append(resRaces, modelsViewApi.RaceInfo{RaceNum: race.Number, RaceID: race.ID})
	}

	response := modelsViewApi.RankingResponse{
		RankingTable: resRankingTable,
		Races:        resRaces,
	}

	c.JSON(http.StatusOK, response)
}
