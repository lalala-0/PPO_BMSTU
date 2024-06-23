package controllers

import (
	"PPO_BMSTU/internal/registry"
	"PPO_BMSTU/ui/middleware"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Services struct {
	Services *registry.Services
}

func RunServer(app *registry.App) error {
	s := Services{
		app.Services,
	}

	router := s.setupRouter(app)

	gin.SetMode(gin.DebugMode)

	port := app.Config.Port
	address := app.Config.Address
	err := router.Run(address + port)
	return err
}

func (s *Services) setupRouter(app *registry.App) *gin.Engine {
	authMiddleware := middleware.NewMiddleware(*app)

	router := gin.Default()

	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.LoadHTMLGlob("ui/templates/**/*")

	router.GET("/", s.index)
	//router.GET("/ratings", s.ratings)

	authGroup := router.Group("/auth")
	{
		authGroup.GET("/signin", s.signinGet)
		authGroup.POST("/signin", s.signinPost)

		authGroup.GET("/logout", s.logout)
	}

	judgeGroup := router.Group("/judge")
	judgeGroup.Use(authMiddleware.JudgeMiddleware())
	{
		judgeGroup.GET("/", s.menu)

		judgeGroup.GET("/profile", s.judgeProfile)

		judgeGroup.GET("/directory", s.judgesDirectory)

		judgeGroup.GET("/create", s.createJudgeGet)
		judgeGroup.POST("/create", s.createJudgePost)

		judgeGroup.GET("/:id", s.judgeDetails)

		judgeGroup.GET("/:id/edit", s.editJudgeGet)
		judgeGroup.POST("/:id/edit", s.editJudgePost)
	}

	//ratingsGroup := router.Group("/ratings")
	//ratingsGroup.Use(authMiddleware.JudgeMiddleware())
	//{
	//	ratingsGroup.GET("/", s.ratings)
	//	ratingsGroup.GET("/create", s.createRatingGet)
	//	ratingsGroup.POST("/create", s.createRatingPost)
	//	ratingsGroup.GET("/:id", s.editRatingGet)
	//	ratingsGroup.POST("/:id", s.editRatingPost)
	//}
	//
	//racesGroup := router.Group("/races")
	//racesGroup.Use(authMiddleware.JudgeMiddleware())
	//{
	//	racesGroup.GET("/create", s.createRaceGet)
	//	racesGroup.POST("/create", s.createRacePost)
	//
	//	racesGroup.GET("/:id", s.editRaceGet)
	//	racesGroup.POST("/:id", s.editRacePost)
	//}
	//
	//participantsGroup := router.Group("/participants")
	//participantsGroup.Use(authMiddleware.JudgeMiddleware())
	//{
	//	participantsGroup.GET("/create", s.createParticipantGet)
	//	participantsGroup.POST("/create", s.createParticipantPost)
	//
	//	participantsGroup.GET("/:id", s.editParticipantGet)
	//	participantsGroup.POST("/:id", s.editParticipantPost)
	//}
	//
	//crewsGroup := router.Group("/crews")
	//crewsGroup.Use(authMiddleware.JudgeMiddleware())
	//{
	//	crewsGroup.GET("/create", s.createCrewGet)
	//	crewsGroup.POST("/create", s.createCrewPost)
	//
	//	crewsGroup.GET("/:id", s.editCrewGet)
	//	crewsGroup.POST("/:id", s.editCrewPost)
	//}
	//
	//protestsGroup := router.Group("/protests")
	//protestsGroup.Use(authMiddleware.JudgeMiddleware())
	//{
	//	protestsGroup.GET("/create", s.createProtestGet)
	//	protestsGroup.POST("/create", s.createProtestPost)
	//
	//	protestsGroup.GET("/:id", s.editProtestGet)
	//	protestsGroup.POST("/:id", s.editProtestPost)
	//}
	return router
}
