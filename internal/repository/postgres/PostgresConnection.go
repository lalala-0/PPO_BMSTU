package postgres

import (
	"PPO_BMSTU/config"
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"PPO_BMSTU/logger"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/otel"
)

type PostgresConnection struct {
	DB     *sql.DB
	Config config.Config
}

func NewPostgresConnection(Postgres config.DbConnectionFlags, logger *logger.CustomLogger) (*PostgresConnection, error) {
	fields := new(PostgresConnection)
	var err error

	fields.Config.DBFlags = Postgres

	fields.DB, err = fields.Config.DBFlags.InitPostgresDB(logger)
	if err != nil {
		logger.Error("POSTGRES! Error parse config for postgreSQL")
		return nil, repository_errors.ConnectionError
	}

	logger.Info("POSTGRES! Successfully create postgres repository fields")

	return fields, nil
}
func CreateCrewRepository(fields *PostgresConnection) repository_interfaces.ICrewRepository {
	dbx := sqlx.NewDb(fields.DB, "pgx")
	tracer := otel.Tracer("repository.Crew")
	tracedDB := NewTracedDB(dbx, tracer)

	return NewCrewRepository(tracedDB)
}

func CreateJudgeRepository(fields *PostgresConnection) repository_interfaces.IJudgeRepository {
	dbx := sqlx.NewDb(fields.DB, "pgx")
	tracer := otel.Tracer("repository.Judge")
	tracedDB := NewTracedDB(dbx, tracer)

	return NewJudgeRepository(tracedDB)
}

func CreateCrewResInRaceRepository(fields *PostgresConnection) repository_interfaces.ICrewResInRaceRepository {
	dbx := sqlx.NewDb(fields.DB, "pgx")
	tracer := otel.Tracer("repository.CrewResInRace")
	tracedDB := NewTracedDB(dbx, tracer)

	return NewCrewResInRaceRepository(tracedDB)
}

func CreateParticipantRepository(fields *PostgresConnection) repository_interfaces.IParticipantRepository {
	dbx := sqlx.NewDb(fields.DB, "pgx")
	tracer := otel.Tracer("repository.Participant")
	tracedDB := NewTracedDB(dbx, tracer)

	return NewParticipantRepository(tracedDB)
}

func CreateProtestRepository(fields *PostgresConnection) repository_interfaces.IProtestRepository {
	dbx := sqlx.NewDb(fields.DB, "pgx")
	tracer := otel.Tracer("repository.Protest")
	tracedDB := NewTracedDB(dbx, tracer)

	return NewProtestRepository(tracedDB)
}

func CreateRaceRepository(fields *PostgresConnection) repository_interfaces.IRaceRepository {
	dbx := sqlx.NewDb(fields.DB, "pgx")
	tracer := otel.Tracer("repository.Race")
	tracedDB := NewTracedDB(dbx, tracer)

	return NewRaceRepository(tracedDB)
}

func CreateRatingRepository(fields *PostgresConnection) repository_interfaces.IRatingRepository {
	dbx := sqlx.NewDb(fields.DB, "pgx")
	tracer := otel.Tracer("repository.Rating")
	tracedDB := NewTracedDB(dbx, tracer)

	return NewRatingRepository(tracedDB)
}
