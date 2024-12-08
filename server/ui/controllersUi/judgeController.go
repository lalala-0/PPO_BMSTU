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

func (s *ServicesUI) judgeProfile(c *gin.Context) {
	judge := s.authenticatedJudge(c)
	judgeView, _ := modelsUI2.FromJudgeModelToStringData(judge)

	c.HTML(200, "judgeProfile", gin.H{
		"title":     "Профиль судьи",
		"judgeView": judgeView,
		"judge":     judge,
	})
}

func (s *ServicesUI) judgeMainMenu(judge *models.Judge) gin.H {

	participants, err := s.Services.ParticipantService.GetAllParticipants()
	if err != nil {
		log.Printf("Error getting races: %v", err)
		participants = []models.Participant{}
	}
	participantsView, _ := modelsUI2.FromParticipantModelsToStringData(participants)

	judgeView, _ := modelsUI2.FromJudgeModelToStringData(judge)

	var result gin.H

	if judge.Role == models.MainJudge {
		judges, err := s.Services.JudgeService.GetAllJudges()
		if err != nil {
			log.Printf("Error getting crews: %v", err)
			judges = []models.Judge{}
		}
		judgesView, _ := modelsUI2.FromJudgeModelsToStringData(judges)

		result = gin.H{
			"title":            "Панель судьи",
			"mainJudge":        true,
			"judgeView":        judgeView,
			"judge":            judge,
			"participants":     participants,
			"participantsView": participantsView,
			"judges":           judges,
			"judgesView":       judgesView,
		}
	} else {
		result = gin.H{
			"title":            "Панель судьи",
			"mainJudge":        false,
			"judgeView":        judgeView,
			"judge":            judge,
			"participants":     participants,
			"participantsView": participantsView,
		}
	}

	return result
}

func (s *ServicesUI) menu(c *gin.Context) {
	judge := s.authenticatedJudge(c)

	c.HTML(200, "judgeDashboard", s.judgeMainMenu(judge))
}

// UPDATE PASSWORD
func (s *ServicesUI) updatePasswordGet(c *gin.Context) {
	judgeID, err := uuid.Parse(c.Param("judgeID"))
	if err != nil {
		c.HTML(400, "updatePassword", gin.H{
			"title": "Обновить пароль",
			"error": "Неверный идентификатор судьи",
		})
		return
	}

	judge, err := s.Services.JudgeService.GetJudgeDataByID(judgeID)
	if err != nil {
		c.HTML(400, "updatePassword", gin.H{
			"title": "Обновить пароль",
			"error": "Судья не найден",
		})
		return
	}

	c.HTML(200, "updatePassword", gin.H{
		"title":    "Обновить пароль",
		"judge":    judge,
		"formData": modelsUI2.PasswordInput{},
	})
}

func (s *ServicesUI) updatePasswordPost(c *gin.Context) {
	var input modelsUI2.PasswordInput
	err := c.Bind(&input)
	if err != nil {
		log.Printf("Error binding data: %v", err)
		c.HTML(400, "updatePassword", gin.H{
			"title":    "Обновить пароль",
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	judgeID, err := uuid.Parse(c.Param("judgeID"))
	if err != nil {
		c.HTML(400, "updatePassword", gin.H{
			"title": "Обновить пароль",
			"error": "Неверный идентификатор судьи",
		})
		return
	}
	judge, err := s.Services.JudgeService.GetJudgeDataByID(judgeID)
	if err != nil {
		c.HTML(400, "updatePassword", gin.H{
			"title": "Обновить пароль",
			"error": "Неверный идентификатор судьи",
		})
		return
	}

	_, err = s.Services.JudgeService.UpdateProfile(judgeID, judge.FIO, judge.Login, input.Password, judge.Role)
	if err != nil {
		log.Printf("Error updating judge: %v", err)
		c.HTML(400, "updatePassword", gin.H{
			"title":    "Обновить пароль",
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

// UPDATE
func (s *ServicesUI) updateJudgeGet(c *gin.Context) {
	judgeID, err := uuid.Parse(c.Param("judgeID"))
	if err != nil {
		c.HTML(400, "updateJudge", gin.H{
			"title": "Редактировать профиль",
			"error": "Неверный идентификатор профиль",
		})
		return
	}

	judge, err := s.Services.JudgeService.GetJudgeDataByID(judgeID)
	if err != nil {
		c.HTML(400, "updateJudge", gin.H{
			"title": "Редактировать профиль",
			"error": "Судья не найден",
		})
		return
	}

	editedJudgeView, _ := modelsUI2.FromJudgeModelToInputData(judge)

	c.HTML(200, "updateJudge", gin.H{
		"title":        "Редактировать профиль",
		"judge":        judge,
		"formData":     editedJudgeView,
		"judgeRoleMap": modelsUI2.JudgeRoleMap,
	})
}

func (s *ServicesUI) updateJudgePost(c *gin.Context) {
	var input modelsUI2.JudgeInput
	err := c.Bind(&input)
	if err != nil {
		log.Printf("Error binding data: %v", err)
		c.HTML(400, "updateJudge", gin.H{
			"title":    "Редактировать профиль",
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	judgeID, err := uuid.Parse(c.Param("judgeID"))
	if err != nil {
		c.HTML(400, "updateJudge", gin.H{
			"title": "Редактировать профиль",
			"error": "Неверный идентификатор профиль",
		})
		return
	}

	judge, err := s.Services.JudgeService.GetJudgeDataByID(judgeID)
	if err != nil {
		c.HTML(400, "updateJudge", gin.H{
			"title": "Редактировать профиль",
			"error": "Судья не найден",
		})
		return
	}

	_, err = s.Services.JudgeService.UpdateProfile(judgeID, input.FIO, input.Login, judge.Password, input.Role)
	if err != nil {
		log.Printf("Error updating judge: %v", err)
		c.HTML(400, "updateJudge", gin.H{
			"title":    "Редактировать профиль",
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

func (s *ServicesUI) getJudgeMenu(c *gin.Context) {
	idStr := c.Param("judgeID")
	judgeID, err := uuid.Parse(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат UUID")
		return
	}
	judgeData, err := s.Services.JudgeService.GetJudgeDataByID(judgeID)
	if judgeData == nil || err != nil {
		c.String(http.StatusNotFound, "Crew not found")
		return
	}

	judge := s.authenticatedJudge(c)

	c.HTML(http.StatusOK, "judge", s.judgeMenu(judgeData, judge))
}

func (s *ServicesUI) judgeMenu(judgeData *models.Judge, judge *models.Judge) gin.H {
	judgeView, _ := modelsUI2.FromJudgeModelToStringData(judgeData)

	result := gin.H{
		"title":     "",
		"judge":     judge,
		"judgeData": judgeData,
		"judgeView": judgeView,
	}

	return result
}

// DELETE
func (s *ServicesUI) deleteJudge(c *gin.Context) {
	judgeID, err := uuid.Parse(c.Param("judgeID"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error-page", gin.H{
			"error": err.Error(),
		})
		return
	}

	err = s.Services.JudgeService.DeleteProfile(judgeID)
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
		referer = u.String()
	}

	c.Redirect(http.StatusFound, referer)
}

// CREATE
func (s *ServicesUI) createJudgeGet(c *gin.Context) {

	c.HTML(200, "createJudge", gin.H{
		"title":        "Создать профиль",
		"formData":     modelsUI2.JudgeInput{},
		"judgeRoleMap": modelsUI2.JudgeRoleMap,
	})
}

func (s *ServicesUI) createJudgePost(c *gin.Context) {
	var input modelsUI2.JudgeInput
	err := c.Bind(&input)
	if err != nil {
		log.Printf("Error binding data: %v", err)
		c.HTML(400, "createJudge", gin.H{
			"title":    "Создать профиль",
			"error":    err.Error(),
			"formData": input,
		})
		return
	}

	_, err = s.Services.JudgeService.CreateProfile(uuid.New(), input.FIO, input.Login, input.Password, input.Role, input.Post)
	if err != nil {
		log.Printf("Error updating judge: %v", err)
		c.HTML(400, "createJudge", gin.H{
			"title":    "Создать профиль",
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
