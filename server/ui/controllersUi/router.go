package controllersUi

import (
	"PPO_BMSTU/internal/registry"
	"PPO_BMSTU/server/ui/middleware"
	"github.com/gin-gonic/gin"
)

type ServicesUI struct {
	Services *registry.Services
}

func SetupRouter(services *registry.Services, router *gin.Engine) {
	s := ServicesUI{Services: services}
	authMiddleware := middleware.NewMiddleware(services)

	router.LoadHTMLGlob("server/ui/templates/**/*")

	router.GET("ui/", s.index)

	authGroup := router.Group("ui/auth")
	{
		authGroup.GET("/signin", s.signinGet)
		authGroup.POST("/signin", s.signinPost)

		authGroup.GET("/logout", s.logout)
	}

	judgeGroup := router.Group("ui/judges")
	judgeGroup.Use(authMiddleware.JudgeMiddleware())
	{
		judgeGroup.GET("/", s.menu)

		judgeGroup.GET("/:judgeID/profile", s.judgeProfile)

		judgeGroup.GET("/:judgeID/profile/updatePassword", s.updatePasswordGet)
		judgeGroup.POST("/:judgeID/profile/updatePassword", s.updatePasswordPost)

		judgeGroup.GET("/:judgeID/profile/update", s.updateJudgeGet)
		judgeGroup.POST("/:judgeID/profile/update", s.updateJudgePost)

		judgeGroup.GET("/:judgeID", s.getJudgeMenu)

		judgeGroup.POST("/:judgeID/delete", s.deleteJudge)

		judgeGroup.GET("/create", s.createJudgeGet)
		judgeGroup.POST("/create", s.createJudgePost)

		judgeGroup.GET("/:judgeID/update", s.updateJudgeGet)
		judgeGroup.POST("/:judgeID/update", s.updateJudgePost)
	}

	ratingsGroup := router.Group("ui/ratings")
	ratingsGroup.Use()
	{
		ratingsGroup.GET("/create", s.createRatingGet)
		ratingsGroup.POST("/create", s.createRatingPost)

		ratingsGroup.GET("/:ratingID/update", s.updateRatingGet)
		ratingsGroup.POST("/:ratingID/update", s.updateRatingPost)

		ratingsGroup.POST("/:ratingID/delete", s.deleteRating)

		ratingsGroup.GET("/:ratingID", s.getRatingMenu)
		ratingsGroup.GET("/:ratingID/ratingTable", s.getRatingTable)
	}

	racesGroup := router.Group("ui/ratings/:ratingID/races")
	racesGroup.Use()
	{
		racesGroup.GET("/:raceID", s.getRaceMenu)

		racesGroup.GET("/create", s.createRaceGet)
		racesGroup.POST("/create", s.createRacePost)

		racesGroup.POST("/:raceID/delete", s.deleteRace)

		racesGroup.GET("/:raceID/update", s.updateRaceGet)
		racesGroup.POST("/:raceID/update", s.updateRacePost)

		racesGroup.GET("/:raceID/start", s.startRaceGet)
		racesGroup.POST("/:raceID/start", s.startRacePost)

		racesGroup.GET("/:raceID/finish", s.finishRaceGet)
		racesGroup.POST("/:raceID/finish", s.finishRacePost)
	}

	participantsGroup := router.Group("ui/ratings/:ratingID/crews/:crewID/participants")
	participantsGroup.Use(authMiddleware.JudgeMiddleware())
	{
		participantsGroup.GET("/:participantID", s.getParticipantMenu)

		participantsGroup.POST("/:participantID/delete", s.deleteParticipant)

		participantsGroup.GET("/create", s.createParticipantGet)
		participantsGroup.POST("/create", s.createParticipantPost)

		participantsGroup.GET("/:participantID/update", s.updateParticipantGet)
		participantsGroup.POST("/:participantID/update", s.updateParticipantPost)
	}

	participantsShortGroup := router.Group("ui/participants")
	participantsShortGroup.Use(authMiddleware.JudgeMiddleware())
	{
		participantsShortGroup.GET("/:participantID", s.getParticipantMenu)

		participantsShortGroup.POST("/:participantID/delete", s.deleteParticipant)

		participantsShortGroup.GET("/create", s.createParticipantGet)
		participantsShortGroup.POST("/create", s.createParticipantPost)

		participantsShortGroup.GET("/:participantID/update", s.updateParticipantGet)
		participantsShortGroup.POST("/:participantID/update", s.updateParticipantPost)
	}

	crewsGroup := router.Group("ui/ratings/:ratingID/crews")
	crewsGroup.Use()
	{
		crewsGroup.GET("/:crewID", s.getCrewMenu)

		crewsGroup.GET("/create", s.createCrewGet)
		crewsGroup.POST("/create", s.createCrewPost)

		crewsGroup.GET("/:crewID/update", s.updateCrewGet)
		crewsGroup.POST("/:crewID/update", s.updateCrewPost)

		crewsGroup.POST("/:crewID/delete", s.deleteCrew)

		crewsGroup.GET("/:crewID/attach", s.attachCrewParticipantGet)
		crewsGroup.POST("/:crewID/attach", s.attachCrewParticipantPost)

		crewsGroup.GET("/:crewID/detach", s.detachCrewParticipantGet)
		crewsGroup.POST("/:crewID/detach", s.detachCrewParticipantPost)

	}

	protestsGroup := router.Group("ui/ratings/:ratingID/races/:raceID/protests")
	protestsGroup.Use()
	{
		protestsGroup.GET("/:protestID", s.getProtestMenu)

		protestsGroup.GET("/create", s.createProtestGet)
		protestsGroup.POST("/create", s.createProtestPost)

		protestsGroup.GET("/:protestID/update", s.updateProtestGet)
		protestsGroup.POST("/:protestID/update", s.updateProtestPost)

		protestsGroup.POST("/:protestID/delete", s.deleteProtest)

		protestsGroup.GET("/:protestID/attach", s.attachProtestParticipantGet)
		protestsGroup.POST("/:protestID/attach", s.attachProtestParticipantPost)

		protestsGroup.GET("/:protestID/detach", s.detachProtestParticipantGet)
		protestsGroup.POST("/:protestID/detach", s.detachProtestParticipantPost)

		protestsGroup.GET("/:protestID/complete", s.completeProtestGet)
		protestsGroup.POST("/:protestID/complete", s.completeProtestPost)
	}
	return
}
