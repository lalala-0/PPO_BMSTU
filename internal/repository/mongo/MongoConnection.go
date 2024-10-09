package mongo

import (
	"PPO_BMSTU/config"
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"github.com/charmbracelet/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoConnection struct {
	DB     *mongo.Database
	Config config.Config
}

func NewMongoConnection(Mongo config.DbConnectionFlags, logger *log.Logger) (*MongoConnection, error) {
	fields := new(MongoConnection)
	var err error

	fields.Config.DBFlags = Mongo

	fields.DB, err = fields.Config.DBFlags.InitMongoDB(logger)
	if err != nil {
		logger.Error("MONGO! Error parse config for mongoDB")
		return nil, repository_errors.ConnectionError
	}

	logger.Info("Mongo! Successfully create mongo repository fields")

	return fields, nil
}

func CreateCrewRepository(fields *MongoConnection) repository_interfaces.ICrewRepository {
	return NewCrewRepository(fields.DB)
}

func CreateJudgeRepository(fields *MongoConnection) repository_interfaces.IJudgeRepository {
	return NewJudgeRepository(fields.DB)
}

func CreateCrewResInRaceRepository(fields *MongoConnection) repository_interfaces.ICrewResInRaceRepository {
	return NewCrewResInRaceRepository(fields.DB)
}

func CreateParticipantRepository(fields *MongoConnection) repository_interfaces.IParticipantRepository {
	return NewParticipantRepository(fields.DB)
}

func CreateProtestRepository(fields *MongoConnection) repository_interfaces.IProtestRepository {
	return NewProtestRepository(fields.DB)
}
func CreateRaceRepository(fields *MongoConnection) repository_interfaces.IRaceRepository {
	return NewRaceRepository(fields.DB)
}
func CreateRatingRepository(fields *MongoConnection) repository_interfaces.IRatingRepository {
	return NewRatingRepository(fields.DB)
}
