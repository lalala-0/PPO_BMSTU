package middleware

import (
	"PPO_BMSTU/internal/registry"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	Services *registry.Services
}

func NewMiddleware(registry registry.App) *Middleware {
	return &Middleware{Services: registry.Services}
}

func (m *Middleware) JudgeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("judgeID")
		if sessionID == nil {
			c.Redirect(http.StatusMovedPermanently, "/auth/signin")
			c.Abort()
			return
		}

		strJudgeId, ok := sessionID.(string)
		if !ok {
			c.Redirect(http.StatusMovedPermanently, "/auth/signin")
			c.Abort()
			return
		}
		// Check if the user exists
		judgeID, err := uuid.Parse(strJudgeId)
		judge, err := m.Services.JudgeService.GetJudgeDataByID(judgeID)
		if err != nil || judge.ID == uuid.Nil {
			c.Redirect(http.StatusMovedPermanently, "/auth/signin")
			c.Abort()
			return
		}

		c.Set("judgeID", judge.ID)
		c.Next()
	}
}
