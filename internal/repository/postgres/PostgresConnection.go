package postgres

import (
	"PPO_BMSTU/config"
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"database/sql"
	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
)

type PostgresConnection struct {
	DB     *sql.DB
	Config config.Config
}

func NewPostgresConnection(Postgres config.DbConnectionFlags, logger *log.Logger) (*PostgresConnection, error) {
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

	return NewCrewRepository(dbx)
}

func CreateJudgeRepository(fields *PostgresConnection) repository_interfaces.IJudgeRepository {
	dbx := sqlx.NewDb(fields.DB, "pgx")

	return NewJudgeRepository(dbx)
}

func CreateCrewResInRaceRepository(fields *PostgresConnection) repository_interfaces.ICrewResInRaceRepository {
	dbx := sqlx.NewDb(fields.DB, "pgx")

	return NewCrewResInRaceRepository(dbx)
}

func CreateParticipantRepository(fields *PostgresConnection) repository_interfaces.IParticipantRepository {
	dbx := sqlx.NewDb(fields.DB, "pgx")

	return NewParticipantRepository(dbx)
}

func CreateProtestRepository(fields *PostgresConnection) repository_interfaces.IProtestRepository {
	dbx := sqlx.NewDb(fields.DB, "pgx")

	return NewProtestRepository(dbx)
}
func CreateRaceRepository(fields *PostgresConnection) repository_interfaces.IRaceRepository {
	dbx := sqlx.NewDb(fields.DB, "pgx")

	return NewRaceRepository(dbx)
}
func CreateRatingRepository(fields *PostgresConnection) repository_interfaces.IRatingRepository {
	dbx := sqlx.NewDb(fields.DB, "pgx")

	return NewRatingRepository(dbx)
}
