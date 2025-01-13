package controllersApi

import (
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/server/api/modelsViewApi"
	modelsUI2 "PPO_BMSTU/server/ui/modelsUI"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// Получить список всех судей
// @Summary Получить список всех судей
// @Description Получает список всех судей с возможностью фильтрации
// @Tags Judge
// @Produce json
// @Param fio query string false "Фильтр по ФИО"
// @Param login query string false "Фильтр по логину"
// @Param role query string false "Фильтр по роли"
// @Param post query string false "Фильтр по должности"
// @Success 200 {array} modelsViewApi.JudgeFormData
// @Failure 400 {object} modelsViewApi.BadRequestError
// @Router /api/judges [get]
func (s *ServicesAPI) getAllJudges(c *gin.Context) {
	// Получаем фильтры из query параметров
	fio := c.Query("fio")
	login := c.Query("login")
	role := c.Query("role")
	post := c.Query("post")

	// Получение моделей уровня сервиса
	judgeModels, err := s.Services.JudgeService.GetAllJudges()
	if err != nil {
		s.handleInternalError(c, "Не удалось получить судей.")
		return
	}

	// Конвертация из модели уровня сервиса в модель API
	judges, err := modelsViewApi.FromJudgeModelsToStringData(judgeModels)
	if err != nil {
		s.handleConversionError(c, err)
		return
	}

	// Применение фильтров
	filteredJudges := s.filterJudges(judges, fio, login, role, post)

	c.JSON(http.StatusOK, filteredJudges)
}

// Функция для фильтрации судей
func (s *ServicesAPI) filterJudges(judges []modelsViewApi.JudgeFormData, fio, login, role, post string) []modelsViewApi.JudgeFormData {
	var filteredJudges []modelsViewApi.JudgeFormData
	for _, judge := range judges {
		if (fio == "" || judge.FIO == fio) &&
			(login == "" || judge.Login == login) &&
			(role == "" || judge.Role == role) &&
			(post == "" || judge.Post == post) {
			filteredJudges = append(filteredJudges, judge)
		}
	}
	return filteredJudges
}

// @Summary Создать нового судью
// @Tags Judge
// @Param body body modelsViewApi.JudgeInput true "Данные для создания судьи"
// @Success 201 {object} modelsViewApi.JudgeInput "Судья успешно создан"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Router /api/judges [post]
func (s *ServicesAPI) createJudge(c *gin.Context) {
	var input modelsViewApi.JudgeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid input data",
			Message: "Failed to parse request body.",
		})
		return
	}

	// Создание нового судьи
	judge, err := s.Services.JudgeService.CreateProfile(uuid.New(), input.FIO, input.Login, input.Password, input.Role, input.Post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, judge)
}

// @Summary Получить информацию о судье
// @Tags Judge
// @Param judgeID path string true "ID судьи" format(uuid)
// @Success 200 {object} modelsViewApi.JudgeFormData "Информация о судье"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Судья не найден"
// @Router /api/judges/{judgeID} [get]
func (s *ServicesAPI) getJudgeByID(c *gin.Context) {
	judgeIDParam := c.Param("judgeID")

	judgeID, err := uuid.Parse(judgeIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid judge ID",
			Message: "The provided judge ID is not a valid UUID.",
		})
		return
	}

	judge, err := s.Services.JudgeService.GetJudgeDataByID(judgeID)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Judge not found",
				Message: "The specified judge ID does not exist.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	judgeFormData, err := modelsViewApi.FromJudgeModelToStringData(judge)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, judgeFormData)
}

// @Summary Удалить судью
// @Tags Judge
// @Param judgeID path string true "ID судьи" format(uuid)
// @Success 204 {string} string "Судья успешно удален"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Судья не найден"
// @Router /api/judges/{judgeID} [delete]
func (s *ServicesAPI) deleteJudge(c *gin.Context) {
	judgeIDParam := c.Param("judgeID")

	judgeID, err := uuid.Parse(judgeIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid judge ID",
			Message: "The provided judge ID is not a valid UUID.",
		})
		return
	}

	err = s.Services.JudgeService.DeleteProfile(judgeID)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Judge not found",
				Message: "The specified judge ID does not exist.",
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

// @Summary Обновить информацию о судье
// @Tags Judge
// @Param judgeID path string true "ID судьи" format(uuid)
// @Param body body modelsViewApi.JudgeInput true "Данные для обновления судьи"
// @Success 200 {object} modelsViewApi.JudgeFormData "Информация о судье успешно обновлена"
// @Failure 400 {object} modelsViewApi.BadRequestError "Ошибка валидации"
// @Failure 404 {object} modelsViewApi.ErrorResponse "Судья не найден"
// @Router /api/judges/{judgeID} [put]
func (s *ServicesAPI) updateJudge(c *gin.Context) {
	judgeIDParam := c.Param("judgeID")

	judgeID, err := uuid.Parse(judgeIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid judge ID",
			Message: "The provided judge ID is not a valid UUID.",
		})
		return
	}

	var input modelsViewApi.JudgeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, modelsViewApi.BadRequestError{
			Error:   "Invalid input data",
			Message: "Failed to parse request body.",
		})
		return
	}

	updatedJudge, err := s.Services.JudgeService.UpdateProfile(judgeID, input.FIO, input.Login, input.Password, input.Role)
	if err != nil {
		if err == repository_errors.DoesNotExist {
			c.JSON(http.StatusNotFound, modelsViewApi.ErrorResponse{
				Error:   "Judge not found",
				Message: "The specified judge ID does not exist.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	judgeFormData, err := modelsViewApi.FromJudgeModelToStringData(updatedJudge)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modelsViewApi.ErrorResponse{
			Error:   "Internal error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, judgeFormData)
}

func (s *ServicesAPI) updatePassword(c *gin.Context) {
	var input modelsUI2.PasswordInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data",
		})
		return
	}

	judgeID, err := uuid.Parse(c.Param("judgeID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid judge ID",
		})
		return
	}

	judge, err := s.Services.JudgeService.GetJudgeDataByID(judgeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Judge not found",
		})
		return
	}

	_, err = s.Services.JudgeService.UpdateProfile(judgeID, judge.FIO, judge.Login, input.Password, judge.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password updated successfully",
	})
}

func (s *ServicesAPI) signin(c *gin.Context) {
	var data modelsViewApi.LoginFormData
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data",
		})
		return
	}

	_, err := s.Services.JudgeService.Login(data.Login, data.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Authentication failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Authentication successful",
	})
}
