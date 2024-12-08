package mongo_init

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func SetupTestDatabase() (testcontainers.Container, *mongo.Database) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	containerReq := testcontainers.ContainerRequest{
		Image:        "mongo:latest",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForListeningPort("27017/tcp"),
		Env: map[string]string{
			"MONGO_INITDB_ROOT_USERNAME": USER,
			"MONGO_INITDB_ROOT_PASSWORD": PASSWORD,
		},
	}

	dbContainer, err := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		},
	)
	if err != nil {
		fmt.Println("Failed to start container: ", err)
		return nil, nil
	}

	// Получаем IP-адрес и порт контейнера
	host, err := dbContainer.Host(ctx)
	if err != nil {
		fmt.Println("Failed to get container host: ", err)
		return nil, nil
	}
	mappedPort, err := dbContainer.MappedPort(ctx, "27017")
	if err != nil {
		fmt.Println("Failed to get mapped port: ", err)
		return nil, nil
	}

	dsnMongoConn := fmt.Sprintf("mongodb://%s:%s@%s:%s", USER, PASSWORD, host, mappedPort.Port())

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsnMongoConn))
	if err != nil {
		fmt.Println("Failed to connect to MongoDB: ", err)
		return dbContainer, nil
	}

	// Send a ping to confirm a successful connection
	var result bson.M
	if err = client.Database(DBNAME).RunCommand(ctx, bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		fmt.Println("Failed to ping MongoDB: ", err)
		return dbContainer, nil
	}

	// Получите ссылку на вашу базу данных
	db := client.Database(DBNAME)

	return dbContainer, db
}
