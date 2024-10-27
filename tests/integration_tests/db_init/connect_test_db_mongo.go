package db_init

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectTestDatabaseMongo подключается к существующему контейнеру MongoDB и возвращает подключение к базе данных
func ConnectTestDatabaseMongo() (*mongo.Database, error) {
	mongoDSN := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(mongoDSN)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Проверка подключения
	if err = client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	// Возвращаем подключение к базе данных
	return client.Database("testDB"), nil // Указываем название базы данных для MongoDB
}
