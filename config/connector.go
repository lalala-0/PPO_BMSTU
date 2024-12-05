package config

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/charmbracelet/log"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type DbConnectionFlags struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
	DBName   string `mapstructure:"dbname"`
}

func (p *DbConnectionFlags) InitPostgresDB(logger *log.Logger) (*sql.DB, error) {
	logger.Debug("POSTGRES! Start init postgreSQL", "user", p.User, "DBName", p.DBName,
		"host", p.Host, "port", p.Port)

	dsnPGConn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		p.User, p.DBName, p.Password,
		p.Host, p.Port)

	db, err := sql.Open("pgx", dsnPGConn)
	if err != nil {
		logger.Fatal("POSTGRES! Error in method open")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Fatal("POSTGRES! Error in method ping")
		return nil, err
	}

	db.SetMaxOpenConns(10)

	logger.Info("POSTGRES! Successfully init postgreSQL")
	return db, nil
}

func (p *DbConnectionFlags) InitMongoDB(logger *log.Logger) (*mongo.Database, error) {
	logger.Debug("MONGO! Start init mongoDB", "user", p.User, "DBName", p.DBName, "host", p.Host, "port", p.Port)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	dsnMongoConn := "mongodb://" + p.Host + ":" + p.Port + "/"

	// Подключаемся к MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsnMongoConn))
	if err != nil {
		logger.Error("MONGO! Error connecting to MongoDB", "error", err)
		return nil, err
	}

	// Пингуем для проверки успешного подключения
	if err = client.Ping(ctx, nil); err != nil {
		logger.Error("MONGO! Failed to ping MongoDB", "error", err)
		return nil, err
	}

	// Сообщаем об успешной инициализации
	logger.Info("MONGO! Successfully initialized MongoDB")

	return client.Database(p.DBName), nil
}
