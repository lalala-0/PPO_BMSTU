package controllersUi

import (
	"PPO_BMSTU/internal/models"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (s *ServicesUI) authenticatedJudge(c *gin.Context) *models.Judge {
	session := sessions.Default(c)
	sessionID := session.Get("judgeID")

	if sessionID != nil {
		strJudgeID, ok := sessionID.(string)
		if ok {
			judgeId, err := uuid.Parse(strJudgeID)
			if err == nil {
				judge, err := s.Services.JudgeService.GetJudgeDataByID(judgeId)
				if err == nil {
					return judge
				}
			}
		}
	}
	return nil
}

type loginFormData struct {
	Login    string `form:"InputLogin"`
	Password string `form:"InputPassword"`
}

func (s *ServicesUI) signinGet(c *gin.Context) {
	c.HTML(200, "signin", gin.H{
		"title": "Вход",
	})
}

func (s *ServicesUI) signinPost(c *gin.Context) {
	var data loginFormData
	if err := c.Bind(&data); err != nil {
		c.HTML(http.StatusBadRequest, "signin", gin.H{
			"title": "Вход",
			"error": err.Error(),
		})
		return
	}

	judge, err := s.Services.JudgeService.Login(data.Login, data.Password)
	if err != nil {
		c.HTML(http.StatusBadRequest, "signin", gin.H{
			"title":    "Вход",
			"error":    "Неверный пароль или имя пользователя",
			"formData": data,
		})
		return
	}

	// Set the session.
	session := sessions.Default(c)
	session.Set("judgeID", judge.ID.String())
	ok := session.Save()
	if ok != nil {
		c.HTML(http.StatusBadRequest, "signin", gin.H{
			"title":    "Вход",
			"error":    "Не удалось сохранить сессию",
			"formData": data,
		})
		return
	}

	c.Redirect(http.StatusFound, "/")
}

func (s *ServicesUI) logout(c *gin.Context) {
	// Delete the session
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		log.Error("Error save session while logout: %s", err)
	}
	c.Status(http.StatusAccepted)
	c.Redirect(http.StatusFound, "/")
}
