package controllers

//
//import (
//	"PPO_BMSTU/internal/models"
//	"log"
//	"net/http"
//
//	"github.com/gin-gonic/gin"
//	"github.com/google/uuid"
//)
//
//func (s *Services) ratings(c *gin.Context) {
//	judge := s.authenticatedJudge(c)
//
//	ratings, err := s.Services.RatingService.GetAllRatings()
//	if err != nil {
//		log.Printf("Error getting ratings: %v", err)
//		ratings = []models.Rating{}
//	}
//
//	c.HTML(http.StatusOK, "ratingsList", gin.H{
//		"title":   "Рейтинги",
//		"judge":   judge,
//		"ratings": ratings,
//	})
//}
//
//func (s *Services) createRatingGet(c *gin.Context) {
//	judge := s.authenticatedJudge(c)
//	categories, err := s.Services.CategoryService.GetAll()
//	if err != nil {
//		log.Printf("Error getting categories: %v", err)
//		categories = []models.Category{}
//	}
//
//	c.HTML(http.StatusOK, "createService", gin.H{
//		"title":      "Создать услугу",
//		"judge":      judge,
//		"categories": categories,
//	})
//}
//
//type ServiceFormData struct {
//	Name           string  `form:"name"`
//	PricePerSingle float64 `form:"pricePerSingle"`
//	Category       int     `form:"category"`
//}
//
//func (s *Services) createServicePost(c *gin.Context) {
//	judge := s.authenticatedJudge(c)
//
//	var data ServiceFormData
//	if err := c.Bind(&data); err != nil {
//		c.HTML(http.StatusBadRequest, "createService", gin.H{
//			"judge":    judge,
//			"title":    "Создать услугу",
//			"error":    err.Error(),
//			"formData": data,
//		})
//		return
//	}
//
//	_, err := s.Services.TaskService.Create(data.Name, data.PricePerSingle, data.Category)
//	if err != nil {
//		c.HTML(http.StatusBadRequest, "createService", gin.H{
//			"judge":    judge,
//			"title":    "Создать услугу",
//			"error":    err.Error(),
//			"formData": data,
//		})
//		return
//	}
//
//	c.Redirect(http.StatusFound, "/services")
//}
//
//func (s *Services) editServiceGet(c *gin.Context) {
//	judge := s.authenticatedJudge(c)
//	serviceID, err := uuid.Parse(c.Param("id"))
//	if err != nil {
//		c.HTML(http.StatusBadRequest, "createService", gin.H{
//			"judge": judge,
//			"title": "Создать услугу",
//			"error": "Неверный идентификатор услуги",
//		})
//		return
//	}
//
//	service, err := s.Services.TaskService.GetTaskByID(serviceID)
//	if err != nil {
//		c.HTML(http.StatusBadRequest, "createService", gin.H{
//			"judge": judge,
//			"title": "Создать услугу",
//			"error": "Услуга не найдена",
//		})
//		return
//	}
//
//	formData := ServiceFormData{
//		Name:           service.Name,
//		PricePerSingle: service.PricePerSingle,
//		Category:       service.Category,
//	}
//
//	c.HTML(http.StatusOK, "createService", gin.H{
//		"title":    "Создать услугу",
//		"judge":    judge,
//		"formData": formData,
//	})
//}
//
//func (s *Services) editServicePost(c *gin.Context) {
//	judge := s.authenticatedJudge(c)
//	serviceID, err := uuid.Parse(c.Param("id"))
//
//	var data ServiceFormData
//	if err := c.Bind(&data); err != nil {
//		c.HTML(http.StatusBadRequest, "createService", gin.H{
//			"judge":    judge,
//			"title":    "Изменить услугу",
//			"error":    err.Error(),
//			"formData": data,
//		})
//		return
//	}
//
//	_, err = s.Services.TaskService.Update(serviceID, data.Category, data.Name, data.PricePerSingle)
//	if err != nil {
//		c.HTML(http.StatusBadRequest, "createService", gin.H{
//			"judge":    judge,
//			"title":    "Изменить услугу",
//			"error":    err.Error(),
//			"formData": data,
//		})
//		return
//	}
//
//	c.Redirect(http.StatusFound, "/services")
//}
