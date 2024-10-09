package controllersApi

import (
	"PPO_BMSTU/server/api/modelsViewApi"
	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Could not retrieve judges.",
		})
		return
	}

	// Конвертация из модели уровня сервиса в модель апи
	judges, err := modelsViewApi.FromJudgeModelsToStringData(judgeModels)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Could not convert judge models.",
		})
		return
	}

	// Применение фильтров
	var filteredJudges []modelsViewApi.JudgeFormData
	for _, judge := range judges {
		if (fio == "" || judge.FIO == fio) &&
			(login == "" || judge.Login == login) &&
			(role == "" || judge.Role == role) &&
			(post == "" || judge.Post == post) {
			filteredJudges = append(filteredJudges, judge)
		}
	}

	c.JSON(http.StatusOK, filteredJudges)
}
