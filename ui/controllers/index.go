package controllers

import (
	"github.com/gin-gonic/gin"
)

func (s *Services) index(c *gin.Context) {
	c.HTML(200, "index", gin.H{
		"title": "Домашняя страница",
		"judge": s.authenticatedJudge(c),
	})
}
