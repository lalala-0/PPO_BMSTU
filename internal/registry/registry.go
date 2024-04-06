package registry

//
//import (
//	"PPO_BMSTU/config"
//	"PPO_BMSTU/internal/repository/repository_interfaces"
//	services "PPO_BMSTU/internal/services"
//	"PPO_BMSTU/internal/services/service_interfaces"
//	"PPO_BMSTU/password_hash"
//	"github.com/charmbracelet/log"
//	"os"
//)
//
//type Services struct {
//	CrewService        service_interfaces.ICrewService
//	JudgeService       service_interfaces.IJudgeService
//	ParticipantService service_interfaces.IParticipantService
//	ProtestService     service_interfaces.IProtestService
//	RaceService        service_interfaces.IRaceService
//	RatingService      service_interfaces.IRatingService
//}
//
//type Repositories struct {
//	CrewRepository          repository_interfaces.ICrewRepository
//	JudgeRepository         repository_interfaces.IJudgeRepository
//	ParticipantRepository   repository_interfaces.IParticipantRepository
//	ProtestRepository       repository_interfaces.IProtestRepository
//	RaceRepository          repository_interfaces.IRaceRepository
//	RatingRepository        repository_interfaces.IRatingRepository
//	CrewResInRaceRepository repository_interfaces.ICrewResInRaceRepository
//}
//
//type App struct {
//	Config       config.Config
//	Repositories *Repositories
//	Services     *Services
//	Logger       *log.Logger
//}
//
//func (a *App) repositoriesInitialization(fields *repositories.PostgresConnection) *Repositories {
//	r := &Repositories{
//		UserRepository:   repositories.CreateUserRepository(fields),
//		WorkerRepository: repositories.CreateWorkerRepository(fields),
//		TaskRepository:   repositories.CreateTaskRepository(fields),
//		OrderRepository:  repositories.CreateOrderRepository(fields),
//	}
//	a.Logger.Info("Success initialization of repositories")
//	return r
//}
//
//func (a *App) servicesInitialization(r *Repositories) *Services {
//	passwordHash := password_hash.NewPasswordHash()
//
//	s := &Services{
//		UserService:   services.NewUserService(r.UserRepository, passwordHash, a.Logger),
//		WorkerService: services.NewWorkerService(r.WorkerRepository, passwordHash, a.Logger),
//		OrderService:  services.NewOrderService(r.OrderRepository, r.WorkerRepository, r.TaskRepository, a.Logger),
//		TaskService:   services.NewTaskService(r.TaskRepository, a.Logger),
//	}
//	a.Logger.Info("Success initialization of services")
//
//	return s
//}
//
//func (a *App) initLogger() {
//	f, err := os.OpenFile(a.Config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	Logger := log.New(f)
//
//	log.SetFormatter(log.LogfmtFormatter)
//	Logger.SetReportTimestamp(true)
//	Logger.SetReportCaller(true)
//
//	if a.Config.LogLevel == "debug" {
//		Logger.SetLevel(log.DebugLevel)
//	} else if a.Config.LogLevel == "info" {
//		Logger.SetLevel(log.InfoLevel)
//	} else {
//		log.Fatal("Error log level")
//	}
//
//	Logger.Info("Success initialization of new Logger!")
//
//	a.Logger = Logger
//}
//
//func (a *App) Init() error {
//	a.initLogger()
//
//	fields, err := repositories.NewPostgresConnection(a.Config.Postgres, a.Logger)
//	if err != nil {
//		a.Logger.Fatal("Error create postgres repository fields", "err", err)
//		return err
//	}
//
//	a.Repositories = a.repositoriesInitialization(fields)
//	a.Services = a.servicesInitialization(a.Repositories)
//
//	return nil
//}
//
//func (a *App) Run() error {
//	err := a.Init()
//
//	if err != nil {
//		a.Logger.Error("Error init app", "err", err)
//		return err
//	}
//
//	return nil
//}
