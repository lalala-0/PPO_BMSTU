package controllersApi

import (
	_ "PPO_BMSTU/docs"
	"PPO_BMSTU/internal/registry"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServicesAPI struct {
	Services *registry.Services
}

func SetupRouter(services *registry.Services, router *gin.Engine) {
	s := ServicesAPI{Services: services}

	// Публичные маршруты (GET)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // Путь к Swagger UI
	router.POST("/api/login", s.Login)
	router.POST("/api/logout", s.Logout)

	// judge routs
	judgeGroup := router.Group("/api/judges")
	{
		// Публичные GET маршруты
		judgeGroup.GET("/", s.getAllJudges)
		judgeGroup.GET("/:judgeID", s.getJudgeByID)

		// Защищённые маршруты
		judgeGroup.Use(s.JWTMiddleware())
		{
			judgeGroup.POST("/", s.createJudge)
			judgeGroup.PUT("/:judgeID", s.updateJudge)
			judgeGroup.DELETE("/:judgeID", s.deleteJudge)
		}
	}

	// rating routs
	ratingsGroup := router.Group("/api/ratings")
	{
		// Публичные GET маршруты
		ratingsGroup.GET("/", s.getAllRatings)
		ratingsGroup.GET("/:ratingID", s.getRating)
		ratingsGroup.GET("/:ratingID/rankings", s.getRankingTable)

		// Защищённые маршруты
		ratingsGroup.Use(s.JWTMiddleware())
		{
			ratingsGroup.POST("/", s.createRating)
			ratingsGroup.PUT("/:ratingID", s.updateRating)
			ratingsGroup.DELETE("/:ratingID", s.deleteRating)
		}
	}

	// race routs
	racesGroup := router.Group("/api/ratings/:ratingID/races")
	{
		// Публичные GET маршруты
		racesGroup.GET("/", s.getRacesByRatingID)
		racesGroup.GET("/:raceID", s.getRaceByID)

		// Защищённые маршруты
		racesGroup.Use(s.JWTMiddleware())
		{
			racesGroup.POST("/", s.createRace)
			racesGroup.PUT("/:raceID", s.updateRace)
			racesGroup.DELETE("/:raceID", s.deleteRace)
			racesGroup.POST("/:raceID/start", s.startProcedure)
			racesGroup.POST("/:raceID/finish", s.finishProcedure)
		}
	}

	// participants routs
	participantsGroup := router.Group("/api/participants")
	{
		// Публичные GET маршруты
		participantsGroup.GET("/", s.getAllParticipants)
		participantsGroup.GET("/:participantID", s.getParticipantById)

		// Защищённые маршруты
		participantsGroup.Use(s.JWTMiddleware())
		{
			participantsGroup.POST("/", s.createParticipant)
			participantsGroup.PUT("/:participantID", s.updateParticipant)
			participantsGroup.DELETE("/:participantID", s.deleteParticipant)
		}
	}

	// crews routs
	crewsGroup := router.Group("/api/ratings/:ratingID/crews")
	{
		// Публичные GET маршруты
		crewsGroup.GET("/", s.getCrewsByRatingID)
		crewsGroup.GET("/:crewID", s.getCrewByID)
		crewsGroup.GET("/:crewID/members", s.getCrewMembersByID)
		crewsGroup.GET("/:crewID/members/:participantID", s.getCrewMember)

		// Защищённые маршруты
		crewsGroup.Use(s.JWTMiddleware())
		{
			crewsGroup.POST("/:crewID/members", s.attachCrewMember)
			crewsGroup.PUT("/:crewID", s.updateCrewSailNumber)
			crewsGroup.DELETE("/:crewID", s.deleteCrewByID)
			crewsGroup.PUT("/:crewID/members/:participantID", s.updateCrewMember)
			crewsGroup.DELETE("/:crewID/members/:participantID", s.detachCrewMember)
		}
	}

	// protest routs
	protestsGroup := router.Group("/api/ratings/:ratingID/races/:raceID/protests")
	{
		// Публичные GET маршруты
		protestsGroup.GET("/", s.getProtests)
		protestsGroup.GET("/:protestID", s.getProtest)
		protestsGroup.GET("/:protestID/members", s.getProtestMembers)

		// Защищённые маршруты
		protestsGroup.Use(s.JWTMiddleware())
		{
			protestsGroup.POST("/", s.createProtest)
			protestsGroup.POST("/:protestID/complete", s.completeProtest)
			protestsGroup.DELETE("/:protestID", s.deleteProtest)
			protestsGroup.PUT("/:protestID", s.updateProtest)
			protestsGroup.POST("/:protestID/members", s.attachProtestMember)
			protestsGroup.DELETE("/:protestID/members/:crewSailNum", s.detachProtestMember)
		}
	}
}
