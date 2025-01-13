package controllersApi

import (
	"github.com/google/uuid"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GenerateCodeHandler обрабатывает запрос на генерацию 2FA-кода.
func (c *ServicesAPI) GenerateCodeHandler(ctx *gin.Context) {
	judgeID := ctx.Param("judgeID")

	judgeUUID, err := uuid.Parse(judgeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	judge, err := c.Services.JudgeService.GetJudgeDataByID(judgeUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	code, err := c.Services.TwoFA.GenerateAndStoreCode(judgeID, judge.Login)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "2FA code generated", "code": code})
}

// VerifyCodeHandler обрабатывает запрос на проверку 2FA-кода.
func (c *ServicesAPI) VerifyCodeHandler(ctx *gin.Context) {
	judgeID := ctx.Param("judgeID")
	var request struct {
		Code string `json:"code" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isValid, err := c.Services.TwoFA.VerifyCode(judgeID, request.Code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if isValid {
		ctx.JSON(http.StatusOK, gin.H{"message": "Code is valid"})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired code"})
	}
}
