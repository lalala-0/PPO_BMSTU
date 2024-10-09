package controllersUi

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/server/ui/modelsUI"
	"github.com/gin-gonic/gin"
	"log"
)

func (s *ServicesUI) index(c *gin.Context) {

	ratings, err := s.Services.RatingService.GetAllRatings()
	if err != nil {
		log.Printf("Error getting ratings: %v", err)
		ratings = []models.Rating{}
	}
	ratingsData, err := modelsUI.FromRatingModelsToStringData(ratings)
	if err != nil {
		log.Printf("Error class to string convert: %v", err)
	}

	c.HTML(200, "index", gin.H{
		"title":   "Список рейтингов",
		"judge":   s.authenticatedJudge(c),
		"ratings": ratingsData,
	})
}
