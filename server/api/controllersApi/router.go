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

	// judge routs
	judgeGroup := router.Group("/api/judges")

	judgeGroup.GET("/", s.getAllJudges)
	judgeGroup.POST("/", s.createJudge)

	judgeGroup.GET("/:judgeID", s.getJudgeByID)
	judgeGroup.PUT("/:judgeID", s.updateJudge)
	judgeGroup.DELETE("/:judgeID", s.deleteJudge)
	//
	//judgeGroup.GET("/:judgeID/profile/updatePassword", s.updatePasswordGet)
	//judgeGroup.POST("/:judgeID/profile/updatePassword", s.updatePasswordPost)
	//

	// rating routs
	ratingsGroup := router.Group("/api/ratings")

	ratingsGroup.GET("/", s.getAllRatings)
	ratingsGroup.POST("/", s.createRating)

	ratingsGroup.GET("/:ratingID", s.getRating)
	ratingsGroup.PUT("/:ratingID", s.updateRating)
	ratingsGroup.DELETE("/:ratingID", s.deleteRating)

	ratingsGroup.GET("/:ratingID/rankings", s.getRankingTable)
	//

	// race routs
	racesGroup := router.Group("/api/ratings/:ratingID/races")

	racesGroup.GET("/", s.getRacesByRatingID)
	racesGroup.POST("/", s.createRace)

	racesGroup.GET("/:raceID", s.getRaceByID)
	racesGroup.PUT("/:raceID", s.updateRace)
	racesGroup.DELETE("/:raceID", s.deleteRace)

	racesGroup.POST("/:raceID/start", s.startProcedure)
	racesGroup.POST("/:raceID/finish", s.finishProcedure)

	// participant routs
	participantsGroup := router.Group("/api/participants")

	participantsGroup.GET("/", s.getAllParticipants)
	participantsGroup.POST("/", s.createParticipant)

	participantsGroup.GET("/:participantID", s.getParticipantById)
	participantsGroup.PUT("/:participantID", s.updateParticipant)
	participantsGroup.DELETE("/:participantID", s.deleteParticipant)

	// crew routs
	crewsGroup := router.Group("/api/ratings/:ratingID/crews")

	crewsGroup.GET("/", s.getCrewsByRatingID)
	crewsGroup.POST("/", s.createCrew)

	crewsGroup.GET("/:crewID", s.getCrewByID)
	crewsGroup.PUT("/:crewID", s.updateCrewSailNumber)
	crewsGroup.DELETE("/:crewID", s.deleteCrewByID)

	crewsGroup.GET("/:crewID/members", s.getCrewMembersByID)
	crewsGroup.POST("/:crewID/members", s.attachCrewMember)

	crewsGroup.GET("/:crewID/members/:participantID", s.getCrewMember)
	crewsGroup.PUT("/:crewID/members/:participantID", s.updateCrewMember)
	crewsGroup.DELETE("/:crewID/members/:participantID", s.detachCrewMember)

	// protest routs

	protestsGroup := router.Group("/api/ratings/:ratingID/races/:raceID/protests")

	protestsGroup.GET("/", s.getProtests)                          // Получить все протесты
	protestsGroup.POST("/", s.createProtest)                       // Создать новый протест
	protestsGroup.GET("/:protestID", s.getProtest)                 // Получить информацию о протесте
	protestsGroup.PATCH("/:protestID/complete", s.completeProtest) // Завершить рассмотрение протеста
	protestsGroup.DELETE("/:protestID", s.deleteProtest)           // Удалить протест
	protestsGroup.PUT("/:protestID", s.updateProtest)              // Обновить информацию о протесте

	// Маршруты для участников протеста
	protestsGroup.POST("/:protestID/members", s.attachProtestMember)                // Добавить команду-участника протеста
	protestsGroup.DELETE("/:protestID/members/:crewSailNum", s.detachProtestMember) // Добавить команду-участника протеста
	protestsGroup.GET("/:protestID/members", s.getProtestMembers)                   // Получить информацию о всех командах-участниках протеста

	// Путь к Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return
}
