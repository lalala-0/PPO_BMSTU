package controllersApi

import (
	"PPO_BMSTU/server/api/modelsViewApi"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// getCrewsByRatingID godoc
// @Summary Получить все команды в рейтинге
// @Tags Crew
// @Produce json
// @Param ratingId path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Success 200 {array} modelsViewApi.CrewFormData "Список команд успешно получен"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Не найдены рейтинг или команды"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 500 {object} modelsViewApi.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/ratings/{ratingId}/crews [get]
func (s *ServicesAPI) getCrewsByRatingID(c *gin.Context) {
	ratingIDParam := c.Param("ratingId")

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
		if err == sql.ErrNoRows {
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

	c.JSON(http.StatusOK, crews)
}

// createCrew godoc
// @Summary Создать новую команду
// @Tags Crew
// @Produce json
// @Param ratingId path string true "Уникальный идентификатор рейтинга" format(uuid)
// @Param input body modelsViewApi.CrewInput true "Данные для создания новой команды"
// @Success 201 {object} modelsViewApi.CrewFormData "Команда успешно создана"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Рейтинг не найден"
// @Router /api/ratings/{ratingId}/crews [post]
func (s *ServicesAPI) createCrew(c *gin.Context) {
	ratingIDParam := c.Param("ratingId")

	// Преобразование ratingID из строки в uuid
	ratingID, err := uuid.Parse(ratingIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid rating ID",
			Message: "The provided rating ID is not a valid UUID.",
		})
		return
	}

	var input modelsViewApi.CrewInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid input data",
			Message: "The provided data for creating the crew is invalid.",
		})
		return
	}

	crew, err := s.Services.CrewService.AddNewCrew(uuid.New(), ratingID, rating.Class, input.SailNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modelsViewApi.BadRequestError{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	// Успешный ответ
	c.JSON(http.StatusCreated, modelsViewApi.CrewFormData{
		ID:       crew.ID,
		RatingID: crew.RatingID,
		SailNum:  crew.SailNum,
		Class:    crew.Class,
	})
}
