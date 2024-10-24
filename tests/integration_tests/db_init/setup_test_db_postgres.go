package db_init

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
)

const (
	USER     = "postgres"
	PASSWORD = "7625"
	DBNAME   = "test"
)

func SetupTestDatabasePostgres() (testcontainers.Container, *sqlx.DB, error) {
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       DBNAME,
			"POSTGRES_PASSWORD": PASSWORD,
			"POSTGRES_USER":     USER,
		},
	}

	ctx := context.Background()
	dbContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: containerReq,
		Started:          true,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to start container: %w", err)
	}

	host, err := dbContainer.Host(ctx)
	if err != nil {
		return dbContainer, nil, fmt.Errorf("failed to get container host: %w", err)
	}
	port, err := dbContainer.MappedPort(ctx, "5432")
	if err != nil {
		return dbContainer, nil, fmt.Errorf("failed to get mapped port: %w", err)
	}

	dsnPGConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port.Int(), USER, PASSWORD, DBNAME)
	db, err := sqlx.Connect("pgx", dsnPGConn)
	if err != nil {
		return dbContainer, nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return dbContainer, nil, fmt.Errorf("failed to ping database: %w", err)
	}

	text, err := os.ReadFile("../../../../PPO_BMSTU/db/init.sql")
	if err != nil {
		return dbContainer, nil, fmt.Errorf("failed to read init.sql: %w", err)
	}

	if _, err := db.Exec(string(text)); err != nil {
		return dbContainer, nil, fmt.Errorf("failed to execute init.sql: %w", err)
	}

	return dbContainer, db, nil
}
