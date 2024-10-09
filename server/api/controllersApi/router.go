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

	judgeGroup := router.Group("/api/judges")

	judgeGroup.GET("/", s.getAllJudges)

	//judgeGroup.GET("/:judgeID/profile", s.judgeProfile)
	//
	//judgeGroup.GET("/:judgeID/profile/updatePassword", s.updatePasswordGet)
	//judgeGroup.POST("/:judgeID/profile/updatePassword", s.updatePasswordPost)
	//
	//judgeGroup.GET("/:judgeID/profile/update", s.updateJudgeGet)
	//judgeGroup.POST("/:judgeID/profile/update", s.updateJudgePost)
	//
	//judgeGroup.GET("/:judgeID", s.getJudgeMenu)
	//
	//judgeGroup.POST("/:judgeID/delete", s.deleteJudge)
	//
	//judgeGroup.GET("/create", s.createJudgeGet)
	//judgeGroup.POST("/create", s.createJudgePost)
	//
	//judgeGroup.GET("/:judgeID/update", s.updateJudgeGet)
	//judgeGroup.POST("/:judgeID/update", s.updateJudgePost)
	//
	//
	ratingsGroup := router.Group("/api/ratings")

	ratingsGroup.GET("/", s.getAllRatings)
	ratingsGroup.POST("/", s.createRating)

	ratingsGroup.GET("/:ratingID", s.getRating)
	ratingsGroup.PUT("/:ratingID", s.updateRating)
	ratingsGroup.DELETE("/:ratingID", s.deleteRating)

	//	ratingsGroup.GET("/:ratingID/ratingTable", s.getRatingTable)
	//
	racesGroup := router.Group("/api/ratings/:ratingID/races")

	racesGroup.GET("/", s.getRacesByRatingID)
	racesGroup.POST("/", s.createRace)

	racesGroup.GET("/:raceID", s.getRaceByID)
	racesGroup.PUT("/:raceID", s.updateRace)
	racesGroup.DELETE("/:raceID", s.deleteRace)

	racesGroup.POST("/:raceID/start", s.startProcedure)
	racesGroup.POST("/:raceID/finish", s.finishProcedure)

	//participantsGroup := router.Group("ratings/:ratingID/crews/:crewID/participants")
	//participantsGroup.Use(authMiddleware.JudgeMiddleware())
	//{
	//	participantsGroup.GET("/:participantID", s.getParticipantMenu)
	//
	//	participantsGroup.POST("/:participantID/delete", s.deleteParticipant)
	//
	//	participantsGroup.GET("/create", s.createParticipantGet)
	//	participantsGroup.POST("/create", s.createParticipantPost)
	//
	//	participantsGroup.GET("/:participantID/update", s.updateParticipantGet)
	//	participantsGroup.POST("/:participantID/update", s.updateParticipantPost)
	//}
	//
	//participantsShortGroup := router.Group("participants")
	//participantsShortGroup.Use(authMiddleware.JudgeMiddleware())
	//{
	//	participantsShortGroup.GET("/:participantID", s.getParticipantMenu)
	//
	//	participantsShortGroup.POST("/:participantID/delete", s.deleteParticipant)
	//
	//	participantsShortGroup.GET("/create", s.createParticipantGet)
	//	participantsShortGroup.POST("/create", s.createParticipantPost)
	//
	//	participantsShortGroup.GET("/:participantID/update", s.updateParticipantGet)
	//	participantsShortGroup.POST("/:participantID/update", s.updateParticipantPost)
	//}
	//
	// crewsGroup := router.Group("/api/ratings/:ratingID/crews")

	//	crewsGroup.GET("/:crewID", s.getCrewMenu)
	//
	//	crewsGroup.GET("/create", s.createCrewGet)
	//	crewsGroup.POST("/create", s.createCrewPost)
	//
	//	crewsGroup.GET("/:crewID/update", s.updateCrewGet)
	//	crewsGroup.POST("/:crewID/update", s.updateCrewPost)
	//
	//	crewsGroup.POST("/:crewID/delete", s.deleteCrew)
	//
	//	crewsGroup.GET("/:crewID/attach", s.attachCrewParticipantGet)
	//	crewsGroup.POST("/:crewID/attach", s.attachCrewParticipantPost)
	//
	//	crewsGroup.GET("/:crewID/detach", s.detachCrewParticipantGet)
	//	crewsGroup.POST("/:crewID/detach", s.detachCrewParticipantPost)
	//
	//
	//protestsGroup := router.Group("/ratings/:ratingID/races/:raceID/protests")
	//protestsGroup.Use()
	//{
	//	protestsGroup.GET("/:protestID", s.getProtestMenu)
	//
	//	protestsGroup.GET("/create", s.createProtestGet)
	//	protestsGroup.POST("/create", s.createProtestPost)
	//
	//	protestsGroup.GET("/:protestID/update", s.updateProtestGet)
	//	protestsGroup.POST("/:protestID/update", s.updateProtestPost)
	//
	//	protestsGroup.POST("/:protestID/delete", s.deleteProtest)
	//
	//	protestsGroup.GET("/:protestID/attach", s.attachProtestParticipantGet)
	//	protestsGroup.POST("/:protestID/attach", s.attachProtestParticipantPost)
	//
	//	protestsGroup.GET("/:protestID/detach", s.detachProtestParticipantGet)
	//	protestsGroup.POST("/:protestID/detach", s.detachProtestParticipantPost)
	//
	//	protestsGroup.GET("/:protestID/complete", s.completeProtestGet)
	//	protestsGroup.POST("/:protestID/complete", s.completeProtestPost)
	//}
	// Путь к Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return
}
