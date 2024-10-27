package db_init

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Импортируйте драйвер PostgreSQL
)

// ConnectTestDatabasePostgres создает контейнер с PostgreSQL и возвращает подключение к базе данных
func ConnectTestDatabasePostgres() (*sqlx.DB, error) {
	postgresDSN := "host=localhost user=testuser dbname=testDB port=5433 sslmode=disable"

	// Подключаемся к базе данных
	db, err := sqlx.Connect("postgres", postgresDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres database: %w", err)
	}

	return db, nil
}
