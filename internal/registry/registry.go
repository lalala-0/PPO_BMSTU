package registry

import (
	"PPO_BMSTU/config"
	"PPO_BMSTU/logger"
	"github.com/prometheus/client_golang/prometheus"

	// "PPO_BMSTU/internal/repository/mongo"
	"PPO_BMSTU/internal/repository/postgres"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	services "PPO_BMSTU/internal/services"
	"PPO_BMSTU/internal/services/service_interfaces"
	"PPO_BMSTU/password_hash"
	"github.com/charmbracelet/log"
	"os"
)

type Services struct {
	CrewService        service_interfaces.ICrewService
	JudgeService       service_interfaces.IJudgeService
	ParticipantService service_interfaces.IParticipantService
	ProtestService     service_interfaces.IProtestService
	RaceService        service_interfaces.IRaceService
	RatingService      service_interfaces.IRatingService
}

type Repositories struct {
	CrewRepository          repository_interfaces.ICrewRepository
	CrewResInRaceRepository repository_interfaces.ICrewResInRaceRepository
	JudgeRepository         repository_interfaces.IJudgeRepository
	ParticipantRepository   repository_interfaces.IParticipantRepository
	ProtestRepository       repository_interfaces.IProtestRepository
	RaceRepository          repository_interfaces.IRaceRepository
	RatingRepository        repository_interfaces.IRatingRepository
}

type App struct {
	Config       config.Config
	Repositories *Repositories
	Services     *Services
	Logger       *logger.CustomLogger
}

func (a *App) postgresRepositoriesInitialization(fields *postgres.PostgresConnection) *Repositories {
	r := &Repositories{
		CrewRepository:          postgres.CreateCrewRepository(fields),
		CrewResInRaceRepository: postgres.CreateCrewResInRaceRepository(fields),
		JudgeRepository:         postgres.CreateJudgeRepository(fields),
		ParticipantRepository:   postgres.CreateParticipantRepository(fields),
		ProtestRepository:       postgres.CreateProtestRepository(fields),
		RaceRepository:          postgres.CreateRaceRepository(fields),
		RatingRepository:        postgres.CreateRatingRepository(fields),
	}
	a.Logger.Info("Success initialization of repositories")
	return r
}

//func (a *App) mongoRepositoriesInitialization(fields *mongo.MongoConnection) *Repositories {
//	r := &Repositories{
//		CrewRepository:          mongo.CreateCrewRepository(fields),
//		CrewResInRaceRepository: mongo.CreateCrewResInRaceRepository(fields),
//		JudgeRepository:         mongo.CreateJudgeRepository(fields),
//		ParticipantRepository:   mongo.CreateParticipantRepository(fields),
//		ProtestRepository:       mongo.CreateProtestRepository(fields),
//		RaceRepository:          mongo.CreateRaceRepository(fields),
//		RatingRepository:        mongo.CreateRatingRepository(fields),
//	}
//	a.Logger.Info("Success initialization of repositories")
//	return r
//}

func (a *App) servicesInitialization(r *Repositories) *Services {
	passwordHash := password_hash.NewPasswordHash()

	s := &Services{
		CrewService:        services.NewCrewService(r.CrewRepository, a.Logger),
		JudgeService:       services.NewJudgeService(r.JudgeRepository, passwordHash, a.Logger),
		ParticipantService: services.NewParticipantService(r.ParticipantRepository, a.Logger),
		ProtestService:     services.NewProtestService(r.ProtestRepository, r.CrewResInRaceRepository, r.CrewRepository, a.Logger),
		RaceService:        services.NewRaceService(r.RaceRepository, r.CrewRepository, r.CrewResInRaceRepository, a.Logger),
		RatingService:      services.NewRatingService(r.RatingRepository, r.JudgeRepository, a.Logger),
	}
	a.Logger.Info("Success initialization of services")

	return s
}

func (a *App) initLogger() {
	// Открытие файла для записи логов
	f, err := os.OpenFile(a.Config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// Инициализация метрик для логирования
	logCount := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "log_count_total",
		Help: "Total number of logs recorded by log level",
	}, []string{"level"})

	logDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "log_duration_seconds",
		Help:    "Duration of log writing in seconds",
		Buckets: prometheus.DefBuckets,
	}, []string{"level"})

	// Регистрация метрик
	prometheus.MustRegister(logCount)
	prometheus.MustRegister(logDuration)

	// Инициализация логгера
	baseLogger := log.New(f)

	baseLogger.SetReportTimestamp(true)
	baseLogger.SetReportCaller(true)

	// Установка уровня логирования
	if a.Config.LogLevel == "debug" {
		baseLogger.SetLevel(log.DebugLevel)
	} else if a.Config.LogLevel == "info" {
		baseLogger.SetLevel(log.InfoLevel)
	} else {
		log.Fatal("Error log level")
	}

	// Оборачивание логгера в кастомный логгер с метриками
	a.Logger = &logger.CustomLogger{
		Logger:      baseLogger,
		LogCount:    logCount,
		LogDuration: logDuration,
	}

	a.Logger.Info("Success initialization of new Logger!")
}

func (a *App) Init() error {
	a.initLogger()

	if a.Config.DBType == "postgres" {
		fields, err := postgres.NewPostgresConnection(a.Config.DBFlags, a.Logger)
		if err != nil {
			a.Logger.Fatal("Error create postgres repository fields", "err", err)
			return err
		}

		a.Repositories = a.postgresRepositoriesInitialization(fields)
		a.Services = a.servicesInitialization(a.Repositories)
	}
	//} else if a.Config.DBType == "mongo" {
	//	fields, err := mongo.NewMongoConnection(a.Config.DBFlags, a.Logger)
	//	if err != nil {
	//		a.Logger.Fatal("Error create mongodb repository fields", "err", err)
	//		return err
	//	}
	//	a.Repositories = a.mongoRepositoriesInitialization(fields)
	//	a.Services = a.servicesInitialization(a.Repositories)
	//
	//}
	return nil
}

func (a *App) Run() error {
	err := a.Init()

	if err != nil {
		a.Logger.Error("Error init app", "err", err)
		return err
	}

	return nil
}
